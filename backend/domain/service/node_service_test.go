package service_test

import (
	"context"
	"testing"

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
		testutil.ExecSQL(t, ctx, mysqlCtr.App, `
			INSERT INTO task (id, title, updated_at, created_at) VALUES (1, 'task1.title', 0, 0);
		`)
	}

	t.Run("success: get task1", func(t *testing.T) {
		{
			nodeService, _ := setUpNodeService(t, mockCtrl)

			node, err := nodeService.Get(ctx, task1ID)
			require.Nil(t, err)

			require.Equal(t, task1ID, node.GetId())
		}
	})
}
