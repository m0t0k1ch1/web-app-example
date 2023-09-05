package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Container struct {
	MySQL *sql.DB
}

func NewContainer(conf Config) (*Container, error) {
	ctr := &Container{}
	{
		db, err := sql.Open("mysql", conf.MySQL.DSN())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to open mysql db: %s", conf.MySQL.DBName)
		}

		ctr.MySQL = db
	}

	return ctr, nil
}
