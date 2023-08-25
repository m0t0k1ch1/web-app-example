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

func (h *AppServiceHandler) CreateTask(ctx context.Context, req *connect.Request[appv1.CreateTaskRequest]) (*connect.Response[appv1.CreateTaskResponse], error) {
	var task mysql.Task

	if err := rdbutil.Transact(ctx, h.env.DB, func(txCtx context.Context, tx *sql.Tx) (txErr error) {
		qtx := mysql.New(tx)

		var id64 int64
		if id64, txErr = qtx.CreateTask(txCtx, req.Msg.Title); txErr != nil {
			return newUnknownError(errors.Wrap(txErr, "failed to create task"))
		}

		if task, txErr = qtx.GetTask(txCtx, idutil.ID(id64)); txErr != nil {
			return newUnknownError(errors.Wrap(txErr, "failed to get task"))
		}

		return
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.CreateTaskResponse{
		Task: newTask(task),
	}), nil
}
