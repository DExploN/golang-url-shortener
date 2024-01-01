package storage

import "errors"

var (
	ErrURLExists    = errors.New("url exists")
	ErrURLNotExists = errors.New("url not exists")
)
