package app

import (
	"auth_service/internal/core/models"
	"context"
)

type Repository interface {
	CreateSession(ctx context.Context, session *models.Session) error
	GetSession(ctx context.Context, sessionId string) (*models.Session, error)
	DeleteSession(ctx context.Context, sessionId string) error
	DeleteAllSessionsByUserID(ctx context.Context, userId string) error
}
