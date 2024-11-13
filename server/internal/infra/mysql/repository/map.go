package repository

import "github.com/yamato0204/cache-server-golang/internal/infra/mysql/cachedb"

func NewModelBulkExecutorMap(
	coinEntityRepository *CoinEntityRepository,

) map[string]cachedb.BulkExecutor {
	return map[string]cachedb.BulkExecutor{
		"user_coin": coinEntityRepository,
	}
}
