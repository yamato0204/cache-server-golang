package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cachedb"
)

type CoinEntityRepository struct {
	db *mysql.ApplicationDB
}

func NewCoinEntityRepository(db *mysql.ApplicationDB) *CoinEntityRepository {
	return &CoinEntityRepository{db: db}
}

func (r *CoinEntityRepository) BulkInsert(ctx context.Context, tx *sqlx.Tx, content []cachedb.CacheContent) error {

	return nil
}

func (r *CoinEntityRepository) BulkUpdate(ctx context.Context, tx *sqlx.Tx, content []cachedb.CacheContent) error {

	return nil
}

func (r *CoinEntityRepository) BulkDelete(ctx context.Context, tx *sqlx.Tx, content []cachedb.CacheContent) error {

	return nil
}
