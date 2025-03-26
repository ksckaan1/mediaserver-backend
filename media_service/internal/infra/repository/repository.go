package repository

import (
	"common/enums/mediatype"
	"context"
	"media_service/internal/core/model"
	"media_service/internal/infra/repository/sqlcgen"

	"github.com/samber/lo"
)

type Repository struct {
	queries *sqlcgen.Queries
}

func New(db sqlcgen.DBTX) (*Repository, error) {
	return &Repository{
		queries: sqlcgen.New(db),
	}, nil
}

func (r *Repository) CreateMedia(ctx context.Context, media *model.Media) error {
	err := r.queries.CreateMedia(ctx, sqlcgen.CreateMediaParams{
		ID:       media.ID,
		Path:     media.Path,
		Type:     media.Type.String(),
		Title:    media.Title,
		MimeType: media.MimeType,
		Size:     int32(media.Size),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetMediaByID(ctx context.Context, id string) (*model.Media, error) {
	media, err := r.queries.GetMediaByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.Media{
		ID:        media.ID,
		CreatedAt: media.CreatedAt.Time,
		UpdatedAt: media.UpdatedAt.Time,
		Title:     media.Title,
		Path:      media.Path,
		Type:      mediatype.FromString(media.Type),
		MimeType:  media.MimeType,
		Size:      int64(media.Size),
	}, nil
}

func (r *Repository) ListMedias(ctx context.Context, limit, offset int64) (*model.MediaList, error) {
	count, err := r.queries.CountMedias(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &model.MediaList{
			List:   make([]*model.Media, 0),
			Count:  count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	medias, err := r.queries.ListMedias(ctx, sqlcgen.ListMediasParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return &model.MediaList{
		List: lo.Map(medias, func(m sqlcgen.Media, _ int) *model.Media {
			return &model.Media{
				ID:        m.ID,
				CreatedAt: m.CreatedAt.Time,
				UpdatedAt: m.UpdatedAt.Time,
				Title:     m.Title,
				Path:      m.Path,
				Type:      mediatype.FromString(m.Type),
				MimeType:  m.MimeType,
				Size:      int64(m.Size),
			}
		}),
		Count:  count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (r *Repository) UpdateMediaByID(ctx context.Context, media *model.Media) error {
	_, err := r.queries.UpdateMediaByID(ctx, sqlcgen.UpdateMediaByIDParams{
		ID:    media.ID,
		Title: media.Title,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteMediaByID(ctx context.Context, id string) error {
	_, err := r.queries.DeleteMediaByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// CreateMedia(ctx context.Context, media *model.Media) error
// 	GetMediaByID(ctx context.Context, id string) (*model.Media, error)
// 	ListMedias(ctx context.Context, limit, offset int64) (*model.MediaList, error)
// 	UpdateMediaByID(ctx context.Context, media *model.Media) error
// 	DeleteMediaByID(ctx context.Context, id string) error
