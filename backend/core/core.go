package core

import (
	"database/sql"
)

type ConfigPath string

func (confPath ConfigPath) String() string {
	return string(confPath)
}

type MySQL *sql.DB
