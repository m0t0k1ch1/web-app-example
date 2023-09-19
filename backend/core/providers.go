package core

import (
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	configloader "github.com/kayac/go-config"

	"app/config"
	appv1 "app/service/app/v1"
)

var (
	ProviderSet = wire.NewSet(
		provideAppConfig,

		provideMySQL,
		provideSentryHandler,
		provideTaskService,
		provideServer,

		provideApp,
	)
)

type ConfigPath string

func (confPath ConfigPath) String() string {
	return string(confPath)
}

type MySQL *sql.DB

func init() {
	configloader.Delims("<%", "%>")
}

func provideAppConfig(path ConfigPath) (config.AppConfig, error) {
	var conf config.AppConfig
	if err := configloader.LoadWithEnv(&conf, path.String()); err != nil {
		return config.AppConfig{}, errors.Wrap(err, "failed to load config")
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(conf); err != nil {
		return config.AppConfig{}, errors.Wrap(err, "invalid config")
	}

	return conf, nil
}

func provideMySQL(conf config.AppConfig) (MySQL, error) {
	db, err := sql.Open("mysql", conf.MySQL.DSN())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open mysql db: %s", conf.MySQL.DBName)
	}

	return MySQL(db), nil
}

func provideSentryHandler(conf config.AppConfig) (*sentryhttp.Handler, error) {
	if len(conf.Sentry.DSN) == 0 {
		return nil, nil
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         conf.Sentry.DSN,
		Environment: conf.Runtime.Env,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to initialize sentry sdk")
	}

	return sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	}), nil
}

func provideTaskService(mysql MySQL) *appv1.TaskService {
	return appv1.NewTaskService(mysql)
}

func provideServer(conf config.AppConfig, sentryHandler *sentryhttp.Handler, taskService *appv1.TaskService) *Server {
	return NewServer(conf.Server, sentryHandler, taskService)
}

func provideApp(conf config.AppConfig, srv *Server) *App {
	return NewApp(conf, srv)
}
