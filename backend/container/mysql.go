package container

import (
	"database/sql"
)

type MySQLContainer struct {
	App *sql.DB
}
