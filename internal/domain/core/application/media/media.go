package media

import (
	"context"
	"fmt"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/pkg/gh"
	"mediaserver/internal/port"
	"mime/multipart"
	"net/http"
)

type StorageService interface {
	Save(ctx context.Context, fh *multipart.FileHeader) (*model.FileInfo, error)
	Delete(ctx context.Context, mediaFilePath string) error
}

type MediaService interface {
	Create(ctx context.Context, fi *model.FileInfo) (string, error)
}

type Media struct {
	logger         port.Logger
	storageService StorageService
	mediaService   MediaService
}

func New(storageService StorageService, mediaService MediaService, logger port.Logger) (*Media, error) {
	return &Media{
		storageService: storageService,
		mediaService:   mediaService,
		logger:         logger,
	}, nil
}

type UploadMediaRequest struct {
	File *multipart.FileHeader `form:"file" validate:"required"`
}

type UploadMediaResponse struct {
	ID string `json:"id"`
}

func (m *Media) UploadMedia(ctx context.Context, req *gh.Request[UploadMediaRequest]) (*gh.Response[*UploadMediaResponse], error) {
	fileInfo, err := m.storageService.Save(ctx, req.Body.File)
	if err != nil {
		return &gh.Response[*UploadMediaResponse]{}, fmt.Errorf("storageService.Save: %w", err)
	}
	id, err := m.mediaService.Create(ctx, fileInfo)
	if err != nil {
		return &gh.Response[*UploadMediaResponse]{}, fmt.Errorf("mediaService.Create: %w", err)
	}
	return &gh.Response[*UploadMediaResponse]{
		Body: &UploadMediaResponse{
			ID: id,
		},
		StatusCode: http.StatusOK,
	}, nil
}
