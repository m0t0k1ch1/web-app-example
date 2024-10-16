package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/samber/oops"

	"app/core"
	_ "app/domain/log"
)

var (
	confPath = flag.String("config", "app.yaml", "path to an app config file")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	appConf, err := core.LoadAppConfig(*confPath)
	if err != nil {
		fatal(ctx, oops.Wrapf(err, "failed to load app config"))
	}

	app, err := core.InitializeApp(ctx, appConf)
	if err != nil {
		fatal(ctx, oops.Wrapf(err, "failed to initialize app"))
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigCh)

	appErrCh := make(chan error, 1)

	go func() {
		defer close(appErrCh)

		slog.InfoContext(ctx, "start app")
		appErrCh <- app.Start()
	}()

	select {
	case err := <-appErrCh:
		fatal(ctx, oops.Wrapf(err, "failed to start app"))
	case <-sigCh:
	}

	slog.InfoContext(ctx, "stop app")
	if err := app.Stop(ctx); err != nil {
		fatal(ctx, oops.Wrapf(err, "failed to stop app"))
	}

	<-appErrCh // wait http.ErrServerClosed
}

func fatal(ctx context.Context, err error) {
	slog.ErrorContext(ctx, err.Error())
	os.Exit(1)
}
