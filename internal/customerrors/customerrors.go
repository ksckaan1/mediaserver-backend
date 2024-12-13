package customerrors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidTitle        = errors.New("invalid title")
	ErrRecordNotFound      = errors.New("record not found")
	ErrUnsupportedFileType = errors.New("unsupported file type")
)

type ErrUnexpectedStatusCode struct {
	StatusCode int
}

func (e ErrUnexpectedStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d", e.StatusCode)
}
