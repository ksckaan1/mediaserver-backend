package app

import (
	"context"
	"setting_service/internal/core/models"
)

type Repository interface {
	Set(ctx context.Context, setting *models.Setting) error
	Get(ctx context.Context, key string) (*models.Setting, error)
	List(ctx context.Context, limit, offset int64) (*models.SettingList, error)
	Delete(ctx context.Context, key string) error
}
