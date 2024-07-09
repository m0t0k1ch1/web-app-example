package core

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/m0t0k1ch1-go/timeutil/v4"
	"github.com/samber/oops"

	"app/config"
	"app/container"
	appv1 "app/domain/service/gql/app/v1"
	"app/gql"
)

var (
	ProviderSet = wire.NewSet(
		provideClock,
		provideMySQLContainer,
		provideGQLResolver,

		provideTaskService,
		provideNodeService,

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
) *appv1.TaskService {
	return appv1.NewTaskService(
		clock,
		mysqlCtr,
	)
}

func provideNodeService(
	mysqlCtr *container.MySQLContainer,
) *appv1.NodeService {
	return appv1.NewNodeService(
		mysqlCtr,
	)
}

func provideGQLResolver(
	taskService *appv1.TaskService,
	nodeService *appv1.NodeService,
) *gql.Resolver {
	return gql.NewResolver(
		taskService,
		nodeService,
	)
}

func provideServer(
	conf config.AppConfig,
	gqlResolver *gql.Resolver,
) *Server {
	return NewServer(
		conf.Server,
		gqlResolver,
	)
}

func provideApp(conf config.AppConfig, srv *Server) *App {
	return NewApp(conf, srv)
}
