package repository

import (
	"Lesson15/internal/models"
	"context"
	"github.com/rs/zerolog"
	"os"
)

func (r *EmployeeRepository) CreateUser(ctx context.Context, user models.User) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "CreateUser").Logger()
	_, err = r.db.ExecContext(ctx, `INSERT INTO users (full_name, username, password)
					VALUES ($1, $2, $3)`,
		user.FullName,
		user.Username,
		user.Password)
	if err != nil {
		logger.Err(err).Msg("error inserting user")
		return r.translateError(err)
	}

	return nil
}

func (r *EmployeeRepository) GetUserByID(ctx context.Context, id int) (user models.User, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "GetUserByID").Logger()
	if err = r.db.GetContext(ctx, &user, `
		SELECT id, full_name, username, password, created_at, updated_at 
		FROM users
		WHERE id = $1`, id); err != nil {
		logger.Err(err).Msg("error selecting user")
		return models.User{}, r.translateError(err)
	}

	return user, nil
}

func (r *EmployeeRepository) GetUserByUsername(ctx context.Context, username string) (user models.User, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "GetUserByUsername").Logger()
	if err = r.db.GetContext(ctx, &user, `
		SELECT id, full_name, username, password, created_at, updated_at 
		FROM users
		WHERE username = $1`, username); err != nil {
		logger.Err(err).Msg("error selecting user")
		return models.User{}, r.translateError(err)
	}

	return user, nil
}
