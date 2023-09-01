package testutil

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"golang.org/x/exp/slog"
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
	mysqlContainer, err := setupMySQLContainer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup mysql container")
	}

	return func() {
		mysqlContainer.Terminate(ctx)
	}, nil
}

func setupMySQLContainer(ctx context.Context) (testcontainers.Container, error) {
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
		WaitingFor: wait.ForSQL(nat.Port("3306"), "mysql", func(host string, port nat.Port) string {
			return fmt.Sprintf(
				"root:@tcp(%s:%s)/test",
				host, port.Port(),
			)
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
