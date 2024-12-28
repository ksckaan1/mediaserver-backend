package series

import (
	"context"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/pkg/gh"
	"net/http"
)

type CreateEpisodeRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	SeasonID    string `json:"season_id"`
	Order       int64  `json:"order"`
	MediaID     string `json:"media_id"`
}

type CreateEpisodeResponse struct {
	ID string `json:"id"`
}

func (s *Series) CreateEpisode(ctx context.Context, req *gh.Request[*CreateEpisodeRequest]) (*gh.Response[*CreateEpisodeResponse], error) {
	id, err := s.seriesService.CreateEpisode(ctx, &model.Episode{
		Title:       req.Body.Title,
		Description: req.Body.Description,
		SeasonID:    req.Body.SeasonID,
		Order:       req.Body.Order,
		MediaID:     req.Body.MediaID,
	})
	if err != nil {
		if errors.Is(err, customerrors.ErrMediaNotFound) {
			return &gh.Response[*CreateEpisodeResponse]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrMediaNotFound
		}
		if errors.Is(err, customerrors.ErrSeasonNotFound) {
			return &gh.Response[*CreateEpisodeResponse]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrSeasonNotFound
		}
		return &gh.Response[*CreateEpisodeResponse]{}, fmt.Errorf("seriesService.CreateEpisode: %w", err)
	}

	return &gh.Response[*CreateEpisodeResponse]{
		Body: &CreateEpisodeResponse{
			ID: id,
		},
		StatusCode: http.StatusCreated,
	}, nil
}

func (s *Series) GetEpisodeByID(ctx context.Context, req *gh.Request[any]) (*gh.Response[*model.GetEpisodeByIDResponse], error) {
	id := req.Params["id"]

	movie, err := s.seriesService.GetEpisodeByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrEpisodeNotFound) {
			return &gh.Response[*model.GetEpisodeByIDResponse]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrEpisodeNotFound
		}
		if errors.Is(err, customerrors.ErrMediaNotFound) {
			return &gh.Response[*model.GetEpisodeByIDResponse]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrMediaNotFound
		}
		return &gh.Response[*model.GetEpisodeByIDResponse]{}, fmt.Errorf("seriesService.GetEpisodeByID: %w", err)
	}

	return &gh.Response[*model.GetEpisodeByIDResponse]{
		Body:       movie,
		StatusCode: http.StatusOK,
	}, nil
}

func (s *Series) ListEpisodesBySeasonID(ctx context.Context, req *gh.Request[any]) (*gh.Response[*model.EpisodeList], error) {
	seasonID := req.Params["season_id"]

	limit, err := req.GetQueryInt64("limit", -1)
	if err != nil {
		return &gh.Response[*model.EpisodeList]{
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("req.GetQueryInt64 (limit): %w", err)
	}

	offset, err := req.GetQueryInt64("offset", 0)
	if err != nil {
		return &gh.Response[*model.EpisodeList]{
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("req.GetQueryInt64 (offset): %w", err)
	}

	episodes, err := s.seriesService.ListEpisodesBySeasonID(ctx, seasonID, limit, offset)
	if err != nil {
		if errors.Is(err, customerrors.ErrSeasonNotFound) {
			return &gh.Response[*model.EpisodeList]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrSeasonNotFound
		}
		return &gh.Response[*model.EpisodeList]{}, fmt.Errorf("seriesService.ListEpisodesBySeasonID: %w", err)
	}

	return &gh.Response[*model.EpisodeList]{
		Body:       episodes,
		StatusCode: http.StatusOK,
	}, nil
}

type UpdateEpisodeByIDRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Order       int64  `json:"order"`
	MediaID     string `json:"media_id"`
}

func (s *Series) UpdateEpisodeByID(ctx context.Context, req *gh.Request[*UpdateEpisodeByIDRequest]) (*gh.Response[any], error) {
	episodeID := req.Params["id"]
	err := s.seriesService.UpdateEpisodeByID(ctx, &model.Episode{
		ID:          episodeID,
		Title:       req.Body.Title,
		Description: req.Body.Description,
		Order:       req.Body.Order,
		MediaID:     req.Body.MediaID,
	})
	if err != nil {
		if errors.Is(err, customerrors.ErrEpisodeNotFound) {
			return &gh.Response[any]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrEpisodeNotFound
		}
		if errors.Is(err, customerrors.ErrMediaNotFound) {
			return &gh.Response[any]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrMediaNotFound
		}
		return &gh.Response[any]{}, fmt.Errorf("seriesSerivce.UpdateEpisodeByID: %w", err)
	}

	return &gh.Response[any]{
		StatusCode: http.StatusNoContent,
	}, nil
}

func (s *Series) DeleteEpisodeByID(ctx context.Context, req *gh.Request[any]) (*gh.Response[any], error) {
	episodeID := req.Params["id"]
	err := s.seriesService.DeleteEpisodeByID(ctx, episodeID)
	if err != nil {
		if errors.Is(err, customerrors.ErrEpisodeNotFound) {
			return &gh.Response[any]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrEpisodeNotFound
		}
		return &gh.Response[any]{}, fmt.Errorf("seriesService.DeleteEpisodeByID: %w", err)
	}
	return &gh.Response[any]{
		StatusCode: http.StatusNoContent,
	}, nil
}
