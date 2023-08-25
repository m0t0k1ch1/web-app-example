package handler

import (
	"database/sql"
)

type Env struct {
	DB *sql.DB
}
