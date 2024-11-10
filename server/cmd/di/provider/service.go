package provider

import (
	"github.com/google/wire"
	"github.com/yamato0204/cache-server-golang/internal/service"
)

var ServiceSet = wire.NewSet(

	service.NewUserRegisterService,
)
