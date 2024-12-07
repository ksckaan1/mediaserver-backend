package movierepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/infrastructure/movierepository/sqlcgen"
)

type MovieRepository struct {
	queries *sqlcgen.Queries
}

func New(db sqlcgen.DBTX) (*MovieRepository, error) {
	return &MovieRepository{
		queries: sqlcgen.New(db),
	}, nil
}

func (m *MovieRepository) CreateMovie(ctx context.Context, movie *model.Movie) error {
	err := m.queries.CreateMovie(ctx, sqlcgen.CreateMovieParams{
		ID:          movie.ID,
		CreatedAt:   movie.CreatedAt,
		Title:       movie.Title,
		TmdbID:      movie.TMDBID,
		Description: movie.Description,
	})
	if err != nil {
		return fmt.Errorf("queries.CreateMovie: %w", err)
	}
	return nil
}

func (m *MovieRepository) GetMovieByID(ctx context.Context, id string) (*model.Movie, error) {
	movie, err := m.queries.GetMovieByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("queries.GetMovieByID: %w", customerrors.ErrRecordNotFound)
		}
		return nil, fmt.Errorf("queries.GetMovieByID: %w", err)
	}
	return &model.Movie{
		ID:          movie.ID,
		CreatedAt:   movie.CreatedAt,
		Title:       movie.Title,
		TMDBID:      movie.TmdbID,
		Description: movie.Description,
	}, nil
}
