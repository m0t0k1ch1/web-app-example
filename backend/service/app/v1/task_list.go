package appv1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	appv1 "app/gen/buf/app/v1"
	"app/gen/sqlc/mysql"
	"app/service"
)

func (s *TaskService) List(ctx context.Context, req *connect.Request[appv1.TaskServiceListRequest]) (*connect.Response[appv1.TaskServiceListResponse], error) {
	qdb := mysql.New(s.Env.DB)

	tasks, err := qdb.ListTasks(ctx)
	if err != nil {
		return nil, service.NewUnknownError(errors.Wrap(err, "failed to list tasks"))
	}

	return connect.NewResponse(&appv1.TaskServiceListResponse{
		Tasks: NewTasks(tasks),
	}), nil
}
