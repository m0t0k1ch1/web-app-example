package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	"github.com/m0t0k1ch1/web-app-sample/backend/config"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	conf := Config{
		Server: config.ServerConfig{
			Port: 8080,
		},
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	api := NewAPI(conf)

	go func() {
		if err := api.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fatal(errors.Wrap(err, "failed to start server"))
		}
	}()

	<-ctx.Done()

	if err := api.Shutdown(context.Background()); err != nil {
		fatal(errors.Wrap(err, "failed to shutdown server"))
	}
}

func fatal(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}
