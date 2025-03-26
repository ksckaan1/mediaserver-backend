package app

import (
	"context"
	"media_service/internal/core/models"
)

type Repository interface {
	CreateMedia(ctx context.Context, media *models.Media) error
	GetMediaByID(ctx context.Context, id string) (*models.Media, error)
	ListMedias(ctx context.Context, limit, offset int64) (*models.MediaList, error)
	UpdateMediaByID(ctx context.Context, media *models.Media) error
	DeleteMediaByID(ctx context.Context, id string) error
}
