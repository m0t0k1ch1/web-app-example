package core

import (
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	configloader "github.com/kayac/go-config"

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

type MySQL *sql.DB

func init() {
	configloader.Delims("<%", "%>")
}

func provideAppConfig(path ConfigPath) (AppConfig, error) {
	var conf AppConfig
	if err := configloader.LoadWithEnv(&conf, path.String()); err != nil {
		return AppConfig{}, errors.Wrap(err, "failed to load config")
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(conf); err != nil {
		return AppConfig{}, errors.Wrap(err, "invalid config")
	}

	return conf, nil
}

func provideMySQL(conf AppConfig) (MySQL, error) {
	db, err := sql.Open("mysql", conf.MySQL.DSN())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open mysql db: %s", conf.MySQL.DBName)
	}

	return MySQL(db), nil
}

func provideTaskService(mysql MySQL) *appv1.TaskService {
	return appv1.NewTaskService(mysql)
}

func provideServer(conf AppConfig, taskService *appv1.TaskService) *Server {
	return NewServer(conf.Server, taskService)
}

func provideApp(conf AppConfig, srv *Server) *App {
	return NewApp(conf, srv)
}
