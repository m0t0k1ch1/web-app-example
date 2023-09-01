package testutil

import (
	"context"
	"database/sql"
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
	mysqlCtr *mysqlContainer
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

func setup(ctx context.Context) (teardown func(), err error) {
	mysqlCtr, err = newMySQLContainer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup mysql container")
	}

	teardown = func() {
		mysqlCtr.Terminate(ctx)
	}

	return
}

func OpenDB(ctx context.Context) (*sql.DB, error) {
	port, err := mysqlCtr.MappedPort(ctx, "3306")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get mapped port")
	}

	return sql.Open("mysql", config.MySQL{
		Host:     "127.0.0.1",
		Port:     port.Int(),
		User:     "root",
		Password: "",
		DBName:   "test",
	}.DSN())
}

func fatal(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}

type mysqlContainer struct {
	testcontainers.Container
}

func newMySQLContainer(ctx context.Context) (*mysqlContainer, error) {
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
			"MYSQL_DATABASE":             "test",
		},
		WaitingFor: wait.ForSQL("3306", "mysql", func(host string, port nat.Port) string {
			slog.Info(host)

			return config.MySQL{
				Host:     host,
				Port:     port.Int(),
				User:     "root",
				Password: "",
				DBName:   "test",
			}.DSN()
		}),
	}

	ctr, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create container")
	}

	return &mysqlContainer{
		Container: ctr,
	}, nil
}
