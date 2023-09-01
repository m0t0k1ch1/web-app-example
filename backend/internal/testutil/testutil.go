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

var (
	mysqlContainer MySQLContainer
)

func Run(m *testing.M) int {
	ctx := context.Background()

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	teardown, err := setup(ctx)
	if err != nil {
		fatal(err)
	}
	defer teardown()

	return m.Run()
}

func setup(ctx context.Context) (func(), error) {
	{
		container, err := newMySQLContainer(ctx, "test")
		if err != nil {
			return nil, errors.Wrap(err, "failed to setup mysql container")
		}

		mysqlContainer = MySQLContainer{
			Container: container,
		}
	}

	return func() {
		mysqlContainer.Terminate(ctx)
	}, nil
}

func newMySQLContainer(ctx context.Context, dbName string) (testcontainers.Container, error) {
	pathToBeMounted, err := filepath.Abs("../../_schema/sql/")
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare absolute path for dir to be mounted")
	}

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
		WaitingFor: wait.ForSQL(nat.Port("3306"), "mysql", func(host string, port nat.Port) string {
			return config.MySQL{
				Host:     host,
				Port:     port.Int(),
				User:     "root",
				Password: "",
				DBName:   dbName,
			}.DSN()
		}),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create container")
	}

	return container, nil
}

func fatal(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}
