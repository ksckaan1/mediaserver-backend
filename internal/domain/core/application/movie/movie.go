package movie

import (
	"context"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/pkg/gh"
	"net/http"
)

type MovieService interface {
	CreateMovie(ctx context.Context, movie *model.Movie) (string, error)
	GetMovieByID(ctx context.Context, id string) (*model.GetMovieByIDResponse, error)
	ListMovies(ctx context.Context, limit, offset int64) (*model.MovieList, error)
	UpdateMovieByID(ctx context.Context, movie *model.Movie) error
	DeleteMovieByID(ctx context.Context, id string) error
}

type Movie struct {
	movieService MovieService
}

func New(movieService MovieService) (*Movie, error) {
	return &Movie{
		movieService: movieService,
	}, nil
}

type CreateMovieRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	TMDBID      string `json:"tmdb_id"`
	MediaID     string `json:"media_id"`
}

type CreateMovieResponse struct {
	ID string `json:"id"`
}

func (m *Movie) CreateMovie(ctx context.Context, req *gh.Request[*CreateMovieRequest]) (*gh.Response[*CreateMovieResponse], error) {
	id, err := m.movieService.CreateMovie(ctx, &model.Movie{
		Title:       req.Body.Title,
		Description: req.Body.Description,
		TMDBID:      req.Body.TMDBID,
		MediaID:     req.Body.MediaID,
	})
	if err != nil {
		if errors.Is(err, customerrors.ErrMediaNotFound) {
			return &gh.Response[*CreateMovieResponse]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrMediaNotFound
		}
		return &gh.Response[*CreateMovieResponse]{}, fmt.Errorf("movieService.CreateMovie: %w", err)
	}

	return &gh.Response[*CreateMovieResponse]{
		Body: &CreateMovieResponse{
			ID: id,
		},
		StatusCode: http.StatusCreated,
	}, nil
}

func (m *Movie) GetMovieByID(ctx context.Context, req *gh.Request[any]) (*gh.Response[*model.GetMovieByIDResponse], error) {
	id := req.Params["id"]

	movie, err := m.movieService.GetMovieByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrMovieNotFound) {
			return &gh.Response[*model.GetMovieByIDResponse]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrMovieNotFound
		}
		return &gh.Response[*model.GetMovieByIDResponse]{}, fmt.Errorf("movieService.GetMovieByID: %w", err)
	}

	return &gh.Response[*model.GetMovieByIDResponse]{
		Body:       movie,
		StatusCode: http.StatusOK,
	}, nil
}

func (m *Movie) ListMovies(ctx context.Context, req *gh.Request[any]) (*gh.Response[*model.MovieList], error) {
	limit, err := req.GetQueryInt64("limit", -1)
	if err != nil {
		return &gh.Response[*model.MovieList]{
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("req.GetQueryInt64 (limit): %w", err)
	}

	offset, err := req.GetQueryInt64("offset", 0)
	if err != nil {
		return &gh.Response[*model.MovieList]{
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("req.GetQueryInt64 (offset): %w", err)
	}

	movies, err := m.movieService.ListMovies(ctx, limit, offset)
	if err != nil {
		return &gh.Response[*model.MovieList]{}, fmt.Errorf("movieService.ListMovies: %w", err)
	}

	return &gh.Response[*model.MovieList]{
		Body:       movies,
		StatusCode: http.StatusOK,
	}, nil
}

type UpdateMovieByIDRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	TMDBID      string `json:"tmdb_id"`
	MediaID     string `json:"media_id"`
}

func (m *Movie) UpdateMovieByID(ctx context.Context, req *gh.Request[*UpdateMovieByIDRequest]) (*gh.Response[any], error) {
	movieID := req.Params["id"]

	err := m.movieService.UpdateMovieByID(ctx, &model.Movie{
		ID:          movieID,
		Title:       req.Body.Title,
		Description: req.Body.Description,
		TMDBID:      req.Body.TMDBID,
		MediaID:     req.Body.MediaID,
	})
	if err != nil {
		if errors.Is(err, customerrors.ErrMovieNotFound) {
			return &gh.Response[any]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrMovieNotFound
		}
		if errors.Is(err, customerrors.ErrMediaNotFound) {
			return &gh.Response[any]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrMediaNotFound
		}
		return &gh.Response[any]{}, fmt.Errorf("movieService.UpdateMovieByID: %w", err)
	}

	return &gh.Response[any]{
		StatusCode: http.StatusNoContent,
	}, nil
}

func (m *Movie) DeleteMovieByID(ctx context.Context, req *gh.Request[any]) (*gh.Response[any], error) {
	movieID := req.Params["id"]

	err := m.movieService.DeleteMovieByID(ctx, movieID)
	if err != nil {
		if errors.Is(err, customerrors.ErrMovieNotFound) {
			return &gh.Response[any]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrMovieNotFound
		}
		return &gh.Response[any]{}, fmt.Errorf("movieService.DeleteMovieByID: %w", err)
	}

	return &gh.Response[any]{
		StatusCode: http.StatusNoContent,
	}, nil
}
