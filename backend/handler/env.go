package handler

import (
	"github.com/m0t0k1ch1/web-app-sample/backend/gen/sqlc/mysql"
)

type Env struct {
	Queries *mysql.Queries
}
