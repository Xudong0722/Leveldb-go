package utils

import "errors"

var (
	ErrKeyNotFound  = errors.New("The key was not found.")
	ErrTypeMismatch = errors.New("Type mismatch")
)
