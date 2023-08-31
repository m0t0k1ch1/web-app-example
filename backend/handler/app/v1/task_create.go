package appv1

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	"backend/converter"
	appv1 "backend/gen/buf/app/v1"
	"backend/gen/sqlc/mysql"
	"backend/handler"
	"backend/library/idutil"
	"backend/library/rdbutil"
)

func (h *TaskServiceHandler) Create(ctx context.Context, req *connect.Request[appv1.TaskServiceCreateRequest]) (*connect.Response[appv1.TaskServiceCreateResponse], error) {
	var task mysql.Task

	if err := rdbutil.Transact(ctx, h.Env.DB, func(txCtx context.Context, tx *sql.Tx) (txErr error) {
		qtx := mysql.New(tx)

		var id64 int64
		if id64, txErr = qtx.CreateTask(txCtx, req.Msg.Title); txErr != nil {
			return handler.NewUnknownError(errors.Wrap(txErr, "failed to create task"))
		}

		if task, txErr = qtx.GetTask(txCtx, idutil.ID(id64)); txErr != nil {
			return handler.NewUnknownError(errors.Wrap(txErr, "failed to get task"))
		}

		return
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceCreateResponse{
		Task: converter.Task(task),
	}), nil
}
