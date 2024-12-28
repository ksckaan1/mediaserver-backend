package series

import (
	"context"
	"fmt"
	"mediaserver/internal/domain/core/model"
)

func (s *Series) CreateSeason(ctx context.Context, season *model.Season) (string, error) {
	season.ID = s.idgen.NewID()
	err := s.repo.CreateSeason(ctx, season)
	if err != nil {
		err = fmt.Errorf("repo.CreateSeason: %w", err)
		s.logger.Error(ctx, "error when creating season",
			"error", err,
		)
		return "", err
	}
	s.logger.Info(ctx, "season created",
		"season_id", season.ID,
		"series_id", season.SeriesID,
	)
	return season.ID, nil
}

func (s *Series) GetSeasonByID(ctx context.Context, seasonID string) (*model.Season, error) {
	season, err := s.repo.GetSeasonByID(ctx, seasonID)
	if err != nil {
		err = fmt.Errorf("repo.GetSeasonByID: %w", err)
		s.logger.Error(ctx, "error when getting season",
			"error", err,
			"season_id", seasonID,
		)
		return nil, err
	}
	s.logger.Info(ctx, "season retrieved",
		"season_id", season.ID,
		"series_id", season.SeriesID,
	)
	return season, nil
}

func (s *Series) ListSeasonsBySeriesID(ctx context.Context, seriesID string, limit, offset int64) (*model.SeasonList, error) {
	_, err := s.repo.GetSeriesByID(ctx, seriesID)
	if err != nil {
		err = fmt.Errorf("repo.GetSeriesByID: %w", err)
		s.logger.Error(ctx, "error when getting series",
			"error", err,
			"series_id", seriesID,
		)
		return nil, err
	}
	seasons, err := s.repo.ListSeasonsBySeriesID(ctx, seriesID, limit, offset)
	if err != nil {
		err = fmt.Errorf("repo.ListSeasonsBySeriesID: %w", err)
		s.logger.Error(ctx, "error when listing seasons",
			"error", err,
			"series_id", seriesID,
		)
		return nil, err
	}
	s.logger.Info(ctx, "seasons retrieved",
		"series_id", seriesID,
		"count", seasons.Count,
	)
	return seasons, nil
}

func (s *Series) UpdateSeasonByID(ctx context.Context, season *model.Season) error {
	err := s.repo.UpdateSeasonByID(ctx, season)
	if err != nil {
		err = fmt.Errorf("repo.UpdateSeasonByID: %w", err)
		s.logger.Error(ctx, "error when updating season",
			"error", err,
			"season_id", season.ID,
		)
		return err
	}
	s.logger.Info(ctx, "season updated",
		"season_id", season.ID,
		"series_id", season.SeriesID,
	)
	return nil
}

func (s *Series) DeleteSeasonByID(ctx context.Context, id string) error {
	err := s.repo.DeleteSeasonByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("repo.DeleteSeasonByID: %w", err)
		s.logger.Error(ctx, "error when deleting season",
			"error", err,
			"season_id", id,
		)
		return err
	}
	s.logger.Info(ctx, "season deleted",
		"season_id", id,
	)
	return nil
}
