package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	kayac_config "github.com/kayac/go-config"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	"backend/config"
)

var (
	confPath = flag.String("config", "app.yaml", "path to config file")
)

func loadConfig(path string) (config.App, error) {
	kayac_config.Delims("<%", "%>")

	var conf config.App
	if err := kayac_config.LoadWithEnv(&conf, path); err != nil {
		return config.App{}, err
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(conf); err != nil {
		return config.App{}, errors.Wrap(err, "invalid config")
	}

	return conf, nil
}

func main() {
	flag.Parse()

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	conf, err := loadConfig(*confPath)
	if err != nil {
		fatal(errors.Wrap(err, "failed to load config"))
	}

	app, err := NewApp(context.Background(), conf)
	if err != nil {
		fatal(errors.Wrap(err, "failed to initialize app"))
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := app.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fatal(errors.Wrap(err, "failed to start app"))
		}
		slog.Info("app started")
	}()

	<-ctx.Done()

	if err := app.Stop(context.Background()); err != nil {
		fatal(errors.Wrap(err, "failed to stop app"))
	}
	slog.Info("app stopped")
}

func fatal(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}
