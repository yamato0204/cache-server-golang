// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/yamato0204/cache-server-golang/cmd/di/provider"
	"github.com/yamato0204/cache-server-golang/config"
	"github.com/yamato0204/cache-server-golang/internal/handler"
	"github.com/yamato0204/cache-server-golang/internal/handler/api"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cachedb"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cacherepository"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/integrationrepository"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/repository"
	"github.com/yamato0204/cache-server-golang/internal/service"
	"github.com/yamato0204/cache-server-golang/internal/usecase"
)

// Injectors from wire.go:

func Inject() (*provider.GameServer, func(), error) {
	uuidGenerator := repository.NewUUIDGenerator()
	configMysql, err := config.NewMysqlConfig()
	if err != nil {
		return nil, nil, err
	}
	applicationDB, cleanup, err := mysql.NewDB(configMysql)
	if err != nil {
		return nil, nil, err
	}
	db := mysql.RetrieveSqlxDB(applicationDB)
	coinEntityRepository := repository.NewCoinEntityRepository(applicationDB)
	v := repository.NewModelBulkExecutorMap(coinEntityRepository)
	cacheDB := cachedb.NewCacheDB(db, v)
	coinEntityCacheRepository := cacherepository.NewCoinEntityCacheRepository(cacheDB)
	coinIntegrationRepository := integrationrepository.NewCoinIntegrationRepository(coinEntityCacheRepository)
	userRegisterService := service.NewUserRegisterService(uuidGenerator, coinIntegrationRepository)
	userUsecase := usecase.NewUserUsecase(userRegisterService)
	userHandler := api.NewUserHandler(userUsecase)
	baseHandler := handler.NewBaseHandler(userHandler)
	router := handler.NewRouter(baseHandler)
	gameServer := provider.NewGameServer(router, applicationDB)
	return gameServer, func() {
		cleanup()
	}, nil
}
