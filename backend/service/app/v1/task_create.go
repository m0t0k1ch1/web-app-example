package appv1

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"

	appv1 "app/gen/buf/app/v1"
	"app/gen/sqlc/mysql"
	"app/library/idutil"
	"app/library/sqlutil"
	"app/service"
)

func (s *TaskService) Create(ctx context.Context, req *connect.Request[appv1.TaskServiceCreateRequest]) (*connect.Response[appv1.TaskServiceCreateResponse], error) {
	var task mysql.Task

	if err := sqlutil.Transact(ctx, s.mysql, func(txCtx context.Context, tx *sql.Tx) (txErr error) {
		qtx := mysql.New(tx)

		var lastInsertID int64
		if lastInsertID, txErr = qtx.CreateTask(txCtx, req.Msg.Title); txErr != nil {
			return service.NewUnknownError(errors.Wrap(txErr, "failed to create task"))
		}

		id := uint64(lastInsertID)

		if txErr = qtx.UpdateTaskDisplayID(txCtx, mysql.UpdateTaskDisplayIDParams{
			ID: id,
			DisplayID: sql.NullString{
				String: idutil.Encode(id),
				Valid:  true,
			},
		}); txErr != nil {
			return service.NewUnknownError(errors.Wrap(txErr, "failed to update task.display_id"))
		}

		if task, txErr = qtx.GetTask(txCtx, id); txErr != nil {
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
