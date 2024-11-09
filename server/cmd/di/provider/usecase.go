package provider

import (
	"github.com/google/wire"
	"github.com/yamato0204/cache-server-golang/internal/usecase"
)

var UsecaseSet = wire.NewSet(
	usecase.NewUserUsecase,
)
