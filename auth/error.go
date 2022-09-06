package auth

import "errors"

var (
	ErrInvalidAccessToken = errors.New("invalid auth token")
	ErrUserDoesNotExist   = errors.New("user does not exist")
	ErrUserAlreadyExists  = errors.New("user with such credentials already exist")
)
