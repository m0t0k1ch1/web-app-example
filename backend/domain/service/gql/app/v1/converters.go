package appv1

import (
	"app/domain/model"
	"app/gen/gqlgen"
	"app/gen/sqlc/mysql"

	"github.com/samber/oops"
)

func ConvertIntoTask(taskInDB mysql.Task) *gqlgen.Task {
	return &gqlgen.Task{
		Id:     EncodeTaskID(taskInDB.ID),
		Title:  taskInDB.Title,
		Status: taskInDB.Status,
	}
}

func ConvertIntoTaskEdgesAndNodes(taskInDBs []mysql.Task, status *gqlgen.TaskStatus) ([]*gqlgen.TaskEdge, []*gqlgen.Task, error) {
	taskEdges := make([]*gqlgen.TaskEdge, len(taskInDBs))
	tasks := make([]*gqlgen.Task, len(taskInDBs))
	{
		for idx, taskInDB := range taskInDBs {
			task := ConvertIntoTask(taskInDB)

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
