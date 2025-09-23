package repository

import (
	"Lesson15/internal/errs"
	"Lesson15/internal/models"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/redis/go-redis/v9"
)

type EmployeeRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (r *EmployeeRepository) translateError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errs.ErrNotfound
	default:
		return err
	}
}

func (r *EmployeeRepository) CreateEmployee(ctx context.Context, employee models.Employee) error {
	return r.db.QueryRowContext(ctx, `INSERT INTO employees (name, email, age) VALUES ($1, $2, $3) RETURNING id`,
		employee.Name, employee.Email, employee.Age).Scan(&employee.ID)
}

func (r *EmployeeRepository) GetEmployeeByID(ctx context.Context, id int) (models.Employee, error) {
	var user models.Employee
	err := r.db.GetContext(ctx, &user, "SELECT * FROM employees WHERE id = $1", id)
	if err != nil {
		return models.Employee{}, r.translateError(err)
	}
	return user, nil
}

func (r *EmployeeRepository) UpdateEmployee(ctx context.Context, employee models.Employee) error {
	_, err := r.db.ExecContext(
		ctx, `UPDATE employees SET name=$1, email=$2, age=$3 WHERE id=$4`,
		employee.Name, employee.Email, employee.Age, employee.ID,
	)
	if err != nil {
		return r.translateError(err)
	}
	return nil
}

func (r *EmployeeRepository) DeleteEmployee(ctx context.Context, id int) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM employees WHERE id = $1", id)
	if err != nil {
		return r.translateError(err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return r.translateError(err)
	}
	if rows == 0 {
		return errs.ErrNotfound
	}
	return nil
}
