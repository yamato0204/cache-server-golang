package provider

import (
	"github.com/google/wire"
	"github.com/yamato0204/cache-server-golang/internal/ifrepository"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cachedb"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/repository"
)

var InfraSet = wire.NewSet(
	repository.NewUUIDGenerator,
	wire.Bind(new(ifrepository.IFUUIDGenerator), new(*repository.UUIDGenerator)),
	mysql.NewDB,
	mysql.RetrieveSqlxDB,
	cachedb.NewCacheDB,
	repository.NewModelBulkExecutorMap,
)
