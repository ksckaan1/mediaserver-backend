package media

import (
	"context"
	"fmt"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/port"
)

type Repository interface {
	CreateMedia(ctx context.Context, media *model.Media) error
	GetMediaByID(ctx context.Context, id string) (*model.Media, error)
	ListMedias(ctx context.Context, limit, offset int64) (*model.MediaList, error)
	DeleteMediaByID(ctx context.Context, id string) error
}

type Media struct {
	repo   Repository
	idgen  port.IDGenerator
	logger port.Logger
}

func New(repo Repository, idgen port.IDGenerator, lg port.Logger) (*Media, error) {
	return &Media{
		repo:   repo,
		idgen:  idgen,
		logger: lg,
	}, nil
}

func (m *Media) Create(ctx context.Context, fi *model.FileInfo) (string, error) {
	err := m.repo.CreateMedia(ctx, &model.Media{
		ID:          fi.ID,
		Path:        fi.Path,
		Type:        fi.Type,
		StorageType: fi.StorageType,
		MimeType:    fi.MimeType,
		Size:        fi.Size,
	})
	if err != nil {
		err = fmt.Errorf("repo.CreateMedia: %w", err)
		m.logger.Error(ctx, "error when creating media",
			"error", err,
		)
		return "", err
	}
	m.logger.Info(ctx, "media created",
		"id", fi.ID,
	)
	return fi.ID, nil
}

func (m *Media) GetMediaByID(ctx context.Context, id string) (*model.Media, error) {
	media, err := m.repo.GetMediaByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("repo.GetMediaByID: %w", err)
		m.logger.Error(ctx, "error when getting media",
			"error", err,
		)
		return nil, err
	}
	m.logger.Info(ctx, "media retrieved",
		"id", media.ID,
	)
	return media, nil
}

func (m *Media) ListMedias(ctx context.Context, limit, offset int64) (*model.MediaList, error) {
	list, err := m.repo.ListMedias(ctx, limit, offset)
	if err != nil {
		err = fmt.Errorf("repo.ListMedias: %w", err)
		m.logger.Error(ctx, "error when listing medias",
			"error", err,
		)
		return nil, err
	}
	m.logger.Info(ctx, "medias listed",
		"count", list.Count,
		"limit", list.Limit,
		"offset", list.Offset,
	)
	return list, nil
}

func (m *Media) DeleteMediaByID(ctx context.Context, id string) error {
	err := m.repo.DeleteMediaByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("repo.DeleteMediaByID: %w", err)
		m.logger.Error(ctx, "error when deleting media",
			"error", err,
		)
		return err
	}
	m.logger.Info(ctx, "media deleted",
		"id", id,
	)
	return nil
}
