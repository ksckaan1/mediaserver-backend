package movie

import (
	"context"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/port"
)

type Repository interface {
	CreateMovie(ctx context.Context, movie *model.Movie) error
	GetMovieByID(ctx context.Context, id string) (*model.Movie, error)
	ListMovies(ctx context.Context, limit, offset int64) (*model.MovieList, error)
	UpdateMovieByID(ctx context.Context, movie *model.Movie) error
	DeleteMovieByID(ctx context.Context, id string) error
}

type Movie struct {
	repo   Repository
	idgen  port.IDGenerator
	logger port.Logger
}

func New(repo Repository, idgen port.IDGenerator, lg port.Logger) (*Movie, error) {
	return &Movie{
		repo:   repo,
		idgen:  idgen,
		logger: lg,
	}, nil
}

func (m *Movie) CreateMovie(ctx context.Context, movie *model.Movie) (string, error) {
	if movie.Title == "" {
		err := customerrors.ErrInvalidTitle
		m.logger.Error(ctx, "error when creating movie",
			"error", err,
		)
		return "", err
	}

	movie.ID = m.idgen.NewID()

	err := m.repo.CreateMovie(ctx, movie)
	if err != nil {
		err = fmt.Errorf("repo.CreateMovie: %w", err)
		m.logger.Error(ctx, "error when creating movie",
			"error", err,
		)
		return "", err
	}
	m.logger.Info(ctx, "movie created",
		"id", movie.ID, "title", movie.Title,
	)
	return movie.ID, nil
}

func (m *Movie) GetMovieByID(ctx context.Context, id string) (*model.Movie, error) {
	movie, err := m.repo.GetMovieByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("repo.GetMovieByID: %w", err)
		m.logger.Error(ctx, "error when getting movie",
			"error", err,
			"id", id,
		)
		return nil, err
	}
	m.logger.Info(ctx, "movie retrieved",
		"id", movie.ID,
		"title", movie.Title,
	)
	return movie, nil
}

func (m *Movie) ListMovies(ctx context.Context, limit, offset int64) (*model.MovieList, error) {
	movies, err := m.repo.ListMovies(ctx, limit, offset)
	if err != nil {
		err = fmt.Errorf("repo.ListMovies: %w", err)
		m.logger.Error(ctx, "error when listing movies",
			"error", err,
		)
		return nil, err
	}
	m.logger.Info(ctx, "movies listed",
		"count", movies.Count,
		"limit", movies.Limit,
		"offset", movies.Offset,
	)
	return movies, nil
}

func (m *Movie) UpdateMovieByID(ctx context.Context, movie *model.Movie) error {
	err := m.repo.UpdateMovieByID(ctx, movie)
	if err != nil {
		err = fmt.Errorf("repo.UpdateMovieByID: %w", err)
		m.logger.Error(ctx, "error when updating movie",
			"error", err,
			"id", movie.ID,
		)
		return err
	}
	m.logger.Info(ctx, "movie updated",
		"id", movie.ID,
		"title", movie.Title,
	)
	return nil
}

func (m *Movie) DeleteMovieByID(ctx context.Context, id string) error {
	err := m.repo.DeleteMovieByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("repo.DeleteMovieByID: %w", err)
		m.logger.Error(ctx, "error when deleting movie",
			"error", err,
			"id", id,
		)
		return err
	}
	m.logger.Info(ctx, "movie deleted",
		"id", id,
	)
	return nil
}
