package testutil

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/m0t0k1ch1-go/sqlutil"
	"github.com/stretchr/testify/require"
)

func Equal(t *testing.T, expected any, actual any, opts ...cmp.Option) {
	t.Helper()

	if diff := cmp.Diff(expected, actual, opts...); len(diff) > 0 {
		t.Errorf("diff: %s", diff)
	}
}

func ExecSQL(t *testing.T, ctx context.Context, db *sql.DB, query string) {
	t.Helper()

	var f *os.File
	{
		var err error

		f, err = os.CreateTemp(t.TempDir(), "")
		require.Nil(t, err)

		_, err = f.WriteString(query)
		require.Nil(t, err)

		require.Nil(t, f.Close())
	}

	require.Nil(t, sqlutil.ExecFile(ctx, db, f.Name()))
}
