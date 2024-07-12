package service

import (
	"app/domain/model"
	"app/gen/gqlgen"
	"app/gen/sqlc/mysql"
	"app/library/idutil"

	"github.com/samber/oops"
)

func convertIntoTask(taskInDB mysql.Task) *gqlgen.Task {
	return &gqlgen.Task{
		Id:     idutil.EncodeTaskID(taskInDB.ID),
		Title:  taskInDB.Title,
		Status: taskInDB.Status,
	}
}

func convertIntoTaskEdgesAndTasks(taskInDBs []mysql.Task, status *gqlgen.TaskStatus) ([]*gqlgen.TaskEdge, []*gqlgen.Task, error) {
	taskEdges := make([]*gqlgen.TaskEdge, len(taskInDBs))
	tasks := make([]*gqlgen.Task, len(taskInDBs))
	{
		for idx, taskInDB := range taskInDBs {
			task := convertIntoTask(taskInDB)

			taskCursor := model.PaginationCursor{
				ID:         task.Id,
				TaskStatus: status,
			}

			encodedTaskCursor, err := taskCursor.Encode()
			if err != nil {
				return nil, nil, oops.Wrapf(err, "failed to encode task cursor")
			}

			taskEdges[idx] = &gqlgen.TaskEdge{
				Cursor: encodedTaskCursor,
				Node:   task,
			}
			tasks[idx] = task
		}
	}

	return taskEdges, tasks, nil
}
