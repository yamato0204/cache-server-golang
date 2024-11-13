package repository

import (
	"github.com/yamato0204/cache-server-golang/internal/infra/mysql"
)

type CoinEntityRepository struct {
	db *mysql.ApplicationDB
}

func NewCoinEntityRepository(db *mysql.ApplicationDB) *CoinEntityRepository {
	return &CoinEntityRepository{db: db}
}
