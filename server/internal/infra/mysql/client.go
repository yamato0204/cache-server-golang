package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/yamato0204/cache-server-golang/config"
)

type ApplicationDB struct {
	Db *sqlx.DB
	Tx *sqlx.Tx
}

func RetrieveSqlxDB(db *ApplicationDB) *sqlx.DB {
	return db.Db
}

func NewDB(mysqlConfig *config.Mysql) (*ApplicationDB, func(), error) {
	db, err := sqlx.Open("mysql", mysqlConfig.DNS())
	if err != nil {
		return nil, func() {}, err
	}

	return &ApplicationDB{Db: db}

}
