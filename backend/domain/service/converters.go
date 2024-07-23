package service

import (
	"github.com/samber/oops"

	"app/domain/model"
	"app/domain/nodeid"
	"app/gen/gqlgen"
	"app/gen/sqlc/mysql"
)

func convertIntoTask(taskInDB mysql.Task) *gqlgen.Task {
	return &gqlgen.Task{
		Id:     nodeid.Encode(taskInDB.ID, nodeid.TypeTask),
		Title:  taskInDB.Title,
		Status: taskInDB.Status,
	}
}

func convertIntoTaskEdges(taskInDBs []mysql.Task, paginationCursorParams model.PaginationCursorParams) ([]*gqlgen.TaskEdge, error) {
	taskEdges := make([]*gqlgen.TaskEdge, len(taskInDBs))
	{
		for idx, taskInDB := range taskInDBs {
			task := convertIntoTask(taskInDB)

			taskCursor := model.PaginationCursor{
				ID:     task.Id,
				Params: paginationCursorParams,
			}

			encodedTaskCursor, err := taskCursor.Encode()
			if err != nil {
				return nil, oops.Wrapf(err, "failed to encode task cursor")
			}

			taskEdges[idx] = &gqlgen.TaskEdge{
				Cursor: encodedTaskCursor,
				Node:   task,
			}
		}
	}

	return taskEdges, nil
}
