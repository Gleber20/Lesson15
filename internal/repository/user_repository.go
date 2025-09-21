package repository

import (
	"Lesson15/internal/errs"
	"Lesson15/internal/models"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) translateError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errs.ErrNotfound
	default:
		return err
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user models.User) error {
	return r.db.QueryRowContext(ctx, `INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id`,
		user.Name, user.Email, user.Age).Scan(&user.ID)
}

func (r *UserRepository) GetUserById(ctx context.Context, id int) (models.User, error) {
	var user models.User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return models.User{}, r.translateError(err)
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user models.User) error {
	_, err := r.db.ExecContext(
		ctx, `UPDATE users SET name=$1, email=$2, age=$3 WHERE id=$4`,
		user.Name, user.Email, user.Age, user.ID,
	)
	if err != nil {
		return r.translateError(err)
	}
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
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
