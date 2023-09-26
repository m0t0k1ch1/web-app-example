package appv1

import (
	"context"

	"connectrpc.com/connect"

	appv1 "app/gen/buf/app/v1"
	"app/service"
)

func (s *TaskService) Get(ctx context.Context, req *connect.Request[appv1.TaskServiceGetRequest]) (*connect.Response[appv1.TaskServiceGetResponse], error) {
	task, err := service.GetTaskOrError(ctx, s.mysql, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceGetResponse{
		Task: NewTask(task),
	}), nil
}
