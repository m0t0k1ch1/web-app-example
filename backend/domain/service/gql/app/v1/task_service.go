package appv1

import (
	"context"
	"database/sql"
	"errors"

	"github.com/m0t0k1ch1-go/sqlutil"
	"github.com/m0t0k1ch1-go/timeutil/v4"
	"github.com/samber/oops"

	"app/container"
	"app/domain/service/gql"
	"app/domain/validation"
	"app/gen/gqlgen"
	"app/gen/sqlc/mysql"
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

	var taskInDB mysql.Task
	{
		if err := sqlutil.Transact(ctx, s.mysqlContainer.App, func(txnCtx context.Context, txn *sql.Tx) (txnErr error) {
			qtxn := mysql.New(txn)

			now := s.clock.Now()

			var id int64
			if id, txnErr = qtxn.CreateTask(txnCtx, mysql.CreateTaskParams{
				Title:     in.Title,
				UpdatedAt: now,
				CreatedAt: now,
			}); txnErr != nil {
				return gql.NewInternalServerError(txnCtx, oops.Wrapf(txnErr, "failed to create task"))
			}

			if taskInDB, txnErr = qtxn.GetTask(txnCtx, uint64(id)); txnErr != nil {
				return gql.NewInternalServerError(txnCtx, oops.Wrapf(txnErr, "failed to get task"))
			}

			return
		}); err != nil {
			return TaskServiceCreateOutput{}, err
		}
	}

	return TaskServiceCreateOutput{
		Task: ConvertIntoTask(taskInDB),
	}, nil
}

type TaskServiceCompleteInput struct {
	ID string `validate:"required" en:"id"`

	idInDB uint64
}

func (in *TaskServiceCompleteInput) Validate() error {
	if err := validation.Struct(in); err != nil {
		return err
	}

	idInDB, err := DecodeTaskID(in.ID)
	if err != nil {
		return oops.Errorf("invalid id")
	}

	in.idInDB = idInDB

	return nil
}

type TaskServiceCompleteOutput struct {
	Task *gqlgen.Task
}

func (s *TaskService) Complete(ctx context.Context, in TaskServiceCompleteInput) (TaskServiceCompleteOutput, error) {
	if err := in.Validate(); err != nil {
		return TaskServiceCompleteOutput{}, gql.NewBadUserInputError(ctx, err)
	}

	qdb := mysql.New(s.mysqlContainer.App)

	taskInDB, err := qdb.GetTask(ctx, in.idInDB)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TaskServiceCompleteOutput{}, nil
		}

		return TaskServiceCompleteOutput{}, gql.NewInternalServerError(ctx, oops.Wrapf(err, "failed to get task"))
	}
	if taskInDB.Status == gqlgen.TaskStatusCompleted {
		return TaskServiceCompleteOutput{}, gql.NewBadUserInputError(ctx, oops.Errorf("task already completed"))
	}

	if err := sqlutil.Transact(ctx, s.mysqlContainer.App, func(txnCtx context.Context, txn *sql.Tx) (txnErr error) {
		qtxn := mysql.New(txn)

		if taskInDB, txnErr = qtxn.GetTaskForUpdate(txnCtx, taskInDB.ID); txnErr != nil {
			if errors.Is(err, sql.ErrNoRows) {
				taskInDB = mysql.Task{}

				return nil
			}

			return gql.NewInternalServerError(txnCtx, oops.Wrapf(txnErr, "failed to get task for update"))
		}
		if taskInDB.Status == gqlgen.TaskStatusCompleted {
			return gql.NewBadUserInputError(ctx, oops.Errorf("task already completed"))
		}

		now := s.clock.Now()

		if txnErr = qtxn.CompleteTask(txnCtx, mysql.CompleteTaskParams{
			ID:        taskInDB.ID,
			UpdatedAt: now,
		}); txnErr != nil {
			return gql.NewInternalServerError(txnCtx, oops.Wrapf(txnErr, "failed to complete task"))
		}

		if taskInDB, txnErr = qtxn.GetTask(txnCtx, taskInDB.ID); txnErr != nil {
			return gql.NewInternalServerError(txnCtx, oops.Wrapf(txnErr, "failed to get task"))
		}

		return
	}); err != nil {
		return TaskServiceCompleteOutput{}, err
	}

	if taskInDB.ID == 0 {
		return TaskServiceCompleteOutput{}, nil
	}

	return TaskServiceCompleteOutput{
		Task: ConvertIntoTask(taskInDB),
	}, nil
}
