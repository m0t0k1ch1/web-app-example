package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"

	"github.com/m0t0k1ch1/web-app-sample/backend/config"
)

func main() {
	conf := Config{
		Server: config.ServerConfig{
			Port:            8080,
			ShutdownTimeout: 5,
		},
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	api := NewAPI(conf)

	go func() {
		if err := api.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(errors.Wrap(err, "failed to start server"))
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), conf.Server.ShutdownTimeout)
	defer cancel()

	if err := api.Shutdown(ctx); err != nil {
		log.Fatal(errors.Wrap(err, "failed to shutdown server"))
	}
}
