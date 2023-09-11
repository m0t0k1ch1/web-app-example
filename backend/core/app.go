package core

import (
	"context"
	"net/http"
)

type App struct {
	server *Server
}

func NewApp(srv *Server) *App {
	return &App{
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
