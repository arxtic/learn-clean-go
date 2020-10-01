package models

import "errors"

var (
	ErrInternalServerError = errors.New("Internal server error")
	ErrNotFound            = errors.New("Request not found")
	ErrBadParamInput       = errors.New("Invalid parameter")
)
