package contracts

import "Lesson15/internal/models"

type ServiceI interface {
	CreateUser(user models.User) error
	GetUserById(userId int) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(userId int) error
}
