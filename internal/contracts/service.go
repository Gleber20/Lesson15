package contracts

import (
	"Lesson15/internal/models"
	"context"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ServiceI interface {
	CreateEmployee(ctx context.Context, user models.Employee) error
	GetEmployeeByID(ctx context.Context, userId int) (models.Employee, error)
	UpdateEmployee(ctx context.Context, user models.Employee) error
	DeleteEmployee(ctx context.Context, userId int) error

	CreateUser(ctx context.Context, user models.User) (err error)
	Authenticate(ctx context.Context, user models.User) (string, string, error)
}
