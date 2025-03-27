package app

import (
	"context"
	"tmdb_service/internal/core/models"
)

type TMDBClient interface {
	GetMovieDetail(ctx context.Context, id string) (map[string]any, error)
	GetSeriesDetail(ctx context.Context, id string) (map[string]any, error)
}

type Repository interface {
	GetTMDBInfo(ctx context.Context, id string) (*models.TMDBInfo, error)
	SetTMDBInfo(ctx context.Context, info *models.TMDBInfo) error
}
