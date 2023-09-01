package core

import (
	"database/sql"

	"backend/config"
)

type Env struct {
	Config config.App
	DB     *sql.DB
}
