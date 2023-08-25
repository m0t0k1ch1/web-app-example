package appv1

import (
	"context"
	"errors"

	"connectrpc.com/connect"

	appv1 "github.com/m0t0k1ch1/web-app-sample/backend/gen/buf/app/v1"
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

func (h *AppServiceHandler) GetTask(context.Context, *connect.Request[appv1.GetTaskRequest]) (*connect.Response[appv1.GetTaskResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("app.v1.AppService.GetTask is not implemented"))
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
