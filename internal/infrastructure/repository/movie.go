package repository

import (
	"cmp"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/infrastructure/repository/sqlcgen"

	"github.com/samber/lo"
)

func (m *Repository) CreateMovie(ctx context.Context, movie *model.Movie) error {
	err := m.queries.CreateMovie(ctx, sqlcgen.CreateMovieParams{
		ID:          movie.ID,
		Title:       movie.Title,
		TmdbID:      movie.TMDBID,
		Description: movie.Description,
	})
	if err != nil {
		return fmt.Errorf("queries.CreateMovie: %w", err)
	}
	return nil
}

func (m *Repository) GetMovieByID(ctx context.Context, id string) (*model.Movie, error) {
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
		UpdatedAt:   cmp.Or(movie.UpdatedAt.Time, movie.CreatedAt),
		Title:       movie.Title,
		TMDBID:      movie.TmdbID,
		Description: movie.Description,
	}, nil
}

func (m *Repository) ListMovies(ctx context.Context, limit, offset int64) (*model.MovieList, error) {
	count, err := m.queries.CountMovies(ctx)
	if err != nil {
		return nil, fmt.Errorf("queries.CountMovies: %w", err)
	}

	if count == 0 {
		return &model.MovieList{
			Movies: make([]*model.Movie, 0),
			Count:  count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}

	movies, err := m.queries.ListMovies(ctx, sqlcgen.ListMoviesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("queries.ListMovies: %w", err)
	}

	return &model.MovieList{
		Movies: lo.Map(movies, func(m sqlcgen.Movie, _ int) *model.Movie {
			return &model.Movie{
				ID:          m.ID,
				CreatedAt:   m.CreatedAt,
				UpdatedAt:   cmp.Or(m.UpdatedAt.Time, m.CreatedAt),
				Title:       m.Title,
				TMDBID:      m.TmdbID,
				Description: m.Description,
			}
		}),
		Count:  count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (m *Repository) UpdateMovieByID(ctx context.Context, movie *model.Movie) error {
	_, err := m.queries.UpdateMovieByID(ctx, sqlcgen.UpdateMovieByIDParams{
		ID:          movie.ID,
		Title:       movie.Title,
		TmdbID:      movie.TMDBID,
		Description: movie.Description,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("queries.UpdateMovieByID: %w", customerrors.ErrRecordNotFound)
		}
		return fmt.Errorf("queries.UpdateMovie: %w", err)
	}
	return nil
}

func (m *Repository) DeleteMovieByID(ctx context.Context, id string) error {
	_, err := m.queries.DeleteMovieByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("queries.DeleteMovieByID: %w", customerrors.ErrRecordNotFound)
		}
		return fmt.Errorf("queries.DeleteMovieByID: %w", err)
	}
	return nil
}
