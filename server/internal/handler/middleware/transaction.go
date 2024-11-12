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

// func (m *TransactionMiddleware) Intercept(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) (err error) {
// 		var commonResp *schema.CommonResponse
// 		ctx := c.Request().Context()
// 		ctx =
