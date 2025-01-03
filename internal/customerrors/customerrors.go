package customerrors

import (
	"errors"
	"fmt"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrRecordNotFound      = errors.New("record not found")
	ErrMovieNotFound       = errors.New("movie not found")
	ErrMediaNotFound       = errors.New("media not found")
	ErrSeriesNotFound      = errors.New("series not found")
	ErrSeasonNotFound      = errors.New("season not found")
	ErrEpisodeNotFound     = errors.New("episode not found")
	ErrUserNotFound        = errors.New("user not found")
	ErrOTPNotFound         = errors.New("otp not found")
	ErrUnsupportedFileType = errors.New("unsupported file type")
	ErrInvalidMediaType    = errors.New("invalid media type")
)

type ErrUnexpectedStatusCode struct {
	StatusCode int
}

func (e ErrUnexpectedStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d", e.StatusCode)
}
