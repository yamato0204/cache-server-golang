package cacherepository

import (
	"context"

	"github.com/yamato0204/cache-server-golang/internal/entity"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cachedb"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cachemodel"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/datamodel"
)

type CoinEntityCacheRepository struct {
	cacheDB *cachedb.CacheDB
}

func NewCoinEntityCacheRepository(cacheDB *cachedb.CacheDB) *CoinEntityCacheRepository {
	return &CoinEntityCacheRepository{cacheDB: cacheDB}
}

func (r *CoinEntityCacheRepository) Create(ctx context.Context, userId int64) (*entity.CoinEntity, error) {
	m := &datamodel.UserCoin{UserID: userId}
	err := r.cacheDB.Insert(ctx, cachemodel.NewUserCoinCacheModel(*m, cachedb.Insert))
	if err != nil {
		return nil, err
	}
	return entity.NewCoinEntity(m), nil
}

// 実装
func (r *CoinEntityCacheRepository) Save(ctx context.Context, e *entity.CoinEntity) error {
	m := datamodel.UserCoin{UserID: e.UserId(), Num: e.Num()}
	err := r.cacheDB.Update(ctx, cachemodel.NewUserCoinCacheModel(m, cachedb.Update))
	if err != nil {
		return err
	}
	return nil
}
