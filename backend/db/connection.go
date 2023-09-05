package db

import (
	"database/sql"

	"github.com/pkg/errors"

	"app/config"
)

type Connection struct {
	*sql.DB
}

func NewConnection(conf config.MySQL) (*Connection, error) {
	db, err := sql.Open("mysql", conf.DSN())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open db: %s", conf.DBName)
	}

	return &Connection{
		DB: db,
	}, nil
}
