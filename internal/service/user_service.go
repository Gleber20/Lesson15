package service

import (
	"Lesson15/internal/errs"
	"Lesson15/internal/models"
	"Lesson15/internal/utils"
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

func (s *EmployeeService) Authenticate(ctx context.Context, user models.User) (int, error) {
	// проверить существует ли пользователь с таким username
	userFromDB, err := s.repo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return 0, errs.ErrUserNotFound
		}

		return 0, err
	}

	// за хэшировать пароль, который получили от пользователя
	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return 0, err
	}

	// проверить правильно ли он указал пароль
	if userFromDB.Password != user.Password {
		return 0, errs.ErrIncorrectUsernameOrPassword
	}

	return userFromDB.ID, nil
}
