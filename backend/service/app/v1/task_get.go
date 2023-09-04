package appv1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	"app/converter"
	appv1 "app/gen/buf/app/v1"
	"app/library/idutil"
	"app/service"
)

func (s *TaskService) Get(ctx context.Context, req *connect.Request[appv1.TaskServiceGetRequest]) (*connect.Response[appv1.TaskServiceGetResponse], error) {
	id, err := idutil.Decode(req.Msg.Id)
	if err != nil {
		return nil, service.NewInvalidArgumentError(errors.Wrap(err, "invalid TaskServiceGetRequest.Id"))
	}

	task, err := service.GetTaskOrError(ctx, s.Env.DB, id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceGetResponse{
		Task: converter.Task(task),
	}), nil
}
