package provider

import (
	"github.com/google/wire"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cachedb"
)

var InfraSet = wire.NewSet(
	mysql.NewDB,
	mysql.RetrieveSqlxDB,
	cachedb.NewCacheDB,
)
