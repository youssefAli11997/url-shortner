package service

import (
	"context"
)

type URLService interface {
	Encode(ctx context.Context, originalURL string) (string, error)
	Decode(ctx context.Context, shortURL string) (string, error)
}

func NewURLService() URLService {
	return &MockURLService{
		EncodeFunc: func(ctx context.Context, originalURL string) (string, error) {
			return "http://localhost:8080/abc123", nil
		},
		DecodeFunc: func(ctx context.Context, shortURL string) (string, error) {
			return "https://google.com", nil
		},
	}
}
