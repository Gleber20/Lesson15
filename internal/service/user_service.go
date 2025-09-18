package service

import (
	"Lesson15/internal/errs"
	"Lesson15/internal/models"
	"Lesson15/internal/repository"
	"errors"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.User) error {
	err := s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserById(userId int) (models.User, error) {
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return models.User{}, errs.ErrUserNotfound
		}
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(user models.User) error {
	_, err := s.repo.GetUserById(user.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotfound
		}
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(userId int) error {
	_, err := s.repo.GetUserById(userId)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotfound
		}
		return err
	}
	err = s.repo.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}
