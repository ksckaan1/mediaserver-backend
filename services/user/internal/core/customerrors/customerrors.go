package customerrors

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUsernameAlreadyInUse = errors.New("username already in use")
	ErrInvalidUserType      = errors.New("invalid user type")
)
