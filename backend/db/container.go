package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Container struct {
	config Config

	MySQL *sql.DB
}

func NewContainer(conf Config) (*Container, error) {
	mysqlDB, err := sql.Open("mysql", conf.MySQL.DSN())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open mysql db: %s", conf.MySQL.DBName)
	}

	return &Container{
		config: conf,

		MySQL: mysqlDB,
	}, nil
}
