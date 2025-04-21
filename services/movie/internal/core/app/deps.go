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

type Searcher interface {
	AddDocument(ctx context.Context, collectionName string, doc any) error
	DeleteDocument(ctx context.Context, collectionName string, docID string) error
	UpdateDocument(ctx context.Context, collectionName string, id string, doc any) error
	Search(ctx context.Context, collectionName, query, queryBy string, limit, offset int, cache bool, v any) (int, error)
}
