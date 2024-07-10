package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/samber/oops"

	"app/container"
	"app/domain/validation"
	"app/gen/gqlgen"
	"app/gen/sqlc/mysql"
	"app/library/gqlerrutil"
	"app/library/idutil"
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

	resourceName string
	idInDB       uint64
}

func (in *NodeServiceGetInput) Validate() error {
	if err := validation.Struct(in); err != nil {
		return err
	}

	resourceName, idInDB, err := idutil.Decode(in.ID)
	if err != nil {
		return oops.Errorf("invalid id")
	}

	in.resourceName = resourceName
	in.idInDB = idInDB

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

	switch in.resourceName {

	case idutil.ResourceNameTask:
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
		return NodeServiceGetOutput{}, gqlerrutil.NewBadUserInputError(ctx, oops.Errorf("unexpected resource name: %s", in.resourceName))
	}
}
