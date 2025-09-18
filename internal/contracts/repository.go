package contracts

import "Lesson15/internal/models"

type RepositoryI interface {
	CreateUser(user models.User) error
	GetUserById(id int) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id int) error
}
