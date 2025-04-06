package app

import (
	"context"
	"user_service/internal/core/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, userID string) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	ListUsers(ctx context.Context, limit, offset int64) (*models.UserList, error)
	UpdateUserPassword(ctx context.Context, user *models.User) error
}
