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
	qdb := mysql.New(h.env.DB)

	var id idutil.ID

	if err := rdbutil.Transact(ctx, h.env.DB, func(txCtx context.Context, tx *sql.Tx) error {
		qtx := mysql.New(tx)

		lastInsertID, err := qtx.CreateTask(txCtx, req.Msg.Title)
		if err != nil {
			return errors.Wrap(err, "failed to create task")
		}

		id = idutil.ID(lastInsertID)

		return nil
	}); err != nil {
		return nil, newUnknownError(err)
	}

	task, err := qdb.GetTask(ctx, id)
	if err != nil {
		return nil, newUnknownError(err)
	}

	return connect.NewResponse(&appv1.CreateTaskResponse{
		Task: newTask(task),
	}), nil
}
