package testutil

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/m0t0k1ch1-go/sqlutil"
	"github.com/stretchr/testify/require"
)

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
