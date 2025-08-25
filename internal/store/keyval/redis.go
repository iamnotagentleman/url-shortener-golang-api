package keyval

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisStore struct {
	rdb *redis.Client
}

func (rs *redisStore) GetUrl(ctx context.Context, key string) (string, error) {
	val, err := rs.rdb.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}

func (rs *redisStore) SetUrl(ctx context.Context, key string, url string, ttl int) error {
	err := rs.rdb.Set(ctx, key, url, time.Duration(ttl)*time.Second).Err()
	return err
}

func NewRedisStore(addr, password string, db int) KeyValueStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &redisStore{
		rdb: rdb,
	}
}
