package appv1

import (
	appv1 "app/gen/buf/app/v1"
	"app/gen/sqlc/mysql"
	"app/library/idutil"
	"app/service"
)

func ConvertTask(row mysql.Task) *appv1.Task {
	return &appv1.Task{
		Id:     idutil.Encode(service.ResourceNameTask, row.ID),
		Title:  row.Title,
		Status: appv1.TaskStatus(appv1.TaskStatus_value[string(row.Status)]),
	}
}

func ConvertTasks(rows []mysql.Task) []*appv1.Task {
	tasks := make([]*appv1.Task, len(rows))

	for idx, row := range rows {
		tasks[idx] = ConvertTask(row)
	}

	return tasks
}
