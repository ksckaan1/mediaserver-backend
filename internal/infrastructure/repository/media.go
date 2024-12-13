package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/enum/mediatype"
	"mediaserver/internal/domain/core/enum/storagetype"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/infrastructure/repository/sqlcgen"
)

func (m *Repository) CreateMedia(ctx context.Context, media *model.Media) error {
	err := m.queries.CreateMedia(ctx, sqlcgen.CreateMediaParams{
		ID:          media.ID,
		Path:        media.Path,
		Type:        media.Type.String(),
		StorageType: media.StorageType.String(),
		Size:        media.Size,
	})
	if err != nil {
		return fmt.Errorf("queries.CreateMovie: %w", err)
	}
	return nil
}

func (m *Repository) GetMediaByID(ctx context.Context, id string) (*model.Media, error) {
	media, err := m.queries.GetMediaByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("queries.GetMediaByID: %w", customerrors.ErrRecordNotFound)
		}
		return nil, fmt.Errorf("queries.GetMediaByID: %w", err)
	}
	return &model.Media{
		ID:          media.ID,
		CreatedAt:   media.CreatedAt,
		Path:        media.Path,
		Type:        mediatype.FromString(media.Type),
		StorageType: storagetype.FromString(media.StorageType),
		Size:        media.Size,
	}, nil
}

func (m *Repository) DeleteMediaByID(ctx context.Context, id string) error {
	_, err := m.queries.DeleteMediaByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("queries.DeleteMediaByID: %w", customerrors.ErrRecordNotFound)
		}
		return fmt.Errorf("queries.DeleteMediaByID: %w", err)
	}
	return nil
}
