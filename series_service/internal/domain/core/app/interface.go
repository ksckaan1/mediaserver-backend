package app

import (
	"context"
	"series_service/internal/domain/core/models"
)

type Repository interface {
	CreateSeries(ctx context.Context, series *models.Series) error
	GetSeriesByID(ctx context.Context, id string) (*models.Series, error)
	ListSeries(ctx context.Context, limit, offset int64) (*models.SeriesList, error)
	ListSeriesWithIDs(ctx context.Context, ids []string, limit, offset int64) (*models.SeriesList, error)
	ListSeriesSearch(ctx context.Context) ([]*models.SeriesSearch, error)
	UpdateSeriesByID(ctx context.Context, series *models.Series) error
	DeleteSeriesByID(ctx context.Context, id string) error
}

type IDGenerator interface {
	NewID() string
}
