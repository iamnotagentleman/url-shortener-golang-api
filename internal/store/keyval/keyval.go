package keyval

import (
	"awesome-url-shortener/internal/config"
	"context"
)

type KeyValueStore interface {
	GetUrl(ctx context.Context, key string) (string, error)
	SetUrl(ctx context.Context, key, value string, ttl int) error
}

func NewKeyValueStore(cfg *config.EnvVars) KeyValueStore {
	switch cfg.Common.KeyValProvider {
	case "redis":
		return NewRedisStore(cfg.Redis.Address, cfg.Redis.Password, cfg.Redis.DB)
	case "mock":
		return NewMockStore()
	default:
		panic("unknown key val store provider")
	}
}
