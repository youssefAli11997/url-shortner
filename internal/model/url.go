package model

import (
	"errors"
	"time"
)

type URL struct {
	ID          int64
	OriginalURL string
	ShortCode   string
	CreatedAt   time.Time
}

var (
	ErrURLNotFound = errors.New("url not found")
	ErrInvalidURL  = errors.New("invalid url")
)
