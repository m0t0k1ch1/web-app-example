package testutil

import (
	"github.com/testcontainers/testcontainers-go"
)

type MySQLContainer struct {
	testcontainers.Container
}
