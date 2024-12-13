package localstorage

import (
	"context"
	"fmt"
	"io"
	"mediaserver/config"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/enum/mediatype"
	"mediaserver/internal/domain/core/enum/storagetype"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/port"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
)

type Repository interface {
	CreateMedia(ctx context.Context, media *model.Media) error
}

type LocalStorage struct {
	repo   Repository
	cfg    *config.Config
	idgen  port.IDGenerator
	logger port.Logger
}

func New(repo Repository, cfg *config.Config, idgen port.IDGenerator, logger port.Logger) (*LocalStorage, error) {
	return &LocalStorage{
		repo:   repo,
		cfg:    cfg,
		idgen:  idgen,
		logger: logger,
	}, nil
}

func (l *LocalStorage) Save(ctx context.Context, fh *multipart.FileHeader) (string, error) {
	id := l.idgen.NewID()

	fileType, fileName, err := l.saveFile(ctx, fh, id)
	if err != nil {
		return "", err
	}

	media := &model.Media{
		ID:          id,
		Path:        fileName,
		Type:        fileType,
		StorageType: storagetype.Local,
		Size:        fh.Size,
	}

	err = l.repo.CreateMedia(ctx, media)
	if err != nil {
		err = fmt.Errorf("repo.CreateMedia: %w", err)
		l.logger.Error(ctx, "error when creating media",
			"error", err,
		)
		return "", err
	}

	return id, nil
}

// saveFile saves a file to the local storage and returns the media type and file name.
func (l *LocalStorage) saveFile(ctx context.Context, fh *multipart.FileHeader, id string) (mediatype.MediaType, string, error) {
	f, err := fh.Open()
	if err != nil {
		err = fmt.Errorf("fh.Open: %w", err)
		l.logger.Error(ctx, "error when opening multipart file header",
			"error", err,
		)
		return "", "", err
	}
	defer func() {
		err = f.Close()
		if err != nil {
			l.logger.Error(ctx, "error when closing multipart file header",
				"error", err,
			)
		}
	}()

	mediaType, err := l.getMediaType(f)
	if err != nil {
		err = fmt.Errorf("isFileSupported: %w", err)
		l.logger.Error(ctx, "error when checking if file is supported",
			"error", err,
		)
		return "", "", err
	}

	ext, err := l.getExtension(f)
	if err != nil {
		err = fmt.Errorf("getExtension: %w", err)
		l.logger.Error(ctx, "error when getting file extension",
			"error", err,
		)
		return "", "", err
	}

	fileName := fmt.Sprintf("%s.%s", id, ext)

	err = l.writeFile(ctx, f, fileName)
	if err != nil {
		err = fmt.Errorf("saveFile: %w", err)
		l.logger.Error(ctx, "error when saving file",
			"error", err,
		)
		return "", "", err
	}

	return mediaType, fileName, nil
}

func (l *LocalStorage) getMediaType(f multipart.File) (mediatype.MediaType, error) {
	head := make([]byte, 261)
	f.Read(head)
	_, err := f.Seek(0, 0)
	if err != nil {
		return "", fmt.Errorf("f.Seek: %w", err)
	}
	if filetype.IsImage(head) {
		return mediatype.Image, nil
	}
	if filetype.IsAudio(head) {
		return mediatype.Audio, nil
	}
	if filetype.IsVideo(head) {
		return mediatype.Video, nil
	}
	return "", customerrors.ErrUnsupportedFileType
}

func (l *LocalStorage) getExtension(f multipart.File) (string, error) {
	match, err := filetype.MatchReader(f)
	if err != nil {
		return "", fmt.Errorf("filetype.MatchReader: %w", err)
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return "", fmt.Errorf("f.Seek: %w", err)
	}
	return match.Extension, nil
}

func (l *LocalStorage) writeFile(ctx context.Context, f multipart.File, fileName string) error {
	savePath := filepath.Join(l.cfg.StoragePath, fileName)
	newFile, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer func() {
		err = newFile.Close()
		if err != nil {
			l.logger.Error(ctx, "error when closing file",
				"error", err,
			)
		}
	}()
	_, err = io.Copy(newFile, f)
	if err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	return nil
}
