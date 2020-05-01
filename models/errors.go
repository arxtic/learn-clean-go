package models

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your Request is not found")
	ErrBadParamInput       = errors.New("Param is invalid")
)
