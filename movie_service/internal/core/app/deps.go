package app

import (
	"context"
	"movie_service/internal/core/models"
)

type Repository interface {
	CreateMovie(ctx context.Context, movie *models.Movie) error
	GetMovieByID(ctx context.Context, id string) (*models.Movie, error)
	ListMovies(ctx context.Context, limit, offset int64) (*models.MovieList, error)
	UpdateMovieByID(ctx context.Context, movie *models.Movie) error
	DeleteMovieByID(ctx context.Context, id string) error
}
