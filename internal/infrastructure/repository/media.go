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

	"github.com/samber/lo"
)

func (m *Repository) CreateMedia(ctx context.Context, media *model.Media) error {
	err := m.queries.CreateMedia(ctx, sqlcgen.CreateMediaParams{
		ID:          media.ID,
		Path:        media.Path,
		Type:        media.Type.String(),
		StorageType: media.StorageType.String(),
		MimeType:    media.MimeType,
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
			return nil, fmt.Errorf("queries.GetMediaByID: %w", customerrors.ErrMediaNotFound)
		}
		return nil, fmt.Errorf("queries.GetMediaByID: %w", err)
	}
	return &model.Media{
		ID:          media.ID,
		CreatedAt:   media.CreatedAt,
		Path:        media.Path,
		Type:        mediatype.FromString(media.Type),
		StorageType: storagetype.FromString(media.StorageType),
		MimeType:    media.MimeType,
		Size:        media.Size,
	}, nil
}

func (m *Repository) ListMedias(ctx context.Context, limit, offset int64) (*model.MediaList, error) {
	count, err := m.queries.CountMedias(ctx)
	if err != nil {
		return nil, fmt.Errorf("queries.CountMedias: %w", err)
	}
	if count == 0 {
		return &model.MediaList{
			List:   make([]*model.Media, 0),
			Count:  count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}

	list, err := m.queries.ListMedias(ctx, sqlcgen.ListMediasParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("queries.ListMedias: %w", err)
	}
	return &model.MediaList{
		List: lo.Map(list, func(m sqlcgen.Media, _ int) *model.Media {
			return &model.Media{
				ID:          m.ID,
				CreatedAt:   m.CreatedAt,
				Path:        m.Path,
				Type:        mediatype.FromString(m.Type),
				StorageType: storagetype.FromString(m.StorageType),
				MimeType:    m.MimeType,
				Size:        m.Size,
			}
		}),
		Count:  count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (m *Repository) DeleteMediaByID(ctx context.Context, id string) error {
	_, err := m.queries.DeleteMediaByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("queries.DeleteMediaByID: %w", customerrors.ErrMediaNotFound)
		}
		return fmt.Errorf("queries.DeleteMediaByID: %w", err)
	}
	return nil
}
