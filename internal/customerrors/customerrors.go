package customerrors

import (
	"errors"
	"fmt"
)

var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrMovieNotFound       = errors.New("movie not found")
	ErrMediaNotFound       = errors.New("media not found")
	ErrUnsupportedFileType = errors.New("unsupported file type")
	ErrInvalidMediaType    = errors.New("invalid media type")
)

type ErrUnexpectedStatusCode struct {
	StatusCode int
}

func (e ErrUnexpectedStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d", e.StatusCode)
}
