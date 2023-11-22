package core

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	configloader "github.com/kayac/go-config"
	"github.com/m0t0k1ch1-go/timeutil/v3"
	"github.com/pkg/errors"

	"app/config"
	"app/container"
	"app/domain/validation"
	appv1 "app/service/proto/app/v1"
)

var (
	ProviderSet = wire.NewSet(
		provideValidator,
		provideAppConfig,

		provideClock,
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

func provideValidator() *validator.Validate {
	return validation.NewValidator()
}

func provideAppConfig(confPath ConfigPath, vldtr *validator.Validate) (config.AppConfig, error) {
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

func provideMySQLContainer(conf config.AppConfig) (*container.MySQLContainer, error) {
	mysqlCtr := &container.MySQLContainer{}
	{
		db, err := sql.Open("mysql", conf.MySQL.App.DSN())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to open mysql db: %s", conf.MySQL.App.Name)
		}

		mysqlCtr.App = db
	}

	return mysqlCtr, nil
}

func provideTaskService(vldtr *validator.Validate, clock timeutil.Clock, mysqlCtr *container.MySQLContainer) *appv1.TaskService {
	return appv1.NewTaskService(vldtr, clock, mysqlCtr)
}

func provideServer(conf config.AppConfig, taskService *appv1.TaskService) *Server {
	return NewServer(conf.Server, taskService)
}

func provideApp(conf config.AppConfig, srv *Server) *App {
	return NewApp(conf, srv)
}
