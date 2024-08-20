package service

import (
	"app/domain/nodeid"
	"app/gen/gqlgen"
	"app/gen/sqlc/mysql"
)

func ConvertIntoTask(taskInDB mysql.Task) *gqlgen.Task {
	return &gqlgen.Task{
		Id:     nodeid.Encode(taskInDB.ID, nodeid.TypeTask),
		Title:  taskInDB.Title,
		Status: taskInDB.Status,
	}
}
