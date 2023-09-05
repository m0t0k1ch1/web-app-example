package testutil

import (
	"context"
	"path/filepath"

	"github.com/docker/go-connections/nat"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"app/db"
)

func SetupDB(ctx context.Context, schemaPath string) (db.Config, func(), error) {
	if !filepath.IsAbs(schemaPath) {
		return db.Config{}, nil, errors.New("schema path must be absolute")
	}

	mysqlCtr, mysqlConf, err := setupMySQL(ctx, schemaPath)
	if err != nil {
		return db.Config{}, nil, errors.Wrap(err, "failed to setup mysql")
	}

	return db.Config{
			MySQL: mysqlConf,
		}, func() {
			mysqlCtr.Terminate(ctx)
		}, nil
}

func setupMySQL(ctx context.Context, schemaPath string) (testcontainers.Container, db.MySQLConfig, error) {
	conf := db.MySQLConfig{
		User:   "root",
		DBName: "test",
	}

	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.0",
		ExposedPorts: []string{"3306/tcp"},
		Mounts: testcontainers.ContainerMounts{
			testcontainers.BindMount(filepath.Join(schemaPath, "sql"), "/docker-entrypoint-initdb.d"),
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
