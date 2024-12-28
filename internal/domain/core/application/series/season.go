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

type CreateSeasonRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	SeriesID    string `json:"series_id"`
	Order       int64  `json:"order"`
}

type CreateSeasonResponse struct {
	ID string `json:"id"`
}

func (s *Series) CreateSeason(ctx context.Context, req *gh.Request[*CreateSeasonRequest]) (*gh.Response[*CreateSeasonResponse], error) {
	id, err := s.seriesService.CreateSeason(ctx, &model.Season{
		Title:       req.Body.Title,
		Description: req.Body.Description,
		SeriesID:    req.Body.SeriesID,
		Order:       req.Body.Order,
	})
	if err != nil {
		return nil, customerrors.ErrInternalServerError
	}
	return &gh.Response[*CreateSeasonResponse]{
		Body: &CreateSeasonResponse{
			ID: id,
		},
		StatusCode: http.StatusCreated,
	}, nil
}

func (s *Series) GetSeasonByID(ctx context.Context, req *gh.Request[any]) (*gh.Response[*model.Season], error) {
	id := req.Params["id"]
	season, err := s.seriesService.GetSeasonByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrSeasonNotFound) {
			return &gh.Response[*model.Season]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrSeasonNotFound
		}
		return nil, customerrors.ErrInternalServerError
	}
	return &gh.Response[*model.Season]{
		Body:       season,
		StatusCode: http.StatusOK,
	}, nil
}

func (s *Series) ListSeasonsBySeriesID(ctx context.Context, req *gh.Request[any]) (*gh.Response[*model.SeasonList], error) {
	seriesID := req.Params["series_id"]
	limit, err := req.GetQueryInt64("limit", -1)
	if err != nil {
		return nil, fmt.Errorf("req.GetQueryInt64 (limit): %w", err)
	}
	offset, err := req.GetQueryInt64("offset", 0)
	if err != nil {
		return nil, fmt.Errorf("req.GetQueryInt64 (offset): %w", err)
	}
	seasons, err := s.seriesService.ListSeasonsBySeriesID(ctx, seriesID, limit, offset)
	if err != nil {
		if errors.Is(err, customerrors.ErrSeriesNotFound) {
			return &gh.Response[*model.SeasonList]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrSeriesNotFound
		}
		return nil, customerrors.ErrInternalServerError
	}
	return &gh.Response[*model.SeasonList]{
		Body:       seasons,
		StatusCode: http.StatusOK,
	}, nil
}

type UpdateSeasonByIDRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Order       int64  `json:"order"`
}

func (s *Series) UpdateSeasonByID(ctx context.Context, req *gh.Request[*UpdateSeasonByIDRequest]) (*gh.Response[any], error) {
	id := req.Params["id"]
	err := s.seriesService.UpdateSeasonByID(ctx, &model.Season{
		ID:          id,
		Title:       req.Body.Title,
		Description: req.Body.Description,
		Order:       req.Body.Order,
	})
	if err != nil {
		if errors.Is(err, customerrors.ErrSeasonNotFound) {
			return &gh.Response[any]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrSeasonNotFound
		}
		return nil, customerrors.ErrInternalServerError
	}
	return &gh.Response[any]{
		StatusCode: http.StatusNoContent,
	}, nil
}

func (s *Series) DeleteSeasonByID(ctx context.Context, req *gh.Request[any]) (*gh.Response[any], error) {
	id := req.Params["id"]
	err := s.seriesService.DeleteSeasonByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrSeasonNotFound) {
			return &gh.Response[any]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrSeasonNotFound
		}
		return nil, customerrors.ErrInternalServerError
	}
	return &gh.Response[any]{
		StatusCode: http.StatusNoContent,
	}, nil
}
