package env

import (
	"app/db"
)

type Container struct {
	DB *db.Connection
}

func NewContainer(db *db.Connection) *Container {
	return &Container{
		DB: db,
	}
}
