package appv1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	appv1 "github.com/m0t0k1ch1/web-app-sample/backend/gen/buf/app/v1"
	"github.com/m0t0k1ch1/web-app-sample/backend/gen/sqlc/mysql"
	"github.com/m0t0k1ch1/web-app-sample/backend/library/idutil"
)

func (h *AppServiceHandler) MakeTaskIncomplete(ctx context.Context, req *connect.Request[appv1.MakeTaskIncompleteRequest]) (*connect.Response[appv1.MakeTaskIncompleteResponse], error) {
	id, err := idutil.Decode(req.Msg.Id)
	if err != nil {
		return nil, newInvalidArgumentError(errors.Wrap(err, "failed to decode id"))
	}

	taskBefore, err := h.mustGetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	taskAfter, err := h.updateTask(ctx, mysql.UpdateTaskParams{
		ID:          id,
		Title:       taskBefore.Title,
		IsCompleted: false,
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.MakeTaskIncompleteResponse{
		Task: newTask(taskAfter),
	}), nil
}
