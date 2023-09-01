package main

import (
	"context"
	"database/sql"
	"net/http"

	"connectrpc.com/connect"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"backend/core"
	"backend/gen/buf/app/v1/appv1connect"
	"backend/handler"
	appv1 "backend/handler/app/v1"
)

type App struct {
	config Config

	env    *core.Env
	server *http.Server
}

func NewApp(ctx context.Context, conf Config) (*App, error) {
	app := &App{
		config: conf,
	}

	if err := app.initEnv(); err != nil {
		return nil, errors.Wrap(err, "failed to initialize env")
	}

	app.initServer()

	return app, nil
}

func (app *App) initEnv() error {
	db, err := sql.Open("mysql", app.config.MySQL.DSN())
	if err != nil {
		return errors.Wrapf(err, "failed to connect to db: %s", app.config.MySQL.DBName)
	}

	app.env = &core.Env{
		DB: db,
	}

	return nil
}

func (app *App) initServer() {
	var grpcHandler http.Handler
	{
		r := chi.NewRouter()

		base := handler.NewHandlerBase(app.env)

		path, h := appv1connect.NewTaskServiceHandler(
			appv1.NewTaskServiceHandler(base),
			connect.WithInterceptors(ValidationInterceptor),
			connect.WithCodec(NewJSONCodec()),
		)

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

		r.Handle(path+"*", h2c.NewHandler(c.Handler(h), &http2.Server{}))

		grpcHandler = r
	}

	r := chi.NewRouter()
	r.Mount("/grpc", http.StripPrefix("/grpc", grpcHandler))

	app.server = &http.Server{
		Addr:    app.config.Server.Addr(),
		Handler: r,
	}
}

func (app *App) Start() error {
	return app.server.ListenAndServe()
}

func (app *App) Shutdown(ctx context.Context) error {
	return app.server.Shutdown(ctx)
}
