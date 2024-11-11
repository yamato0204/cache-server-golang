package cachemodel

import (
	"errors"
	"fmt"

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

func (m *UserCoinCacheModel) SetCacheStatus(status cachedb.CacheStatus) {
	m.cacheStatus = status
}

func (m *UserCoinCacheModel) GetCacheStatus() cachedb.CacheStatus {
	return m.cacheStatus
}

func (m *UserCoinCacheModel) UniqueKeyColumnValueStr() string {
	return fmt.Sprintf("user_id=%d", m.UserID)
}

func (m *UserCoinCacheModel) Update(content cachedb.CacheContent) error {
	model, ok := content.(*UserCoinCacheModel)
	if !ok {
		return errors.New("unexpected type")
	}
	m.UserID = model.UserID
	m.Num = model.Num
	return nil
}
