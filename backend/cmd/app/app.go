package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/m0t0k1ch1/web-app-sample/backend/gen/buf/app/v1/appv1connect"
	"github.com/m0t0k1ch1/web-app-sample/backend/gen/sqlc/mysql"
	"github.com/m0t0k1ch1/web-app-sample/backend/handler"
	appv1 "github.com/m0t0k1ch1/web-app-sample/backend/handler/app/v1"
)

type App struct {
	server *http.Server
}

func NewApp(ctx context.Context, conf Config) (*App, error) {
	var env *handler.Env
	{
		db, err := sql.Open("mysql", conf.MySQL.DSN())
		if err != nil {
			return nil, errors.Wrap(err, "failed to connect to MySQL")
		}

		env = &handler.Env{
			Queries: mysql.New(db),
		}
	}

	var srv *http.Server
	{
		grpcMux := http.NewServeMux()
		grpcMux.Handle(appv1connect.NewAppServiceHandler(appv1.NewAppServiceHandler(env)))

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

		srv = &http.Server{
			Addr:    conf.Server.Addr(),
			Handler: c.Handler(h2c.NewHandler(mux, &http2.Server{})),
		}
	}

	return &App{
		server: srv,
	}, nil
}

func (app *App) Start() error {
	return app.server.ListenAndServe()
}

func (app *App) Shutdown(ctx context.Context) error {
	return app.server.Shutdown(ctx)
}
