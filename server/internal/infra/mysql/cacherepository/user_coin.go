package cacherepository

import (
	"context"

	"github.com/yamato0204/cache-server-golang/internal/entity"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cachedb"
)

type CoinEntityCacheRepository struct {
	cacheDB *cachedb.CacheDB
}

func NewCoinEntityCacheRepository(cacheDB *cachedb.CacheDB) *CoinEntityCacheRepository {
	return &CoinEntityCacheRepository{cacheDB: cacheDB}
}

func (r *CoinEntityCacheRepository) Create(ctx context.Context, userId int64) (*entity.CoinEntity, error) {
	// 実装
	return nil, nil
}

func (r *CoinEntityCacheRepository) Save(ctx context.Context, e *entity.CoinEntity) error {
	// 実装
	return nil
}
