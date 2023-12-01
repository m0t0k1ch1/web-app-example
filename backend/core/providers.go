package core

import (
	"database/sql"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	configloader "github.com/kayac/go-config"
	"github.com/m0t0k1ch1-go/timeutil/v3"
	"github.com/pkg/errors"

	"app/config"
	"app/container"
	"app/domain/log"
	appv1 "app/domain/service/proto/app/v1"
	"app/domain/validation"
)

var (
	ProviderSet = wire.NewSet(
		provideValidator,
		provideAppConfig,

		provideClock,
		provideLogger,
		provideMySQLContainer,
		provideTaskService,
		provideServer,

		provideApp,
	)
)

type ConfigPath string

func (confPath ConfigPath) String() string {
	return string(confPath)
}

func init() {
	configloader.Delims("<%", "%>")
}

func provideValidator() (*validation.Validator, error) {
	return validation.NewValidator()
}

func provideAppConfig(confPath ConfigPath, vldtr *validation.Validator) (config.AppConfig, error) {
	var conf config.AppConfig
	if err := configloader.LoadWithEnv(&conf, confPath.String()); err != nil {
		return config.AppConfig{}, errors.Wrap(err, "failed to load config")
	}

	if err := vldtr.Struct(conf); err != nil {
		return config.AppConfig{}, errors.Wrap(err, "invalid config")
	}

	return conf, nil
}

func provideClock() timeutil.Clock {
	return timeutil.NewClock()
}

func provideLogger() *slog.Logger {
	return log.NewLogger()
}

func provideMySQLContainer(conf config.AppConfig) (*container.MySQLContainer, error) {
	ctr := &container.MySQLContainer{}
	{
		db, err := sql.Open("mysql", conf.MySQL.App.DSN())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to open mysql db: %s", conf.MySQL.App.Name)
		}

		ctr.App = db
	}

	return ctr, nil
}

func provideTaskService(clock timeutil.Clock, mysqlCtr *container.MySQLContainer) *appv1.TaskService {
	return appv1.NewTaskService(clock, mysqlCtr)
}

func provideServer(conf config.AppConfig, logger *slog.Logger, taskService *appv1.TaskService) *Server {
	return NewServer(conf.Server, logger, taskService)
}

func provideApp(conf config.AppConfig, srv *Server) *App {
	return NewApp(conf, srv)
}
