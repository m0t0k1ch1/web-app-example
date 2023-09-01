package main

import (
	"os"
	"testing"

	"backend/internal/testutil"
)

func TestMain(m *testing.M) {
	os.Exit(testutil.Run(m))
}

func TestApp(*testing.T) {
	// TODO
}
