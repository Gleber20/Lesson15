package service

import (
	"Lesson15/internal/models"
	"Lesson15/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUser(user models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) GetUserById(userId int) (models.User, error) {
	return s.repo.GetUserById(userId)
}

func (s *UserService) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}
