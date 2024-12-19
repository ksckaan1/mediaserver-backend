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
	CreateSeries(ctx context.Context, series *model.Series) (string, error)
	GetSeriesByID(ctx context.Context, id string) (*model.Series, error)
	ListSeries(ctx context.Context, limit, offset int64) (*model.SeriesList, error)
	UpdateSeriesByID(ctx context.Context, series *model.Series) error
	DeleteSeriesByID(ctx context.Context, id string) error
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
	TMDBID      int64  `json:"tmdb_id"`
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

func (s *Series) GetSeriesByID(ctx context.Context, req *gh.Request[any]) (*gh.Response[*model.Series], error) {
	id := req.Params["id"]
	series, err := s.seriesService.GetSeriesByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrSeriesNotFound) {
			return &gh.Response[*model.Series]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrSeriesNotFound
		}
		return &gh.Response[*model.Series]{}, fmt.Errorf("seriesService.GetSeriesByID: %w", err)
	}
	return &gh.Response[*model.Series]{
		Body:       series,
		StatusCode: http.StatusOK,
	}, nil
}
