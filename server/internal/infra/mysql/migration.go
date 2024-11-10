package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (db *ApplicationDB) Migrate() error {
	driver, err := mysql.WithInstance(db.Db.DB, &mysql.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql", driver)
	if err != nil {
		fmt.Println(err)
		fmt.Println("111111")
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
		return err
	}

	return nil

}
