package core

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/m0t0k1ch1-go/timeutil/v4"
	"github.com/samber/oops"

	"app/config"
	"app/container"
)

var (
	ProviderSet = wire.NewSet(
		provideClock,
		provideMySQLContainer,
		provideServer,
		provideApp,
	)
)

func provideClock() timeutil.Clock {
	return timeutil.NewClock()
}

func provideMySQLContainer(conf config.AppConfig) (*container.MySQLContainer, error) {
	ctr := &container.MySQLContainer{}
	{
		db, err := sql.Open("mysql", conf.MySQL.App.DSN())
		if err != nil {
			return nil, oops.Wrapf(err, "failed to open mysql db: %s", conf.MySQL.App.Name)
		}

		ctr.App = db
	}

	return ctr, nil
}

func provideServer(
	conf config.AppConfig,
) *Server {
	return NewServer(
		conf.Server,
	)
}

func provideApp(conf config.AppConfig, srv *Server) *App {
	return NewApp(conf, srv)
}
