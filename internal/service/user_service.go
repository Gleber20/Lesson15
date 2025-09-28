package service

import (
	"Lesson15/internal/config"
	"Lesson15/internal/errs"
	"Lesson15/internal/models"
	"Lesson15/internal/utils"
	"Lesson15/pkg"
	"context"
	"errors"
	_ "golang.org/x/tools/go/cfg"
)

func (s *EmployeeService) CreateUser(ctx context.Context, user models.User) (err error) {
	// Проверить существует ли пользователь с таким username'ом в бд
	_, err = s.repo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return err
		}
	} else {
		return errs.ErrUsernameAlreadyExists
	}

	// За хэшировать пароль
	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return err
	}

	// Добавить пользователя в бд
	if err = s.repo.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *EmployeeService) Authenticate(ctx context.Context, user models.User) (string, string, error) {
	// проверить существует ли пользователь с таким username
	userFromDB, err := s.repo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return "", "", errs.ErrUserNotFound
		}

		return "", "", err
	}

	// за хэшировать пароль, который получили от пользователя
	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return "", "", err
	}

	// проверить правильно ли он указал пароль
	if userFromDB.Password != user.Password {
		return "", "", errs.ErrIncorrectUsernameOrPassword
	}

	cfg := config.LoadConfig()
	accessToken, err := pkg.GenerateToken(
		userFromDB.ID,
		cfg.AuthConfig.AccessTokenTTLMinutes,
		cfg.AuthConfig.JWTSecret,
		false)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := pkg.GenerateToken(
		userFromDB.ID,
		cfg.AuthConfig.RefreshTokenTTLDays,
		cfg.AuthConfig.JWTSecret,
		true)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
