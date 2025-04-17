package port

import (
	"context"
	"io"
)

type Storage interface {
	Save(ctx context.Context, object *Object) (string, error)
	Delete(ctx context.Context, filePath string) error
	CreatePresignedURL(ctx context.Context, id string) (string, map[string]string, error)
}

type Object struct {
	Content   io.ReadSeeker
	Size      int64
	MimeType  string
	Extension string
}
