package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
)

var (
	confPath = flag.String("config", "app.yaml", "path to config file")
)

func main() {
	flag.Parse()

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	conf, err := LoadConfig(*confPath)
	if err != nil {
		fatal(errors.Wrap(err, "failed to load config"))
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := NewApp(conf)

	go func() {
		if err := app.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fatal(errors.Wrap(err, "failed to start app"))
		}
	}()

	<-ctx.Done()

	if err := app.Shutdown(context.Background()); err != nil {
		fatal(errors.Wrap(err, "failed to shutdown app"))
	}
}

func fatal(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}
