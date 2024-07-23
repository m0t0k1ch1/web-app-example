package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/samber/oops"

	"app/container"
	"app/domain/nodeid"
	"app/domain/validation"
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

type NodeServiceGetInput struct {
	ID string `validate:"required" en:"id"`

	idInDB     uint64
	nodeIDType nodeid.Type
}

func (in *NodeServiceGetInput) Validate() error {
	if err := validation.Struct(in); err != nil {
		return err
	}

	idInDB, nodeIDType, err := nodeid.Decode(in.ID)
	if err != nil {
		return oops.Errorf("invalid id")
	}

	in.idInDB = idInDB
	in.nodeIDType = nodeIDType

	return nil
}

type NodeServiceGetOutput struct {
	Node gqlgen.Node
}

func (s *NodeService) Get(ctx context.Context, in NodeServiceGetInput) (NodeServiceGetOutput, error) {
	if err := in.Validate(); err != nil {
		return NodeServiceGetOutput{}, gqlerrutil.NewBadUserInputError(ctx, err)
	}

	qdb := mysql.New(s.mysqlContainer.App)

	switch in.nodeIDType {

	case nodeid.TypeTask:
		taskInDB, err := qdb.GetTask(ctx, in.idInDB)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return NodeServiceGetOutput{}, nil
			}

			return NodeServiceGetOutput{}, gqlerrutil.NewInternalServerError(ctx, oops.Wrapf(err, "failed to get task"))
		}

		return NodeServiceGetOutput{
			Node: convertIntoTask(taskInDB),
		}, nil

	default:
		return NodeServiceGetOutput{}, gqlerrutil.NewBadUserInputError(ctx, oops.Errorf("unexpected node id type: %s", in.nodeIDType))
	}
}
