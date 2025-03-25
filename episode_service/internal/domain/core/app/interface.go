package app

import (
	"context"

	"episode_service/internal/domain/core/models"
)

type Repository interface {
	CreateEpisode(ctx context.Context, episode *models.Episode) error
	GetEpisodeByID(ctx context.Context, episodeID string) (*models.Episode, error)
	ListEpisodesBySeasonID(ctx context.Context, seasonID string) ([]*models.Episode, error)
	UpdateEpisodeByID(ctx context.Context, episode *models.Episode) error
	UpdateEpisodeOrder(ctx context.Context, episode *models.Episode) error
	DeleteEpisodeByID(ctx context.Context, episodeID string) error
	DeleteAllEpisodesBySeasonID(ctx context.Context, seasonID string) error
}

type IDGenerator interface {
	NewID() string
}
