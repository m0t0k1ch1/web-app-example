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

	if _, err := testutil.OpenDB(ctx); err != nil {
		t.Fatal(errors.Wrap(err, "failed to open db"))
	}
}
