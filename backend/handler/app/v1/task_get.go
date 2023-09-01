package appv1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	"backend/converter"
	appv1 "backend/gen/buf/app/v1"
	"backend/handler"
	"backend/library/idutil"
)

func (s *TaskService) Get(ctx context.Context, req *connect.Request[appv1.TaskServiceGetRequest]) (*connect.Response[appv1.TaskServiceGetResponse], error) {
	id, err := idutil.Decode(req.Msg.Id)
	if err != nil {
		return nil, handler.NewInvalidArgumentError(errors.Wrap(err, "invalid TaskServiceGetRequest.Id"))
	}

	task, err := s.MustGetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceGetResponse{
		Task: converter.Task(task),
	}), nil
}
