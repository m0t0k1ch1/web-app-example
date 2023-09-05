package env

import (
	"app/db"
)

type Container struct {
	DB *db.Container
}

func NewContainer(dbCtr *db.Container) *Container {
	return &Container{
		DB: dbCtr,
	}
}
