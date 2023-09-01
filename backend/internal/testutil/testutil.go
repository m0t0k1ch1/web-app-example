package testutil

import (
	"context"
	"os"
	"testing"

	"golang.org/x/exp/slog"
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
