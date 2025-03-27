package app

import (
	"context"
	"series_service/internal/core/models"
)

type Repository interface {
	CreateSeries(ctx context.Context, series *models.Series) error
	GetSeriesByID(ctx context.Context, id string) (*models.Series, error)
	ListSeries(ctx context.Context, limit, offset int64) (*models.SeriesList, error)
	UpdateSeriesByID(ctx context.Context, series *models.Series) error
	DeleteSeriesByID(ctx context.Context, id string) error
}
