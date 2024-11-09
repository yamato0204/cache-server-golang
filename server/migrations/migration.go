package migrations

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (db *DB) Migrate() error {
	driver, err := mysql.WithInstance(db.DB.DB, &mysql.Config{})
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

	tables, err := db.ListTables()
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Println(tables)

	return nil

}
