package middleware

import (
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cachedb"
)

type TransactionMiddleware struct {
	db      *mysql.ApplicationDB
	cacheDB *cachedb.CacheDB
}

func NewTransactionMiddleware(db *mysql.ApplicationDB, cacheDB *cachedb.CacheDB) *TransactionMiddleware {
	return &TransactionMiddleware{db: db, cacheDB: cacheDB}
}
