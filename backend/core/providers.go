package core

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	configloader "github.com/kayac/go-config"
	"github.com/m0t0k1ch1-go/timeutil/v3"
	"github.com/pkg/errors"

	"app/config"
	"app/container"
	appv1 "app/service/proto/app/v1"
)

var (
	ProviderSet = wire.NewSet(
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

func provideAppConfig(path ConfigPath) (config.AppConfig, error) {
	var conf config.AppConfig
	if err := configloader.LoadWithEnv(&conf, path.String()); err != nil {
		return config.AppConfig{}, errors.Wrap(err, "failed to load config")
	}

	if err := config.NewValidator().Struct(conf); err != nil {
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

func provideTaskService(clock timeutil.Clock, mysqlCtr *container.MySQLContainer) *appv1.TaskService {
	return appv1.NewTaskService(clock, mysqlCtr)
}

func provideServer(conf config.AppConfig, taskService *appv1.TaskService) *Server {
	return NewServer(conf.Server, taskService)
}

func provideApp(conf config.AppConfig, srv *Server) *App {
	return NewApp(conf, srv)
}
