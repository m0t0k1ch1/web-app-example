package appv1

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	appv1 "app/gen/buf/app/v1"
	"app/gen/sqlc/mysql"
	"app/handler"
	"app/library/idutil"
	"app/library/rdbutil"
)

func (s *TaskService) Delete(ctx context.Context, req *connect.Request[appv1.TaskServiceDeleteRequest]) (*connect.Response[appv1.TaskServiceDeleteResponse], error) {
	id, err := idutil.Decode(req.Msg.Id)
	if err != nil {
		return nil, handler.NewInvalidArgumentError(errors.Wrap(err, "invalid TaskServiceDeleteRequest.Id"))
	}

	task, err := s.MustGetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := rdbutil.Transact(ctx, s.Env.DB, func(txCtx context.Context, tx *sql.Tx) (txErr error) {
		qtx := mysql.New(tx)

		if task, txErr = qtx.GetTaskForUpdate(txCtx, task.ID); txErr != nil {
			if errors.Is(txErr, sql.ErrNoRows) {
				return handler.NewNotFoundError(errors.Wrap(txErr, "task not found"))
			}

			return handler.NewUnknownError(errors.Wrap(txErr, "failed to get task for update"))
		}

		if txErr = qtx.DeleteTask(txCtx, task.ID); txErr != nil {
			return handler.NewUnknownError(errors.Wrap(txErr, "failed to delete task"))
		}

		return
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceDeleteResponse{}), nil
}
