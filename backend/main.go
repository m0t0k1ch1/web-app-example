package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cockroachdb/errors"
	"golang.org/x/exp/slog"

	"app/config"
	"app/core"
)

var (
	confPath = flag.String("config", "app.yaml", "path to config file")
)

func main() {
	flag.Parse()

	app, err := core.InitializeApp(context.Background(), config.ConfigPath(*confPath))
	if err != nil {
		fatal(errors.Wrap(err, "failed to initialize app"))
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		info("start app")
		if err := app.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fatal(errors.Wrap(err, "failed to start app"))
		}
	}()

	<-ctx.Done()

	info("stop app")
	if err := app.Stop(context.Background()); err != nil {
		fatal(errors.Wrap(err, "failed to stop app"))
	}
}

func info(msg string) {
	slog.Info(msg)
}

func fatal(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}
