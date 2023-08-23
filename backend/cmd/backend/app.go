package main

import (
	"context"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/m0t0k1ch1/web-app-sample/backend/gen/buf/app/v1/appv1connect"
	appv1 "github.com/m0t0k1ch1/web-app-sample/backend/handler/app/v1"
)

type App struct {
	server *http.Server
}

func NewApp(conf Config) *App {
	grpcMux := http.NewServeMux()
	grpcMux.Handle(appv1connect.NewAppServiceHandler(appv1.NewAppServiceHandler()))

	mux := http.NewServeMux()
	mux.Handle("/grpc/", http.StripPrefix("/grpc", grpcMux))

	return &App{
		server: &http.Server{
			Addr:    conf.Server.Addr(),
			Handler: h2c.NewHandler(mux, &http2.Server{}),
		},
	}
}

func (api *App) Start() error {
	return api.server.ListenAndServe()
}

func (api *App) Shutdown(ctx context.Context) error {
	return api.server.Shutdown(ctx)
}
