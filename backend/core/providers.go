package core

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/m0t0k1ch1-go/timeutil/v4"
	"github.com/samber/oops"

	"app/config"
	"app/container"
	"app/domain/resolver"
	"app/domain/service"
)

var (
	ProviderSet = wire.NewSet(
		provideClock,
		provideMySQLContainer,

		provideTaskService,
		provideNodeService,

		provideResolver,
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

func provideTaskService(
	clock timeutil.Clock,
	mysqlCtr *container.MySQLContainer,
) *service.TaskService {
	return service.NewTaskService(
		clock,
		mysqlCtr,
	)
}

func provideNodeService(
	mysqlCtr *container.MySQLContainer,
) *service.NodeService {
	return service.NewNodeService(
		mysqlCtr,
	)
}

func provideResolver(
	taskService *service.TaskService,
	nodeService *service.NodeService,
) *resolver.Resolver {
	return resolver.NewResolver(
		taskService,
		nodeService,
	)
}

func provideServer(
	conf config.AppConfig,
	resolver *resolver.Resolver,
) *Server {
	return NewServer(
		conf.Server,
		resolver,
	)
}

func provideApp(conf config.AppConfig, srv *Server) *App {
	return NewApp(conf, srv)
}
