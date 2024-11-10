//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/yamato0204/cache-server-golang/cmd/di/provider"
)

func Inject() (*provider.GameServer, func(), error) {
	wire.Build(
		provider.NewGameServer,
		provider.HandlerSet,
		provider.UsecaseSet,
		provider.ServiceSet,
		provider.RepositorySet,
		provider.InfraSet,
		provider.ConfigSet,
	)
	return nil, nil, nil
}
