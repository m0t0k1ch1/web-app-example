package appv1

import (
	"app/gen/gqlgen"
	"app/gen/sqlc/mysql"
)

func ConvertIntoTask(row mysql.Task) *gqlgen.Task {
	return &gqlgen.Task{
		Id:     EncodeTaskID(row.ID),
		Title:  row.Title,
		Status: row.Status,
	}
}
