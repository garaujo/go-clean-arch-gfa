package domain

import "errors"

var (
	// ErrInternalServerError error message
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound error message
	ErrNotFound = errors.New("Your request item was not found")
	// ErrConflict error message
	ErrConflict = errors.New("Your item already exists")
	// ErrBadParamInput error message
	ErrBadParamInput = errors.New("The given param is not valid")
)
