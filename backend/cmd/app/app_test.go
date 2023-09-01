package main

import (
	"context"
	"os"
	"testing"

	"github.com/pkg/errors"

	"backend/internal/testutil"
)

func TestMain(m *testing.M) {
	os.Exit(testutil.Run(m))
}

func TestApp(t *testing.T) {
	ctx := context.Background()

	_, teardown, err := testutil.SetupMySQL(ctx, "test")
	if err != nil {
		t.Fatal(errors.Wrap(err, "failed to setup mysql"))
	}
	defer teardown()
}
