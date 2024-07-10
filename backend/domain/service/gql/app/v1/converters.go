package appv1

import (
	"app/gen/gqlgen"
	"app/gen/sqlc/mysql"
)

func ConvertIntoTask(taskInDB mysql.Task) *gqlgen.Task {
	return &gqlgen.Task{
		Id:     EncodeTaskID(taskInDB.ID),
		Title:  taskInDB.Title,
		Status: taskInDB.Status,
	}
}
