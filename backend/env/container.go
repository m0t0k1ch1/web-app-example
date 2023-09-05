package env

import (
	"app/db"
)

type Container struct {
	DB *db.Container
}

func NewContainer(db *db.Container) *Container {
	return &Container{
		DB: db,
	}
}
