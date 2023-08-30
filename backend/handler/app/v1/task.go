package appv1

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	appv1 "backend/gen/buf/app/v1"
	"backend/gen/sqlc/mysql"
	"backend/handler"
	"backend/library/idutil"
)

type TaskServiceHandler struct {
	env *handler.Env
}

func NewTaskServiceHandler(env *handler.Env) *TaskServiceHandler {
	return &TaskServiceHandler{
		env: env,
	}
}

func (h *TaskServiceHandler) mustGetTask(ctx context.Context, id idutil.ID) (mysql.Task, error) {
	task, err := mysql.New(h.env.DB).GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mysql.Task{}, newNotFoundError(errors.Wrap(err, "task not found"))
		}

		return mysql.Task{}, newUnknownError(errors.Wrap(err, "failed to get task"))
	}

	return task, nil
}

func newTask(row mysql.Task) *appv1.Task {
	return &appv1.Task{
		Id:        row.ID.Encode(),
		Title:     row.Title,
		Status:    appv1.TaskStatus(row.Status),
		UpdatedAt: row.UpdatedAt.Unix(),
		CreatedAt: row.CreatedAt.Unix(),
	}
}

func newTasks(rows []mysql.Task) []*appv1.Task {
	tasks := make([]*appv1.Task, len(rows))

	for idx, row := range rows {
		tasks[idx] = newTask(row)
	}

	return tasks
}

func newUnknownError(err error) error {
	return connect.NewError(connect.CodeUnknown, err)
}

func newInvalidArgumentError(err error) error {
	return connect.NewError(connect.CodeInvalidArgument, err)
}

func newNotFoundError(err error) error {
	return connect.NewError(connect.CodeNotFound, err)
}
