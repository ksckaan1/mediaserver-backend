package app

import (
	"context"
	"season_service/internal/core/models"
)

type Repository interface {
	CreateSeason(ctx context.Context, season *models.Season) error
	DeleteSeasonByID(ctx context.Context, id string) error
	GetSeasonByID(ctx context.Context, id string) (*models.Season, error)
	ListSeasonsBySeriesID(ctx context.Context, seriesID string) ([]*models.Season, error)
	UpdateSeasonByID(ctx context.Context, season *models.Season) error
	UpdateSeasonOrderByID(ctx context.Context, season *models.Season) error
}
