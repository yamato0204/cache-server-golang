package provider

import (
	"github.com/google/wire"
	"github.com/yamato0204/cache-server-golang/internal/ifrepository"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cacherepository"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/integrationrepository"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/repository"
)

var RepositorySet = wire.NewSet(
	repository.NewCoinEntityRepository,
	cacherepository.NewCoinEntityCacheRepository,

	integrationrepository.NewCoinIntegrationRepository,
	wire.Bind(new(ifrepository.IFCoinEntityRepository), new(*integrationrepository.CoinIntegrationRepository)),
)
