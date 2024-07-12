package core

import (
	"context"
	"net/http"

	"app/config"
)

type App struct {
	Config config.AppConfig

	server *Server
}

func NewApp(conf config.AppConfig, srv *Server) *App {
	return &App{
		Config: conf,

		server: srv,
	}
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.server.Handler.ServeHTTP(w, r)
}

func (app *App) Start() error {
	return app.server.ListenAndServe()
}

func (app *App) Stop(ctx context.Context) error {
	return app.server.Shutdown(ctx)
}
