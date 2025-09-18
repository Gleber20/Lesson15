package repository

import (
	"Lesson15/internal/errs"
	"Lesson15/internal/models"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) translateError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errs.ErrNotfound
	default:
		return err
	}
}

func (r *UserRepository) CreateUser(user models.User) error {
	return r.db.QueryRow(`INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id`,
		user.Name, user.Email, user.Age).Scan(&user.ID)
}

func (r *UserRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return models.User{}, r.translateError(err)
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user models.User) error {
	_, err := r.db.Exec(
		`UPDATE users SET name=$1, email=$2, age=$3 WHERE id=$4`,
		user.Name, user.Email, user.Age, user.ID,
	)
	if err != nil {
		return r.translateError(err)
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return r.translateError(err)
	}
	return nil
}
