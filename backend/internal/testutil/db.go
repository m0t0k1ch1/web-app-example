package testutil

import (
	"context"
	"path/filepath"

	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"app/db"
)

func SetupDB(ctx context.Context) (db.Config, func(), error) {
	mysqlCtr, mysqlConf, err := setupMySQL(ctx)
	if err != nil {
		return db.Config{}, nil, errors.Wrap(err, "failed to setup mysql")
	}

	return db.Config{
			MySQL: mysqlConf,
		}, func() {
			mysqlCtr.Terminate(ctx)
		}, nil
}

func setupMySQL(ctx context.Context) (testcontainers.Container, db.MySQLConfig, error) {
	conf := db.MySQLConfig{
		User:   "root",
		DBName: "test",
	}

	pathToBeMounted, err := filepath.Abs("../../_schema/sql/")
	if err != nil {
		return nil, db.MySQLConfig{}, errors.Wrap(err, "failed to prepare absolute path for dir to be mounted")
	}

	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.0",
		ExposedPorts: []string{"3306/tcp"},
		Mounts: testcontainers.ContainerMounts{
			testcontainers.BindMount(pathToBeMounted, "/docker-entrypoint-initdb.d/"),
		},
		Env: map[string]string{
			"MYSQL_ALLOW_EMPTY_PASSWORD": "yes",
			"MYSQL_DATABASE":             conf.DBName,
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
		return nil, db.MySQLConfig{}, errors.Wrap(err, "failed to create container")
	}

	return ctr, conf, nil
}
