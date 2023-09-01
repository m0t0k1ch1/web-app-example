package testutil

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"golang.org/x/exp/slog"

	"backend/config"
)

func Run(m *testing.M) int {
	ctx := context.Background()

	teardown, err := setup(ctx)
	if err != nil {
		fatal(err)
	}
	defer teardown()

	return m.Run()
}

func setup(ctx context.Context) (teardown func(), err error) {
	return func() {}, nil
}

func fatal(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}

func SetupMySQL(ctx context.Context, dbName string) (config.MySQL, func(), error) {
	pathToBeMounted, err := filepath.Abs("../../_schema/sql/")
	if err != nil {
		return config.MySQL{}, nil, errors.Wrap(err, "failed to prepare absolute path for dir to be mounted")
	}

	var conf config.MySQL

	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.0",
		ExposedPorts: []string{"3306/tcp"},
		Mounts: testcontainers.ContainerMounts{
			testcontainers.BindMount(pathToBeMounted, "/docker-entrypoint-initdb.d/"),
		},
		Env: map[string]string{
			"MYSQL_ALLOW_EMPTY_PASSWORD": "yes",
			"MYSQL_DATABASE":             dbName,
		},
		WaitingFor: wait.ForSQL("3306", "mysql", func(host string, port nat.Port) string {
			conf = config.MySQL{
				Host:     host,
				Port:     port.Int(),
				User:     "root",
				Password: "",
				DBName:   dbName,
			}
			return conf.DSN()
		}),
	}

	ctr, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return config.MySQL{}, nil, errors.Wrap(err, "failed to create container")
	}

	return conf, func() {
		ctr.Terminate(ctx)
	}, nil
}
