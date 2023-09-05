package db

import (
	"database/sql"

	"github.com/pkg/errors"
)

type Connection struct {
	MySQL *sql.DB
}

func NewConnection(conf Config) (*Connection, error) {
	conn := &Connection{}
	{
		db, err := sql.Open("mysql", conf.MySQL.DSN())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to open db: %s", conf.MySQL.DBName)
		}

		conn.MySQL = db
	}

	return conn, nil
}
