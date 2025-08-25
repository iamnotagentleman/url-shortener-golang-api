package keyval

import (
	"context"
)

type mockStore struct {
}

func (rs *mockStore) GetUrl(ctx context.Context, key string) (string, error) {
	return key, nil
}

func (rs *mockStore) SetUrl(ctx context.Context, key string, url string, ttl int) error {
	return nil
}

func NewMockStore() KeyValueStore {
	return &mockStore{}
}
