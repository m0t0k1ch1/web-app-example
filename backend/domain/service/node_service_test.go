package service_test

import (
	"context"
	"os"
	"testing"

	"github.com/m0t0k1ch1-go/sqlutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"app/domain/nodeid"
	"app/domain/service"
	"app/internal/testutil"
)

func setUpNodeService(t *testing.T, _ *gomock.Controller) (*service.NodeService, *Mocks) {
	t.Helper()

	return service.NewNodeService(
		mysqlCtr,
	), &Mocks{}
}

func TestNodeService(t *testing.T) {
	setup(t)
	t.Cleanup(func() {
		teardown(t)
	})

	ctx := context.Background()

	mockCtrl := gomock.NewController(t)

	var (
		task1ID = nodeid.Encode(1, nodeid.TypeTask)
	)
	{
		var f *os.File
		{
			var err error

			f, err = os.CreateTemp(t.TempDir(), "")
			require.Nil(t, err)

			_, err = f.WriteString(`
				INSERT INTO task (id, title, updated_at, created_at) VALUES (1, 'task1.title', 0, 0);
			`)
			require.Nil(t, err)

			require.Nil(t, f.Close())
		}

		require.Nil(t, sqlutil.ExecFile(ctx, mysqlCtr.App, f.Name()))
	}

	t.Run("success: get task1", func(t *testing.T) {
		{
			nodeService, _ := setUpNodeService(t, mockCtrl)

			node, err := nodeService.Get(ctx, task1ID)
			require.Nil(t, err)

			testutil.Equal(t, task1ID, node.GetId())
		}
	})
}
