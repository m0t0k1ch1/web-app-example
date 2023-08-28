package appv1

import (
	"context"

	"connectrpc.com/connect"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"github.com/pkg/errors"

	appv1 "github.com/m0t0k1ch1/web-app-sample/backend/gen/buf/app/v1"
	"github.com/m0t0k1ch1/web-app-sample/backend/gen/sqlc/mysql"
	"github.com/m0t0k1ch1/web-app-sample/backend/library/idutil"
)

func (h *TaskServiceHandler) Update(ctx context.Context, req *connect.Request[appv1.TaskServiceUpdateRequest]) (*connect.Response[appv1.TaskServiceUpdateResponse], error) {
	id, err := idutil.Decode(req.Msg.Id)
	if err != nil {
		return nil, newInvalidArgumentError(errors.Wrap(err, "invalid TaskServiceUpdateRequest.Id"))
	}

	taskBefore, err := h.mustGetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	params := mysql.UpdateTaskParams{
		ID:          taskBefore.ID,
		Title:       taskBefore.Title,
		IsCompleted: taskBefore.IsCompleted,
	}
	{
		mask, err := fieldmask_utils.MaskFromPaths(req.Msg.FieldMask.Paths, nil)
		if err != nil {
			return nil, newUnknownError(errors.Wrap(err, "failed to create mask"))
		}

		if err := fieldmask_utils.StructToStruct(mask, req.Msg, &params); err != nil {
			return nil, newUnknownError(errors.Wrap(err, "failed to apply mask"))
		}
	}

	taskAfter, err := h.updateTask(ctx, params)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceUpdateResponse{
		Task: newTask(taskAfter),
	}), nil
}
