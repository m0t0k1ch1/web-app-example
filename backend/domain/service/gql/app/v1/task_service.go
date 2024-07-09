package appv1

import (
	"context"

	"github.com/m0t0k1ch1-go/timeutil/v4"
	"github.com/samber/oops"

	"app/container"
	"app/domain/service/gql"
	"app/domain/validation"
	"app/gen/gqlgen"
)

type TaskService struct {
	clock          timeutil.Clock
	mysqlContainer *container.MySQLContainer
}

func NewTaskService(
	clock timeutil.Clock,
	mysqlCtr *container.MySQLContainer,
) *TaskService {
	return &TaskService{
		clock:          clock,
		mysqlContainer: mysqlCtr,
	}
}

type TaskServiceListInput struct {
	Status *gqlgen.TaskStatus `validate:"" en:"status"`
	After  *string            `validate:"" en:"after"`
	First  int32              `validate:"gte=0,lte=100" en:"first"`
}

func (in *TaskServiceListInput) Validate() error {
	return validation.Struct(in)
}

type TaskServiceListOutput struct {
	TaskConnection *gqlgen.TaskConnection
}

func (s *TaskService) List(ctx context.Context, in TaskServiceListInput) (TaskServiceListOutput, error) {
	if err := in.Validate(); err != nil {
		return TaskServiceListOutput{}, gql.NewBadUserInputError(ctx, err)
	}

	// TODO

	return TaskServiceListOutput{}, nil
}

type TaskServiceCreateInput struct {
	Title string `validate:"min=1,max=32" en:"title"`
}

func (in *TaskServiceCreateInput) Validate() error {
	return validation.Struct(in)
}

type TaskServiceCreateOutput struct {
	Task *gqlgen.Task
}

func (s *TaskService) Create(ctx context.Context, in TaskServiceCreateInput) (TaskServiceCreateOutput, error) {
	if err := in.Validate(); err != nil {
		return TaskServiceCreateOutput{}, gql.NewBadUserInputError(ctx, err)
	}

	// TODO

	return TaskServiceCreateOutput{}, nil
}

type TaskServiceCompleteInput struct {
	ID string `validate:"required" en:"id"`

	id uint64
}

func (in *TaskServiceCompleteInput) Validate() error {
	if err := validation.Struct(in); err != nil {
		return err
	}

	id, err := DecodeTaskID(in.ID)
	if err != nil {
		return oops.Errorf("invalid id")
	}

	in.id = id

	return nil
}

type TaskServiceCompleteOutput struct {
	Task *gqlgen.Task
}

func (s *TaskService) Complete(ctx context.Context, in TaskServiceCompleteInput) (TaskServiceCompleteOutput, error) {
	if err := in.Validate(); err != nil {
		return TaskServiceCompleteOutput{}, gql.NewBadUserInputError(ctx, err)
	}

	// TODO

	return TaskServiceCompleteOutput{}, nil
}
