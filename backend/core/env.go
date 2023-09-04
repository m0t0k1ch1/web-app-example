package core

import (
	"database/sql"

	"app/config"
)

type Env struct {
	Config config.App
	DB     *sql.DB
}
