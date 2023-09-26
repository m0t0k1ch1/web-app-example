package appv1

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"

	appv1 "app/gen/buf/app/v1"
	"app/gen/sqlc/mysql"
	"app/library/sqlutil"
	"app/service"
)

func (s *TaskService) Create(ctx context.Context, req *connect.Request[appv1.TaskServiceCreateRequest]) (*connect.Response[appv1.TaskServiceCreateResponse], error) {
	var task mysql.Task

	if err := sqlutil.Transact(ctx, s.mysql, func(txCtx context.Context, tx *sql.Tx) (txErr error) {
		qtx := mysql.New(tx)

		var id int64
		if id, txErr = qtx.CreateTask(txCtx, req.Msg.Title); txErr != nil {
			return service.NewUnknownError(errors.Wrap(txErr, "failed to create task"))
		}

		if task, txErr = qtx.GetTask(txCtx, uint64(id)); txErr != nil {
			return service.NewUnknownError(errors.Wrap(txErr, "failed to get task"))
		}

		return
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceCreateResponse{
		Task: NewTask(task),
	}), nil
}
