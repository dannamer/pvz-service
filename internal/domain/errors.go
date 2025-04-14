package domain

import "errors"

var (
	ErrInvalidEmailFormat     = errors.New("invalid email format")
	ErrInvalidRole            = errors.New("invalid role")

	ErrEmailAlreadyExists     = errors.New("email already registered")
	ErrInvalidCredentials     = errors.New("invalid email or password")

	ErrInternalServerError    = errors.New("internal server error")

	ErrModeratorAccessDenied  = errors.New("access denied: moderator only")
	ErrEmployeeAccessDenied   = errors.New("access denied: employee only")

	ErrReceptionAlreadyExists = errors.New("reception already exists: in progress")
	ErrReceptionAlreadyClosed = errors.New("reception already closed")

	ErrPvzNotFound            = errors.New("pvz not found")
)
