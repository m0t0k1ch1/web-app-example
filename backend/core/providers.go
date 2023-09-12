package core

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	configloader "github.com/kayac/go-config"
	"github.com/pkg/errors"

	"app/config"
	appv1 "app/service/app/v1"
)

var (
	ProviderSet = wire.NewSet(
		provideAppConfig,

		provideMySQL,
		provideTaskService,
		provideServer,

		provideApp,
	)
)

func provideAppConfig(path ConfigPath) (config.AppConfig, error) {
	configloader.Delims("<%", "%>")

	var conf config.AppConfig
	if err := configloader.LoadWithEnv(&conf, path.String()); err != nil {
		return config.AppConfig{}, err
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

func provideTaskService(mysql MySQL) *appv1.TaskService {
	return appv1.NewTaskService(mysql)
}

func provideServer(conf config.AppConfig, taskService *appv1.TaskService) *Server {
	return NewServer(conf.Server, taskService)
}

func provideApp(srv *Server) *App {
	return NewApp(srv)
}
