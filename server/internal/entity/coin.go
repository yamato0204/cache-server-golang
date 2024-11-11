package entity

import "github.com/yamato0204/cache-server-golang/internal/infra/mysql/datamodel"

type CoinEntity struct {
	useId int64
	num   int64
}

func NewCoinEntity(m *datamodel.UserCoin) *CoinEntity {
	return &CoinEntity{
		useId: m.UserID,
		num:   m.Num,
	}
}

func (e *CoinEntity) Add(num int64) {
	e.num += num
}

func (e *CoinEntity) UserId() int64 {
	return e.useId
}

func (e *CoinEntity) Num() int64 {
	return e.num
}
