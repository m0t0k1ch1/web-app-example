package main

import (
	"context"
	"net/http"

	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders: []string{
			"Accept-Encoding",
			"Content-Encoding",
			"Content-Type",
			"Connect-Protocol-Version",
			"Connect-Timeout-Ms",
			"Connect-Accept-Encoding",  // Unused in web browsers, but added for future-proofing
			"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
			"Grpc-Timeout",             // Used for gRPC-web
			"X-Grpc-Web",               // Used for gRPC-web
			"X-User-Agent",             // Used for gRPC-web
		},
		ExposedHeaders: []string{
			"Content-Encoding",         // Unused in web browsers, but added for future-proofing
			"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
			"Grpc-Status",              // Required for gRPC-web
			"Grpc-Message",             // Required for gRPC-web
		},
	})

	return &App{
		server: &http.Server{
			Addr:    conf.Server.Addr(),
			Handler: c.Handler(h2c.NewHandler(mux, &http2.Server{})),
		},
	}
}

func (app *App) Start() error {
	return app.server.ListenAndServe()
}

func (app *App) Shutdown(ctx context.Context) error {
	return app.server.Shutdown(ctx)
}
