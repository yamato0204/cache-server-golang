package cachedb

import (
	"context"
	"errors"

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

func (cdb *CacheDB) Insert(ctx context.Context, content CacheContent) error {
	cacheManager, err := extractDBOperationCacheManager(ctx)
	if err != nil {
		return err
	}

	cacheManager.mutex.Lock()
	defer cacheManager.mutex.Unlock()

	cachedContent, err := cdb.extractCacheContent(ctx, content)
	if err != nil {
		return err
	}
	if cachedContent != nil {
		content.SetCacheStatus(Insert)

		return nil
	}

	switch cachedContent.GetCacheStatus() {
	case Insert, Select, Update:
		return errors.New("already exists")
	case Delete:
		cachedContent.SetCacheStatus(Update)
		return cachedContent.Update(content)
	case None:
		cachedContent.SetCacheStatus(Insert)
		return cachedContent.Update(content)
	default:
		return errors.New("no status found")
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
