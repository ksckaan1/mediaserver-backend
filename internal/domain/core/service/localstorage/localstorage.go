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

type LocalStorage struct {
	cfg    *config.Config
	idgen  port.IDGenerator
	logger port.Logger
}

func New(cfg *config.Config, idgen port.IDGenerator, logger port.Logger) (*LocalStorage, error) {
	return &LocalStorage{
		cfg:    cfg,
		idgen:  idgen,
		logger: logger,
	}, nil
}

func (l *LocalStorage) Save(ctx context.Context, fh *multipart.FileHeader) (*model.FileInfo, error) {
	id := l.idgen.NewID()
	if fh == nil {
		err := fmt.Errorf("fh is nil")
		l.logger.Error(ctx, "error when saving file",
			"error", err,
		)
		return nil, err
	}
	f, err := fh.Open()
	if err != nil {
		return nil, fmt.Errorf("fh.Open: %w", err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			l.logger.Error(ctx, "error when closing multipart file header",
				"error", err,
			)
		}
	}()
	// get header for file info
	head := make([]byte, 261)
	f.Read(head)
	_, err = f.Seek(0, 0)
	if err != nil {
		return nil, fmt.Errorf("f.Seek: %w", err)
	}
	// get file info
	mediaType, err := l.getMediaType(head)
	if err != nil {
		return nil, fmt.Errorf("getMediaType: %w", err)
	}
	ext, err := l.getExtension(head)
	if err != nil {
		return nil, fmt.Errorf("getExtension: %w", err)
	}
	mimeType, err := l.getMimeType(head)
	if err != nil {
		return nil, fmt.Errorf("getMimeType: %w", err)
	}
	// write file
	fileName := fmt.Sprintf("%s.%s", id, ext)
	err = l.writeFile(ctx, f, fileName)
	if err != nil {
		return nil, fmt.Errorf("writeFile: %w", err)
	}
	return &model.FileInfo{
		ID:          id,
		Path:        fileName,
		Type:        mediaType,
		StorageType: storagetype.Local,
		MimeType:    mimeType,
		Size:        fh.Size,
	}, nil
}

func (l *LocalStorage) getMediaType(head []byte) (mediatype.MediaType, error) {
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

func (l *LocalStorage) getExtension(head []byte) (string, error) {
	match, err := filetype.Match(head)
	if err != nil {
		return "", fmt.Errorf("filetype.Match: %w", err)
	}
	return match.Extension, nil
}

func (l *LocalStorage) getMimeType(head []byte) (string, error) {
	match, err := filetype.Match(head)
	if err != nil {
		return "", fmt.Errorf("filetype.Match: %w", err)
	}
	return match.MIME.Value, nil
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

func (l *LocalStorage) Delete(ctx context.Context, mediaFilePath string) error {
	dst := filepath.Join(l.cfg.StoragePath, mediaFilePath)
	err := os.Remove(dst)
	if err != nil {
		return fmt.Errorf("os.Remove: %w", err)
	}
	l.logger.Info(ctx, "file deleted",
		"path", dst,
	)
	return nil
}
