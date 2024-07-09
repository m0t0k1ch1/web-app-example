package appv1

import (
	"context"

	"github.com/samber/oops"

	"app/container"
	"app/domain/service/gql"
	"app/domain/validation"
	"app/gen/gqlgen"
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
	id           uint64
}

func (in *NodeServiceGetInput) Validate() error {
	if err := validation.Struct(in); err != nil {
		return err
	}

	resourceName, id, err := idutil.Decode(in.ID)
	if err != nil {
		return oops.Errorf("invalid id")
	}

	in.resourceName = resourceName
	in.id = id

	return nil
}

type NodeServiceGetOutput struct {
	Node gqlgen.Node
}

func (s *NodeService) Get(ctx context.Context, in NodeServiceGetInput) (NodeServiceGetOutput, error) {
	if err := in.Validate(); err != nil {
		return NodeServiceGetOutput{}, gql.NewBadUserInputError(ctx, err)
	}

	switch in.resourceName {

	// TODO

	default:
		return NodeServiceGetOutput{}, gql.NewBadUserInputError(ctx, oops.Errorf("unexpected resource name: %s", in.resourceName))
	}
}
