package media

import (
	"context"
	"fmt"
	"mediaserver/internal/pkg/gh"
	"mediaserver/internal/port"
	"mime/multipart"
	"net/http"
)

type StorageService interface {
	Save(ctx context.Context, fh *multipart.FileHeader) (string, error)
}

type Media struct {
	logger         port.Logger
	storageService StorageService
}

func New(storageService StorageService, logger port.Logger) (*Media, error) {
	return &Media{
		storageService: storageService,
		logger:         logger,
	}, nil
}

type UploadMediaRequest struct {
	File *multipart.FileHeader `form:"file"`
}

type UploadMediaResponse struct {
	ID string `json:"id"`
}

func (m *Media) UploadMedia(ctx context.Context, req *gh.Request[UploadMediaRequest]) (*gh.Response[*UploadMediaResponse], error) {
	id, err := m.storageService.Save(ctx, req.Body.File)
	if err != nil {
		return &gh.Response[*UploadMediaResponse]{}, fmt.Errorf("storageService.Save: %w", err)
	}

	return &gh.Response[*UploadMediaResponse]{
		Body: &UploadMediaResponse{
			ID: id,
		},
		StatusCode: http.StatusOK,
	}, nil
}
