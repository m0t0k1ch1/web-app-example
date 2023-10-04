package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	"app/core"
)

var (
	confPath = flag.String("config", "app.yaml", "path to config file")
)

func main() {
	flag.Parse()

	app, err := core.InitializeApp(context.Background(), core.ConfigPath(*confPath))
	if err != nil {
		fatal(errors.Wrap(err, "failed to initialize app"))
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigCh)

	appErrCh := make(chan error, 1)

	go func() {
		defer close(appErrCh)

		info("start app")
		appErrCh <- app.Start()
	}()

	select {
	case err := <-appErrCh:
		fatal(errors.Wrap(err, "failed to start app"))
	case <-sigCh:
	}

	info("stop app")
	if err := app.Stop(context.Background()); err != nil {
		fatal(errors.Wrap(err, "failed to stop app"))
	}

	<-appErrCh // wait http.ErrServerClosed
}

func info(msg string) {
	slog.Info(msg)
}

func fatal(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}
