package service

import (
	"awesome-url-shortener/internal/config"
	"awesome-url-shortener/internal/models"
	"awesome-url-shortener/internal/store/keyval"
	"context"
)

// compile-time proofs of service interface implementation
var _ Service = (*service)(nil)

type service struct {
	store  keyval.KeyValueStore
	config config.Common
}

type Service interface {
	ShortUrl(ctx context.Context, req models.UrlShortPayload) models.UrlShortCreateResponse
	ResolveShortUrl(ctx context.Context, req models.UrlShortGetInput) models.UrlShortGetResponse
}

func NewService(store keyval.KeyValueStore, cfg config.Common) Service {
	return &service{
		store:  store,
		config: cfg,
	}
}
