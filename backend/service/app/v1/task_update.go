package appv1

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	appv1 "app/gen/buf/app/v1"
	"app/gen/sqlc/mysql"
	"app/library/idutil"
	"app/library/rdbutil"
	"app/service"
)

func (s *TaskService) Update(ctx context.Context, req *connect.Request[appv1.TaskServiceUpdateRequest]) (*connect.Response[appv1.TaskServiceUpdateResponse], error) {
	id, err := idutil.Decode(req.Msg.Id)
	if err != nil {
		return nil, service.NewInvalidArgumentError(errors.Wrap(err, "invalid TaskServiceUpdateRequest.Id"))
	}

	task, err := service.GetTaskOrError(ctx, s.Env.DB.MySQL, id)
	if err != nil {
		return nil, err
	}

	if err := rdbutil.Transact(ctx, s.Env.DB.MySQL, func(txCtx context.Context, tx *sql.Tx) (txErr error) {
		qtx := mysql.New(tx)

		if task, txErr = qtx.GetTaskForUpdate(txCtx, task.ID); txErr != nil {
			if errors.Is(txErr, sql.ErrNoRows) {
				return service.NewNotFoundError(errors.Wrap(txErr, "task not found"))
			}

			return service.NewUnknownError(errors.Wrap(txErr, "failed to get task for update"))
		}

		if txErr = qtx.UpdateTask(txCtx, mysql.UpdateTaskParams{
			ID:     task.ID,
			Title:  req.Msg.Title,
			Status: int32(req.Msg.Status),
		}); txErr != nil {
			return service.NewUnknownError(errors.Wrap(txErr, "failed to update task"))
		}

		if task, txErr = qtx.GetTask(txCtx, task.ID); txErr != nil {
			return service.NewUnknownError(errors.Wrap(txErr, "failed to get task"))
		}

		return
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceUpdateResponse{
		Task: NewTask(task),
	}), nil
}