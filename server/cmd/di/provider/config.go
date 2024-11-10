package provider

import (
	"github.com/google/wire"
	"github.com/yamato0204/cache-server-golang/config"
)

var ConfigSet = wire.NewSet(
	config.NewMysqlConfig,
)
