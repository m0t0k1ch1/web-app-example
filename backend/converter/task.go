package converter

import (
	appv1 "backend/gen/buf/app/v1"
	"backend/gen/sqlc/mysql"
)

func Task(row mysql.Task) *appv1.Task {
	return &appv1.Task{
		Id:        row.ID.Encode(),
		Title:     row.Title,
		Status:    appv1.TaskStatus(row.Status),
		UpdatedAt: row.UpdatedAt.Unix(),
		CreatedAt: row.CreatedAt.Unix(),
	}
}

func Tasks(rows []mysql.Task) []*appv1.Task {
	tasks := make([]*appv1.Task, len(rows))

	for idx, row := range rows {
		tasks[idx] = Task(row)
	}

	return tasks
}
