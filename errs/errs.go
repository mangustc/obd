package errs

import (
	"errors"
)

var (
	ErrInternalServer      = errors.New("Internal server error")
	ErrNotFound            = errors.New("Not found")
	ErrUnprocessableEntity = errors.New("Unprocessable Entity")
	ErrUnauthorized        = errors.New("Unauthorized")
)
