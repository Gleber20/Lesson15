package contracts

import (
	"Lesson15/internal/models"
	"context"
)

type RepositoryI interface {
	CreateEmployee(ctx context.Context, user models.Employee) error
	GetEmployeeByID(ctx context.Context, id int) (models.Employee, error)
	UpdateEmployee(ctx context.Context, user models.Employee) error
	DeleteEmployee(ctx context.Context, id int) error
}
