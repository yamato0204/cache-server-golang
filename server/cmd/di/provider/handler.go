package provider

import (
	"github.com/google/wire"
	"github.com/yamato0204/cache-server-golang/internal/handler"
	"github.com/yamato0204/cache-server-golang/internal/handler/api"
)

var HandlerSet = wire.NewSet(
	handler.NewBaseHandler,
	handler.NewRouter,
	api.NewUserHandler,
)
