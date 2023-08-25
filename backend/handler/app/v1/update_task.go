package appv1

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	appv1 "github.com/m0t0k1ch1/web-app-sample/backend/gen/buf/app/v1"
	"github.com/m0t0k1ch1/web-app-sample/backend/gen/sqlc/mysql"
	"github.com/m0t0k1ch1/web-app-sample/backend/library/idutil"
	"github.com/m0t0k1ch1/web-app-sample/backend/library/rdbutil"
)

func (h *AppServiceHandler) UpdateTask(ctx context.Context, req *connect.Request[appv1.UpdateTaskRequest]) (*connect.Response[appv1.UpdateTaskResponse], error) {
	id, err := idutil.Decode(req.Msg.Id)
	if err != nil {
		return nil, newInvalidArgumentError(errors.Wrap(err, "failed to decode id"))
	}

	task, err := h.mustGetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := rdbutil.Transact(ctx, h.env.DB, func(txCtx context.Context, tx *sql.Tx) (txErr error) {
		qtx := mysql.New(tx)

		if task, txErr = qtx.GetTaskForUpdate(txCtx, id); txErr != nil {
			if errors.Is(txErr, sql.ErrNoRows) {
				return newNotFoundError(errors.Wrap(txErr, "task not found"))
			}

			return newUnknownError(errors.Wrap(txErr, "failed to get task for update"))
		}

		if txErr = qtx.UpdateTask(txCtx, mysql.UpdateTaskParams{
			ID:          task.ID,
			Title:       req.Msg.Title,
			IsCompleted: req.Msg.IsCompleted,
		}); txErr != nil {
			return newUnknownError(errors.Wrap(txErr, "failed to update task"))
		}

		if task, txErr = qtx.GetTask(txCtx, id); txErr != nil {
			return newUnknownError(errors.Wrap(txErr, "failed to get task"))
		}

		return
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.UpdateTaskResponse{
		Task: newTask(task),
	}), nil
}
