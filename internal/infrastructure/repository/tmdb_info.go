package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/infrastructure/repository/sqlcgen"
)

func (m *Repository) GetTMDBInfoByID(ctx context.Context, id int64) (*model.TMDBInfo, error) {
	info, err := m.queries.GetTMDBInfo(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("queries.GetTMDBInfo: %w", customerrors.ErrRecordNotFound)
		}
		return nil, fmt.Errorf("queries.GetTMDBInfo: %w", err)
	}
	return &model.TMDBInfo{
		ID:            id,
		Title:         info.Title,
		OriginalTitle: info.OriginalTitle,
		PosterPath:    info.PosterPath,
		BackdropPath:  info.BackdropPath,
		VoteAverage:   info.VoteAverage,
		VoteCount:     info.VoteCount,
		Popularity:    info.Popularity,
		ReleaseDate:   info.ReleaseDate,
	}, nil
}

func (m *Repository) SetTMDBInfo(ctx context.Context, info *model.TMDBInfo) error {
	err := m.queries.SetTMDBInfo(ctx, sqlcgen.SetTMDBInfoParams{
		ID:            info.ID,
		Title:         info.Title,
		OriginalTitle: info.OriginalTitle,
		PosterPath:    info.PosterPath,
		BackdropPath:  info.BackdropPath,
		VoteAverage:   info.VoteAverage,
		VoteCount:     info.VoteCount,
		Popularity:    info.Popularity,
		ReleaseDate:   info.ReleaseDate,
	})
	if err != nil {
		return fmt.Errorf("queries.SetTMDBInfo: %w", err)
	}
	return nil
}
