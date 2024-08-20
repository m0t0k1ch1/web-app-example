package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/samber/oops"

	"app/container"
	"app/domain/nodeid"
	"app/gen/gqlgen"
	"app/gen/sqlc/mysql"
	"app/library/gqlerrutil"
)

type NodeService struct {
	mysqlContainer *container.MySQLContainer
}

func NewNodeService(
	mysqlCtr *container.MySQLContainer,
) *NodeService {
	return &NodeService{
		mysqlContainer: mysqlCtr,
	}
}

func (s *NodeService) Get(ctx context.Context, id string) (gqlgen.Node, error) {
	var (
		idInDB     uint64
		nodeIDType nodeid.Type
	)
	{
		{
			var err error

			idInDB, nodeIDType, err = nodeid.Decode(id)
			if err != nil {
				return nil, gqlerrutil.NewBadUserInputError(ctx, oops.Errorf("invalid id"))
			}
		}
	}

	qdb := mysql.New(s.mysqlContainer.App)

	switch nodeIDType {

	case nodeid.TypeTask:
		taskInDB, err := qdb.GetTask(ctx, idInDB)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, oops.Wrapf(err, "failed to get task")
		}

		return ConvertIntoTask(taskInDB), nil

	default:
		return nil, gqlerrutil.NewBadUserInputError(ctx, oops.Errorf("unexpected node id type: %s", nodeIDType))
	}
}
