package cachemodel

import (
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/cachedb"
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql/datamodel"
)

type UserCoinCacheModel struct {
	datamodel.UserCoin
	cacheStatus cachedb.CacheStatus
}

func NewUserCoinCacheModel(userCoin datamodel.UserCoin, cacheStatus cachedb.CacheStatus) *UserCoinCacheModel {
	return &UserCoinCacheModel{UserCoin: userCoin, cacheStatus: cacheStatus}
}

func (m *UserCoinCacheModel) Table() string {
	return "user_coin"
}
