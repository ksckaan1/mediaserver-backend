package series

import (
	"context"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/pkg/gh"
	"mediaserver/internal/port"
	"net/http"
)

type SeriesService interface {
	// Series
	CreateSeries(ctx context.Context, series *model.Series) (string, error)
	GetSeriesByID(ctx context.Context, id string) (*model.GetSeriesResponse, error)
	ListSeries(ctx context.Context, limit, offset int64) (*model.SeriesList, error)
	UpdateSeriesByID(ctx context.Context, series *model.Series) error
	DeleteSeriesByID(ctx context.Context, id string) error

	// Season
	CreateSeason(ctx context.Context, season *model.Season) (string, error)
	GetSeasonByID(ctx context.Context, id string) (*model.Season, error)
	ListSeasonsBySeriesID(ctx context.Context, seriesID string, limit, offset int64) (*model.SeasonList, error)
	UpdateSeasonByID(ctx context.Context, season *model.Season) error
	DeleteSeasonByID(ctx context.Context, id string) error
}

type Series struct {
	seriesService SeriesService
	logger        port.Logger
}

func New(seriesService SeriesService, logger port.Logger) (*Series, error) {
	return &Series{
		seriesService: seriesService,
		logger:        logger,
	}, nil
}

type CreateSeriesRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	TMDBID      string `json:"tmdb_id"`
	MediaID     string `json:"media_id"`
}

type CreateSeriesResponse struct {
	ID string `json:"id"`
}

func (s *Series) CreateSeries(ctx context.Context, req *gh.Request[*CreateSeriesRequest]) (*gh.Response[*CreateSeriesResponse], error) {
	id, err := s.seriesService.CreateSeries(ctx, &model.Series{
		Title:       req.Body.Title,
		Description: req.Body.Description,
		TMDBID:      req.Body.TMDBID,
	})
	if err != nil {
		if errors.Is(err, customerrors.ErrSeriesNotFound) {
			return &gh.Response[*CreateSeriesResponse]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrMediaNotFound
		}
		return &gh.Response[*CreateSeriesResponse]{}, fmt.Errorf("seriesService.CreateSeries: %w", err)
	}
	return &gh.Response[*CreateSeriesResponse]{
		Body: &CreateSeriesResponse{
			ID: id,
		},
		StatusCode: http.StatusCreated,
	}, nil
}

func (s *Series) GetSeriesByID(ctx context.Context, req *gh.Request[any]) (*gh.Response[*model.GetSeriesResponse], error) {
	id := req.Params["id"]
	series, err := s.seriesService.GetSeriesByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrSeriesNotFound) {
			return &gh.Response[*model.GetSeriesResponse]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrSeriesNotFound
		}
		return &gh.Response[*model.GetSeriesResponse]{}, fmt.Errorf("seriesService.GetSeriesByID: %w", err)
	}
	return &gh.Response[*model.GetSeriesResponse]{
		Body:       series,
		StatusCode: http.StatusOK,
	}, nil
}

func (s *Series) ListSeries(ctx context.Context, req *gh.Request[any]) (*gh.Response[*model.SeriesList], error) {
	limit, err := req.GetQueryInt64("limit", -1)
	if err != nil {
		return &gh.Response[*model.SeriesList]{
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("req.GetQueryInt64 (limit): %w", err)
	}

	offset, err := req.GetQueryInt64("offset", 0)
	if err != nil {
		return &gh.Response[*model.SeriesList]{
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("req.GetQueryInt64 (offset): %w", err)
	}

	seriesList, err := s.seriesService.ListSeries(ctx, limit, offset)
	if err != nil {
		return &gh.Response[*model.SeriesList]{}, fmt.Errorf("seriesService.ListSeries: %w", err)
	}
	return &gh.Response[*model.SeriesList]{
		Body:       seriesList,
		StatusCode: http.StatusOK,
	}, nil
}

type UpdateSeriesByIDRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	TMDBID      string `json:"tmdb_id"`
	MediaID     string `json:"media_id"`
}

func (s *Series) UpdateSeriesByID(ctx context.Context, req *gh.Request[*UpdateSeriesByIDRequest]) (*gh.Response[any], error) {
	seriesID := req.Params["id"]

	err := s.seriesService.UpdateSeriesByID(ctx, &model.Series{
		ID:          seriesID,
		Title:       req.Body.Title,
		Description: req.Body.Description,
		TMDBID:      req.Body.TMDBID,
	})
	if err != nil {
		if errors.Is(err, customerrors.ErrSeriesNotFound) {
			return &gh.Response[any]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrSeriesNotFound
		}
		return &gh.Response[any]{}, fmt.Errorf("seriesService.UpdateSeriesByID: %w", err)
	}
	return &gh.Response[any]{
		StatusCode: http.StatusOK,
	}, nil
}

func (s *Series) DeleteSeriesByID(ctx context.Context, req *gh.Request[any]) (*gh.Response[any], error) {
	seriesID := req.Params["id"]

	err := s.seriesService.DeleteSeriesByID(ctx, seriesID)
	if err != nil {
		if errors.Is(err, customerrors.ErrSeriesNotFound) {
			return &gh.Response[any]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrSeriesNotFound
		}
		return &gh.Response[any]{}, fmt.Errorf("seriesService.DeleteSeriesByID: %w", err)
	}
	return &gh.Response[any]{
		StatusCode: http.StatusNoContent,
	}, nil
}
