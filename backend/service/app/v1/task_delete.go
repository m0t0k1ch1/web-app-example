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

func (s *TaskService) Delete(ctx context.Context, req *connect.Request[appv1.TaskServiceDeleteRequest]) (*connect.Response[appv1.TaskServiceDeleteResponse], error) {
	task, err := service.GetTaskOrError(ctx, s.mysql, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	if err := sqlutil.Transact(ctx, s.mysql, func(txCtx context.Context, tx *sql.Tx) (txErr error) {
		qtx := mysql.New(tx)

		if task, txErr = qtx.GetTaskForUpdate(txCtx, task.ID); txErr != nil {
			if errors.Is(txErr, sql.ErrNoRows) {
				return service.NewNotFoundError(errors.Wrap(txErr, "task not found"))
			}

			return service.NewUnknownError(errors.Wrap(txErr, "failed to get task for update"))
		}

		if txErr = qtx.DeleteTask(txCtx, task.ID); txErr != nil {
			return service.NewUnknownError(errors.Wrap(txErr, "failed to delete task"))
		}

		return
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceDeleteResponse{}), nil
}
