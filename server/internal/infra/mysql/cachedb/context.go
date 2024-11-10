package cachedb

import (
	"context"
	"errors"
)

type cacheKey struct{}

func WithCacheContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, cacheKey{}, newDBOperationCacheManager())
}

func extractDBOperationCacheManager(ctx context.Context) (*DBOperationCacheManager, error) {
	v := ctx.Value(cacheKey{})
	if v == nil {
		return nil, errors.New("cache manager not found")
	}
	cache, ok := v.(*DBOperationCacheManager)
	if !ok {
		return nil, errors.New("cache manager not found")
	}
	return cache, nil
}
