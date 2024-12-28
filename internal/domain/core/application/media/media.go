package media

import (
	"context"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
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
	GetMediaByID(ctx context.Context, id string) (*model.Media, error)
	ListMedias(ctx context.Context, limit, offset int64) (*model.MediaList, error)
	DeleteMediaByID(ctx context.Context, id string) error
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
		return &gh.Response[*UploadMediaResponse]{}, customerrors.ErrInternalServerError
	}
	id, err := m.mediaService.Create(ctx, fileInfo)
	if err != nil {
		return &gh.Response[*UploadMediaResponse]{}, customerrors.ErrInternalServerError
	}
	return &gh.Response[*UploadMediaResponse]{
		Body: &UploadMediaResponse{
			ID: id,
		},
		StatusCode: http.StatusOK,
	}, nil
}

func (m *Media) GetMediaByID(ctx context.Context, req *gh.Request[any]) (*gh.Response[*model.Media], error) {
	id := req.Params["id"]
	media, err := m.mediaService.GetMediaByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrMediaNotFound) {
			return &gh.Response[*model.Media]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrMediaNotFound
		}
		return &gh.Response[*model.Media]{}, customerrors.ErrInternalServerError
	}
	return &gh.Response[*model.Media]{
		Body:       media,
		StatusCode: http.StatusOK,
	}, nil
}

func (m *Media) ListMedias(ctx context.Context, req *gh.Request[any]) (*gh.Response[*model.MediaList], error) {
	limit, err := req.GetQueryInt64("limit", -1)
	if err != nil {
		return &gh.Response[*model.MediaList]{
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("req.GetQueryInt64 (limit): %w", err)
	}
	offset, err := req.GetQueryInt64("offset", 0)
	if err != nil {
		return &gh.Response[*model.MediaList]{
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("req.GetQueryInt64 (offset): %w", err)
	}
	medias, err := m.mediaService.ListMedias(ctx, limit, offset)
	if err != nil {
		return &gh.Response[*model.MediaList]{}, customerrors.ErrInternalServerError
	}
	return &gh.Response[*model.MediaList]{
		Body:       medias,
		StatusCode: http.StatusOK,
	}, nil
}

func (m *Media) DeleteMediaByID(ctx context.Context, req *gh.Request[any]) (*gh.Response[any], error) {
	id := req.Params["id"]
	media, err := m.mediaService.GetMediaByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrMediaNotFound) {
			return &gh.Response[any]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrMediaNotFound
		}
		return &gh.Response[any]{}, customerrors.ErrInternalServerError
	}
	err = m.mediaService.DeleteMediaByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrMediaNotFound) {
			return &gh.Response[any]{
				StatusCode: http.StatusNotFound,
			}, customerrors.ErrMediaNotFound
		}
		return &gh.Response[any]{}, customerrors.ErrInternalServerError
	}
	err = m.storageService.Delete(ctx, media.Path)
	if err != nil {
		return &gh.Response[any]{}, customerrors.ErrInternalServerError
	}
	return &gh.Response[any]{
		StatusCode: http.StatusNoContent,
	}, nil
}
