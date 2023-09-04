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

	"app/config"
	"app/core"
	"app/gen/buf/app/v1/appv1connect"
	"app/service"
	appv1 "app/service/app/v1"
)

type App struct {
	env *core.Env
	srv *http.Server
}

func NewApp(ctx context.Context, conf config.App) (*App, error) {
	env, err := newEnv(conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize env")
	}

	srv := newServer(env)

	return &App{
		env: env,
		srv: srv,
	}, nil
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.srv.Handler.ServeHTTP(w, r)
}

func (app *App) Start() error {
	return app.srv.ListenAndServe()
}

func (app *App) Stop(ctx context.Context) error {
	return app.srv.Shutdown(ctx)
}

func newEnv(conf config.App) (*core.Env, error) {
	db, err := sql.Open("mysql", conf.MySQL.DSN())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open db: %s", conf.MySQL.DBName)
	}

	return &core.Env{
		Config: conf,
		DB:     db,
	}, nil
}

func newServer(env *core.Env) *http.Server {
	var grpcHandler http.Handler
	{
		r := chi.NewRouter()

		base := service.NewBase(env)

		path, h := appv1connect.NewTaskServiceHandler(
			appv1.NewTaskService(base),
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

	return &http.Server{
		Addr:    env.Config.Server.Addr(),
		Handler: r,
	}
}
