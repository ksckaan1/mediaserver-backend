package movie

import (
	"context"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/port"
	"time"
)

type MovieRepository interface {
	CreateMovie(ctx context.Context, movie *model.Movie) error
	GetMovieByID(ctx context.Context, id string) (*model.Movie, error)
}

type Movie struct {
	movieRepository MovieRepository
	idgen           port.IDGenerator
	logger          port.Logger
}

func New(movieRepository MovieRepository, idgen port.IDGenerator, lg port.Logger) (*Movie, error) {
	return &Movie{
		movieRepository: movieRepository,
		idgen:           idgen,
		logger:          lg,
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
	movie.CreatedAt = time.Now()

	err := m.movieRepository.CreateMovie(ctx, movie)
	if err != nil {
		err = fmt.Errorf("movieRepository.CreateMovie: %w", err)
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
	movie, err := m.movieRepository.GetMovieByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("movieRepository.GetMovieByID: %w", err)
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
