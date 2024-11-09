package provider

import (
	"github.com/google/wire"
	"github.com/yamato0204/cache-server-golang/internal/service"
)

// userIdを提供するプロバイダー関数
func ProvideUserID() int64 {
	return 0 // 必要な値に置き換えてください
}

var ServiceSet = wire.NewSet(
	ProvideUserID,
	service.NewUserRegisterService,
)
