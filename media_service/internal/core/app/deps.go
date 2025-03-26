package app

import (
	"context"
	"media_service/internal/core/model"
)

type Repository interface {
	CreateMedia(ctx context.Context, media *model.Media) error
	GetMediaByID(ctx context.Context, id string) (*model.Media, error)
	ListMedias(ctx context.Context, limit, offset int64) (*model.MediaList, error)
	UpdateMediaByID(ctx context.Context, media *model.Media) error
	DeleteMediaByID(ctx context.Context, id string) error
}
