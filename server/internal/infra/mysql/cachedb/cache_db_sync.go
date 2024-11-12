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
