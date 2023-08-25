package appv1

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	appv1 "github.com/m0t0k1ch1/web-app-sample/backend/gen/buf/app/v1"
	"github.com/m0t0k1ch1/web-app-sample/backend/gen/sqlc/mysql"
	"github.com/m0t0k1ch1/web-app-sample/backend/library/idutil"
)

func (h *AppServiceHandler) GetTask(ctx context.Context, req *connect.Request[appv1.GetTaskRequest]) (*connect.Response[appv1.GetTaskResponse], error) {
	qdb := mysql.New(h.env.DB)

	id, err := idutil.Decode(req.Msg.Id)
	if err != nil {
		return nil, newInvalidArgumentError(errors.Wrap(err, "failed to decode id"))
	}

	task, err := qdb.GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, newNotFoundError(errors.Wrap(err, "task not found"))
		}

		return nil, newUnknownError(errors.Wrap(err, "failed to get task"))
	}

	return connect.NewResponse(&appv1.GetTaskResponse{
		Task: newTask(task),
	}), nil
}
