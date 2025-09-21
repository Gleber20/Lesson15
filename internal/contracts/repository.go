package contracts

import (
	"Lesson15/internal/models"
	"context"
)

type RepositoryI interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserById(ctx context.Context, id int) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id int) error
}
