package testutil

import (
	"context"
	"database/sql"
	"path/filepath"

	"github.com/docker/go-connections/nat"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"app/config"
)

func SetupMySQL(ctx context.Context, dbName string, schemaPath string) (*sql.DB, func(), error) {
	if !filepath.IsAbs(schemaPath) {
		return nil, nil, errors.New("schema path must be absolute")
	}

	conf := config.MySQLConfig{
		User:   "root",
		DBName: dbName,
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
		return nil, nil, errors.Wrap(err, "failed to create container")
	}

	db, err := sql.Open("mysql", conf.DSN())
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to open mysql db: %s", conf.DBName)
	}

	return db, func() {
		ctr.Terminate(ctx)
	}, nil
}
