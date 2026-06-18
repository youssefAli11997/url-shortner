package service

import "context"

type MockURLService struct {
	EncodeFunc func(ctx context.Context, originalURL string) (string, error)
	DecodeFunc func(ctx context.Context, shortURL string) (string, error)
}

func (m *MockURLService) Encode(ctx context.Context, originalURL string) (string, error) {
	return m.EncodeFunc(ctx, originalURL)
}

func (m *MockURLService) Decode(ctx context.Context, shortURL string) (string, error) {
	return m.DecodeFunc(ctx, shortURL)
}
