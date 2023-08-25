package appv1

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	appv1 "github.com/m0t0k1ch1/web-app-sample/backend/gen/buf/app/v1"
	"github.com/m0t0k1ch1/web-app-sample/backend/gen/sqlc/mysql"
	"github.com/m0t0k1ch1/web-app-sample/backend/handler"
	"github.com/m0t0k1ch1/web-app-sample/backend/library/idutil"
	"github.com/m0t0k1ch1/web-app-sample/backend/library/rdbutil"
)

type AppServiceHandler struct {
	env *handler.Env
}

func NewAppServiceHandler(env *handler.Env) *AppServiceHandler {
	return &AppServiceHandler{
		env: env,
	}
}

func (h *AppServiceHandler) mustGetTask(ctx context.Context, id idutil.ID) (mysql.Task, error) {
	task, err := mysql.New(h.env.DB).GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mysql.Task{}, newNotFoundError(errors.Wrap(err, "task not found"))
		}

		return mysql.Task{}, newUnknownError(errors.Wrap(err, "failed to get task"))
	}

	return task, nil
}

func (h *AppServiceHandler) updateTask(ctx context.Context, params mysql.UpdateTaskParams) (mysql.Task, error) {
	var task mysql.Task

	if err := rdbutil.Transact(ctx, h.env.DB, func(txCtx context.Context, tx *sql.Tx) (txErr error) {
		qtx := mysql.New(tx)

		if task, txErr = qtx.GetTaskForUpdate(txCtx, params.ID); txErr != nil {
			if errors.Is(txErr, sql.ErrNoRows) {
				return newNotFoundError(errors.Wrap(txErr, "task not found"))
			}

			return newUnknownError(errors.Wrap(txErr, "failed to get task for update"))
		}

		if txErr = qtx.UpdateTask(txCtx, params); txErr != nil {
			return newUnknownError(errors.Wrap(txErr, "failed to update task"))
		}

		if task, txErr = qtx.GetTask(txCtx, task.ID); txErr != nil {
			return newUnknownError(errors.Wrap(txErr, "failed to get task"))
		}

		return
	}); err != nil {
		return mysql.Task{}, err
	}

	return task, nil
}

func newTask(row mysql.Task) *appv1.Task {
	return &appv1.Task{
		Id:          row.ID.Encode(),
		Title:       row.Title,
		IsCompleted: row.IsCompleted,
		UpdatedAt:   row.UpdatedAt.Unix(),
		CreatedAt:   row.CreatedAt.Unix(),
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
