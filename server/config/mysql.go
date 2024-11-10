package config

import (
	"fmt"

	"github.com/caarlos0/env/v8"
)

type Mysql struct {
	User     string `env:"MYSQL_USER_NAME" envDefault:"root"`
	PassWord string `env:"MYSQL_PASSWORD" envDefault:""`
	Host     string `env:"MYSQL_HOST" envDefault:"127.0.0.1"`
	Port     int    `env:"DB_PORT" envDefault:"3306"`
	Database string `env:"MYSQL_DB_NAME" envDefault:"game_server_example"`
}

func NewMysqlConfig() (*Mysql, error) {
	m := &Mysql{}
	if err := env.Parse(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (m Mysql) DNS() string {
	if m.PassWord == "" {
		return fmt.Sprintf(
			"%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
			m.User, m.Host, m.Port, m.Database,
		)
	}
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		m.User, m.PassWord, m.Host, m.Port, m.Database,
	)
}
