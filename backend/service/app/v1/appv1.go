package appv1

import (
	"app/domain"
	appv1 "app/gen/buf/app/v1"
	"app/gen/sqlc/mysql"
	"app/library/idutil"
)

func NewTask(row mysql.Task) *appv1.Task {
	return &appv1.Task{
		Id:     idutil.Encode(domain.ResourceNameTask, row.ID),
		Title:  row.Title,
		Status: appv1.TaskStatus(row.Status),
	}
}

func NewTasks(rows []mysql.Task) []*appv1.Task {
	tasks := make([]*appv1.Task, len(rows))

	for idx, row := range rows {
		tasks[idx] = NewTask(row)
	}

	return tasks
}
