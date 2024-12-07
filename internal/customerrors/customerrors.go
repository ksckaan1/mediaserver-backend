package customerrors

import "errors"

var (
	ErrInvalidTitle   = errors.New("invalid title")
	ErrRecordNotFound = errors.New("record not found")
)
