package testutil

import (
	"context"
	"database/sql"

	"github.com/docker/go-connections/nat"
	_ "github.com/go-sql-driver/mysql"
	"github.com/samber/oops"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"app/config"
)

func SetupMySQL(ctx context.Context) (*sql.DB, func(), error) {
	conf := config.MySQLDBConfig{
		User: "root",
		Name: "test",
	}

	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.0",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ALLOW_EMPTY_PASSWORD": "yes",
			"MYSQL_DATABASE":             conf.Name,
		},
		WaitingFor: wait.ForSQL("3306", "mysql", func(host string, port nat.Port) string {
			conf.Host = host
			conf.Port = port.Int()

			return conf.DSN()
		}),
	}

	ctr, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, nil, oops.Wrapf(err, "failed to create container")
	}

	db, err := sql.Open("mysql", conf.DSN())
	if err != nil {
		return nil, nil, oops.Wrapf(err, "failed to open mysql db: %s", conf.Name)
	}

	return db, func() {
		ctr.Terminate(ctx)
	}, nil
}
