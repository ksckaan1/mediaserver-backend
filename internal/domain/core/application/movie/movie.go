package movie

import (
	"context"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/pkg/generichandler"
	"net/http"
	"time"
)

type MovieService interface {
	CreateMovie(ctx context.Context, movie *model.Movie) (string, error)
	GetMovieByID(ctx context.Context, id string) (*model.Movie, error)
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
	Title       string `json:"title"`
	Description string `json:"description"`
	TMDBID      string `json:"tmdb_id"`
}

type CreateMovieResponse struct {
	ID string `json:"id"`
}

func (m *Movie) CreateMovie(ctx context.Context, req *generichandler.Request[*CreateMovieRequest]) (*generichandler.Response[*CreateMovieResponse], error) {
	id, err := m.movieService.CreateMovie(ctx, &model.Movie{
		Title:       req.Body.Title,
		Description: req.Body.Description,
		TMDBID:      req.Body.TMDBID,
	})
	if err != nil {
		return &generichandler.Response[*CreateMovieResponse]{}, fmt.Errorf("movieService.CreateMovie: %w", err)
	}

	return &generichandler.Response[*CreateMovieResponse]{
		Body: &CreateMovieResponse{
			ID: id,
		},
		StatusCode: http.StatusCreated,
	}, nil
}

type GetMovieByIDResponse struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TMDBID      string    `json:"tmdb_id"`
}

func (m *Movie) GetMovieByID(ctx context.Context, req *generichandler.Request[any]) (*generichandler.Response[*GetMovieByIDResponse], error) {
	id := req.Params["id"]
	movie, err := m.movieService.GetMovieByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrRecordNotFound) {
			return &generichandler.Response[*GetMovieByIDResponse]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrRecordNotFound
		}

		return &generichandler.Response[*GetMovieByIDResponse]{}, fmt.Errorf("movieService.GetMovieByID: %w", err)
	}
	return &generichandler.Response[*GetMovieByIDResponse]{
		Body: &GetMovieByIDResponse{
			ID:          id,
			CreatedAt:   movie.CreatedAt,
			Title:       movie.Title,
			Description: movie.Description,
			TMDBID:      movie.TMDBID,
		},
		StatusCode: http.StatusOK,
	}, nil
}
