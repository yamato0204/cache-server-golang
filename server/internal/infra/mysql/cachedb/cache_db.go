package cachedb

import (
	"github.com/jmoiron/sqlx"
)

// type BulkExecutor interface {
// 	BulkInsert(ctx context.Context, tx *sqlx.Tx, contents []CacheContent) error
// 	BulkUpdate(ctx context.Context, tx *sqlx.Tx, contents []CacheContent) error
// 	BulkDelete(ctx context.Context, tx *sqlx.Tx, contents []CacheContent) error
// }

type CacheDB struct {
	db *sqlx.DB
	//tx                   *sqlx.Tx
	//modelBulkExecutorMap map[string]BulkExecutor
}

func NewCacheDB(db *sqlx.DB) *CacheDB {
	return &CacheDB{
		db: db,
		//modelBulkExecutorMap: modelBulkExecutorMap,
	}
}
