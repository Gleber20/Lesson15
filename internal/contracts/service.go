package contracts

import (
	"Lesson15/internal/models"
	"context"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ServiceI interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserById(ctx context.Context, userId int) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, userId int) error
}
