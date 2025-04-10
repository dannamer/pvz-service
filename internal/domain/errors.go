package domain

import "errors"

var (
	ErrInvalidEmailFormat  = errors.New("invalid email format")
	ErrInvalidRole         = errors.New("invalid role")
	ErrPasswordTooShort    = errors.New("password too short")
	ErrEmailAlreadyExists  = errors.New("email already registered")
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrInternalServerError = errors.New("internal server error")
)
