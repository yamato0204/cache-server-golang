package cachedb

import (
	"context"
	"errors"
)

func (cdb *CacheDB) Begin() error {
	tx, err := cdb.db.Beginx()
	if err != nil {
		return err
	}
	cdb.tx = tx
	return nil

}

func (cdb *CacheDB) Commit(ctx context.Context) error {
	_, err := extractDBOperationCacheManager(ctx)
	if err != nil {
		return err
	}
	if cdb.tx == nil {
		return errors.New("transaction is not started")
	}

	return cdb.tx.Commit()
}

func (cdb *CacheDB) Rollback(ctx context.Context) error {
	_, err := extractDBOperationCacheManager(ctx)
	if err != nil {
		return err
	}
	if cdb.tx == nil {
		return errors.New("transaction is not started")
	}

	return cdb.tx.Rollback()
}

func (cdb *CacheDB) SyncedToDB(ctx context.Context) (map[string]map[CacheStatus][]CacheContent, error) {
	if cdb.tx == nil {
		return nil, errors.New("transaction is not started")
	}

	cacheManager, err := extractDBOperationCacheManager(ctx)
	if err != nil {
		return nil, err
	}

	cachesMapByStatus := make(map[string]map[CacheStatus][]CacheContent)

	for table, resultMap := range cacheManager.dbOperationResult {
		cachesMapByStatus[table] = make(map[CacheStatus][]CacheContent)
		for _, result := range resultMap {
			cacheStatus := result.GetCacheStatus()
			if cacheStatus == Select || cacheStatus == None {
				continue
			}

			if cachesMapByStatus[table][cacheStatus] == nil {
				cachesMapByStatus[table][cacheStatus] = make([]CacheContent, 0)
			}
			cachesMapByStatus[table][cacheStatus] = append(cachesMapByStatus[table][cacheStatus], result.CreateCopy())

			result.SetCacheStatus(None)
		}
	}

	for table, cachesMap := range cachesMapByStatus {
		for cacheStatus, caches := range cachesMap {
			switch cacheStatus {
			case Insert:
				err = cdb.modelBulkExecutorMap[table].BulkInsert(ctx, cdb.tx, caches)
				if err != nil {
					return nil, err
				}
			case Update:
				err = cdb.modelBulkExecutorMap[table].BulkUpdate(ctx, cdb.tx, caches)
				if err != nil {
					return nil, err
				}
			case Delete:
				err = cdb.modelBulkExecutorMap[table].BulkDelete(ctx, cdb.tx, caches)
				if err != nil {
					return nil, err
				}

			}
		}
	}

	return cachesMapByStatus, nil
}
