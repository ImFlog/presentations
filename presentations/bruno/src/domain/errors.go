package domain

import "errors"

var (
	ErrInvalidLogin    = errors.New("bad login password")
	ErrProductNotFound = errors.New("product not found in catalogue")
)
