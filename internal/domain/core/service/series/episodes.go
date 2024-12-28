package series

import (
	"context"
	"fmt"
	"mediaserver/internal/domain/core/model"
)

func (s *Series) CreateEpisode(ctx context.Context, episode *model.Episode) (string, error) {
	_, err := s.repo.GetSeasonByID(ctx, episode.SeasonID)
	if err != nil {
		err = fmt.Errorf("repo.GetSeasonByID: %w", err)
		s.logger.Error(ctx, "error when getting season",
			"error", err,
		)
		return "", err
	}
	if episode.MediaID != "" {
		_, err := s.repo.GetMediaByID(ctx, episode.MediaID)
		if err != nil {
			err = fmt.Errorf("repo.GetMediaByID: %w", err)
			s.logger.Error(ctx, "error when getting media",
				"error", err,
			)
			return "", err
		}
	}
	episode.ID = s.idgen.NewID()
	err = s.repo.CreateEpisode(ctx, episode)
	if err != nil {
		err = fmt.Errorf("repo.CreateEpisode: %w", err)
		s.logger.Error(ctx, "error when creating episode",
			"error", err,
		)
		return "", err
	}
	s.logger.Info(ctx, "episode created",
		"episode_id", episode.ID,
		"season_id", episode.SeasonID,
	)
	return episode.ID, nil
}

func (s *Series) GetEpisodeByID(ctx context.Context, episodeID string) (*model.GetEpisodeByIDResponse, error) {
	episode, err := s.repo.GetEpisodeByID(ctx, episodeID)
	if err != nil {
		err = fmt.Errorf("repo.GetEpisodeByID: %w", err)
		s.logger.Error(ctx, "error when getting episode",
			"error", err,
		)
		return nil, err
	}
	var mediaInfo *model.Media
	if episode.MediaID != "" {
		media, err := s.repo.GetMediaByID(ctx, episode.MediaID)
		if err != nil {
			err = fmt.Errorf("repo.GetMediaByID: %w", err)
			s.logger.Error(ctx, "error when getting media",
				"error", err,
			)
			return nil, err
		}
		mediaInfo = media
	}
	s.logger.Info(ctx, "episode retrieved",
		"episode_id", episode.ID,
		"season_id", episode.SeasonID,
	)
	return &model.GetEpisodeByIDResponse{
		ID:          episodeID,
		CreatedAt:   episode.CreatedAt,
		UpdatedAt:   episode.UpdatedAt,
		Title:       episode.Title,
		Description: episode.Description,
		SeasonID:    episode.SeasonID,
		Order:       episode.Order,
		MediaInfo:   mediaInfo,
	}, nil
}

func (s *Series) ListEpisodesBySeasonID(ctx context.Context, seasonID string, limit, offset int64) (*model.EpisodeList, error) {
	_, err := s.repo.GetSeasonByID(ctx, seasonID)
	if err != nil {
		err = fmt.Errorf("repo.GetSeasonByID: %w", err)
		s.logger.Error(ctx, "error when getting season",
			"error", err,
		)
		return nil, err
	}
	episodes, err := s.repo.ListEpisodesBySeasonID(ctx, seasonID, limit, offset)
	if err != nil {
		err = fmt.Errorf("repo.ListEpisodesBySeasonID: %w", err)
		s.logger.Error(ctx, "error when listing episodes",
			"error", err,
		)
		return nil, err
	}
	s.logger.Info(ctx, "episodes listed",
		"season_id", seasonID,
		"limit", limit,
		"offset", offset,
	)
	return episodes, nil
}

func (s *Series) UpdateEpisodeByID(ctx context.Context, episode *model.Episode) error {
	if episode.MediaID != "" {
		_, err := s.repo.GetMediaByID(ctx, episode.MediaID)
		if err != nil {
			err = fmt.Errorf("repo.GetMediaByID: %w", err)
			s.logger.Error(ctx, "error when getting media",
				"error", err,
			)
			return err
		}
	}
	err := s.repo.UpdateEpisodeByID(ctx, episode)
	if err != nil {
		err = fmt.Errorf("repo.UpdateEpisodeByID: %w", err)
		s.logger.Error(ctx, "error when updating episode",
			"error", err,
		)
		return err
	}
	s.logger.Info(ctx, "episode updated",
		"episode_id", episode.ID,
		"season_id", episode.SeasonID,
	)
	return nil
}

func (s *Series) DeleteEpisodeByID(ctx context.Context, episodeID string) error {
	err := s.repo.DeleteEpisodeByID(ctx, episodeID)
	if err != nil {
		err = fmt.Errorf("repo.DeleteEpisodeByID: %w", err)
		s.logger.Error(ctx, "error when deleting episode",
			"error", err,
		)
		return err
	}
	s.logger.Info(ctx, "episode deleted",
		"episode_id", episodeID,
	)
	return nil
}
