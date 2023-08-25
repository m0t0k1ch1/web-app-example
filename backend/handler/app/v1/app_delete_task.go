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

func (h *AppServiceHandler) DeleteTask(ctx context.Context, req *connect.Request[appv1.DeleteTaskRequest]) (*connect.Response[appv1.DeleteTaskResponse], error) {
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

		if txErr = qtx.DeleteTask(txCtx, task.ID); txErr != nil {
			return newUnknownError(errors.Wrap(txErr, "failed to delete task"))
		}

		return
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.DeleteTaskResponse{}), nil
}
