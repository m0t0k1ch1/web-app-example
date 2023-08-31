package appv1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	"backend/converter"
	appv1 "backend/gen/buf/app/v1"
	"backend/gen/sqlc/mysql"
	"backend/handler"
)

func (h *TaskServiceHandler) List(ctx context.Context, req *connect.Request[appv1.TaskServiceListRequest]) (*connect.Response[appv1.TaskServiceListResponse], error) {
	qdb := mysql.New(h.Env.DB)

	tasks, err := qdb.ListTasks(ctx)
	if err != nil {
		return nil, handler.NewUnknownError(errors.Wrap(err, "failed to list tasks"))
	}

	return connect.NewResponse(&appv1.TaskServiceListResponse{
		Tasks: converter.Tasks(tasks),
	}), nil
}
