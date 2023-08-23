package handler

import (
	"context"

	"connectrpc.com/connect"

	appv1 "github.com/m0t0k1ch1/web-app-sample/backend/gen/buf/app/v1"
)

func (h *Handler) Ping(ctx context.Context, req *connect.Request[appv1.PingRequest]) (*connect.Response[appv1.PingResponse], error) {
	return connect.NewResponse(&appv1.PingResponse{}), nil
}
