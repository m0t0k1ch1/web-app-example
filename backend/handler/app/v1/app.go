package appv1

import (
	"context"
	"errors"

	"connectrpc.com/connect"

	appv1 "github.com/m0t0k1ch1/web-app-sample/backend/gen/buf/app/v1"
	"github.com/m0t0k1ch1/web-app-sample/backend/gen/sqlc/mysql"
	"github.com/m0t0k1ch1/web-app-sample/backend/handler"
)

type AppServiceHandler struct {
	env *handler.Env
}

func NewAppServiceHandler(env *handler.Env) *AppServiceHandler {
	return &AppServiceHandler{
		env: env,
	}
}

func newTask(task mysql.Task) *appv1.Task {
	return &appv1.Task{
		Id:          task.ID.Encode(),
		Title:       task.Title,
		IsCompleted: task.IsCompleted,
		UpdatedAt:   task.UpdatedAt.Unix(),
		CreatedAt:   task.CreatedAt.Unix(),
	}
}

func newUnknownError(err error) error {
	return connect.NewError(connect.CodeUnknown, err)
}

func newInvalidArgumentError(err error) error {
	return connect.NewError(connect.CodeInvalidArgument, err)
}

func newNotFoundError(err error) error {
	return connect.NewError(connect.CodeNotFound, err)
}

func (h *AppServiceHandler) ListTasks(context.Context, *connect.Request[appv1.ListTasksRequest]) (*connect.Response[appv1.ListTasksResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("app.v1.AppService.ListTasks is not implemented"))
}

func (h *AppServiceHandler) UpdateTask(context.Context, *connect.Request[appv1.UpdateTaskRequest]) (*connect.Response[appv1.UpdateTaskResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("app.v1.AppService.UpdateTask is not implemented"))
}

func (h *AppServiceHandler) DeleteTask(context.Context, *connect.Request[appv1.DeleteTaskRequest]) (*connect.Response[appv1.DeleteTaskResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("app.v1.AppService.DeleteTask is not implemented"))
}
