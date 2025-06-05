package errors

import "errors"

var (
	ErrUnexpected = errors.New("unexpected error.")
	ErrNotFound   = errors.New("not found.")
)
