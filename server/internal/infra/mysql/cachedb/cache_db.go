package cachedb

import (
	"context"

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

func (cdb *CacheDB) extractCacheContent(ctx context.Context, content CacheContent) (CacheContent, error) {
	cacheManager, err := extractDBOperationCacheManager(ctx)
	if err != nil {
		return nil, err
	}

	_, ok := cacheManager.dbOperationResult[content.Table()]
	if !ok {
		cacheManager.dbOperationResult[content.Table()] = make(map[string]CacheContent)
	}
	cachedContent, ok := cacheManager.dbOperationResult[content.Table()][content.UniqueKeyColumnValueStr()]
	if !ok {
		return nil, nil
	}
	return cachedContent, nil

}
