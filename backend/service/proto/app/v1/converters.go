package appv1

import (
	appv1 "app/gen/buf/app/v1"
	"app/gen/sqlc/mysql"
	"app/library/idutil"
)

func ConvertTask(row mysql.Task) *appv1.Task {
	return &appv1.Task{
		Id:        idutil.Encode(ResourceNameTask, row.ID),
		Title:     row.Title,
		Status:    row.Status,
		UpdatedAt: row.UpdatedAt.Time().Unix(),
		CreatedAt: row.CreatedAt.Time().Unix(),
	}
}

func ConvertTasks(rows []mysql.Task) []*appv1.Task {
	tasks := make([]*appv1.Task, len(rows))

	for idx, row := range rows {
		tasks[idx] = ConvertTask(row)
	}

	return tasks
}
