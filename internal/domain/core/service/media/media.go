package media

import (
	"context"
	"fmt"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/port"
)

type Repository interface {
	CreateMedia(ctx context.Context, media *model.Media) error
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
	id := m.idgen.NewID()
	err := m.repo.CreateMedia(ctx, &model.Media{
		ID:          id,
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
	return id, nil
}
