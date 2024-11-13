package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/yamato0204/cache-server-golang/internal/handler/acontext"
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

func (m *TransactionMiddleware) Intercept(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		ctx := c.Request().Context()
		ctx = acontext.WithAPIResponse(ctx)
		ctx = cachedb.WithCacheContext(ctx)
		c.SetRequest(c.Request().WithContext(ctx))

		defer func() {
			if err != nil {
				m.cacheDB.Rollback(ctx)
			}
		}()

		if err = next(c); err != nil {
			return err
		}

		if err = m.cacheDB.Commit(ctx); err != nil {
			return err
		}
		return nil
	}

}
