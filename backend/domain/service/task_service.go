package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/go-cmp/cmp"
	"github.com/m0t0k1ch1-go/gqlutil"
	"github.com/m0t0k1ch1-go/sqlutil"
	"github.com/m0t0k1ch1-go/timeutil/v4"
	"github.com/samber/oops"

	"app/container"
	"app/domain/e"
	"app/domain/model"
	"app/domain/nodeid"
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

func (s *TaskService) List(ctx context.Context, status *gqlgen.TaskStatus, after *string, first int32) (*gqlgen.TaskConnection, error) {
	var (
		afterCursor model.PaginationCursor
	)
	{
		{
			if after != nil {
				var (
					err             error
					errInvalidAfter = oops.Errorf("invalid after")
				)

				if afterCursor, err = model.DecodePaginationCursor(*after); err != nil {
					return nil, e.NewGQLError(ctx, errInvalidAfter, gqlgen.ErrorCodeBadUserInput)
				}
				if !cmp.Equal(afterCursor.Params.TaskStatus, status) {
					return nil, e.NewGQLError(ctx, errInvalidAfter, gqlgen.ErrorCodeBadUserInput)
				}
			}
		}
		{
			if err := validation.Struct(struct {
				First int32 `validate:"gte=0,lte=100" en:"first"`
			}{
				First: first,
			}); err != nil {
				return nil, e.NewGQLError(ctx, err, gqlgen.ErrorCodeBadUserInput)
			}
		}
	}

	var listTasksParams mysql.ListTasksParams
	{
		if status != nil {
			listTasksParams.SetStatus = 1
			listTasksParams.Status = *status
		}

		listTasksParams.Limit = first + 1
		listTasksParams.Offset = afterCursor.Offset
	}

	var countTasksParams mysql.CountTasksParams
	{
		if status != nil {
			countTasksParams.SetStatus = 1
			countTasksParams.Status = *status
		}
	}

	qdb := mysql.New(s.mysqlContainer.App)

	var (
		edges       []*gqlgen.TaskEdge
		hasNextPage bool
	)
	{
		taskInDBs, err := qdb.ListTasks(ctx, listTasksParams)
		if err != nil {
			return nil, oops.Wrapf(err, "failed to list tasks")
		}

		hasNextPage = len(taskInDBs) > int(first)
		if hasNextPage {
			taskInDBs = taskInDBs[:first]
		}

		edges = make([]*gqlgen.TaskEdge, len(taskInDBs))
		{
			for idx, taskInDB := range taskInDBs {
				task := ConvertIntoTask(taskInDB)

				cursor, err := model.PaginationCursor{
					ID:     task.Id,
					Offset: listTasksParams.Offset + int32(idx) + 1,
					Params: model.PaginationCursorParams{
						TaskStatus: status,
					},
				}.Encode()
				if err != nil {
					return nil, oops.Wrapf(err, "failed to encode cursor")
				}

				edges[idx] = &gqlgen.TaskEdge{
					Cursor: cursor,
					Node:   task,
				}
			}
		}
	}

	var endCursor *string
	{
		if len(edges) > 0 {
			endCursor = &edges[len(edges)-1].Cursor
		} else {
			endCursor = nil
		}
	}

	totalCnt, err := qdb.CountTasks(ctx, countTasksParams)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to count tasks")
	}

	return &gqlgen.TaskConnection{
		Edges: edges,
		PageInfo: &gqlgen.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: hasNextPage,
		},
		TotalCount: gqlutil.Int64(totalCnt),
	}, nil
}

func (s *TaskService) Create(ctx context.Context, input gqlgen.CreateTaskInput) (*gqlgen.CreateTaskPayload, error) {
	if err := validation.Struct(input); err != nil {
		return &gqlgen.CreateTaskPayload{
			Error: gqlgen.BadRequestError{
				Message: err.Error(),
			},
		}, nil
	}

	var taskInDB mysql.Task
	{
		if err := sqlutil.Transact(ctx, s.mysqlContainer.App, func(txnCtx context.Context, txn *sql.Tx) (txnErr error) {
			qtxn := mysql.New(txn)

			now := s.clock.Now()

			var id int64
			if id, txnErr = qtxn.CreateTask(txnCtx, mysql.CreateTaskParams{
				Title:     input.Title,
				UpdatedAt: now,
				CreatedAt: now,
			}); txnErr != nil {
				return oops.Wrapf(txnErr, "failed to create task")
			}

			if taskInDB, txnErr = qtxn.GetTask(txnCtx, uint64(id)); txnErr != nil {
				return oops.Wrapf(txnErr, "failed to get task")
			}

			return
		}); err != nil {
			return nil, err
		}
	}

	return &gqlgen.CreateTaskPayload{
		ClientMutationId: input.ClientMutationId,
		Task:             ConvertIntoTask(taskInDB),
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

	idInDB, err := nodeid.DecodeByType(in.ID, nodeid.TypeTask)
	if err != nil {
		return oops.Errorf("invalid id")
	}

	in.idInDB = idInDB

	return nil
}

type TaskServiceCompleteOutput struct {
	Task *gqlgen.Task
}

func (s *TaskService) Complete(ctx context.Context, input gqlgen.CompleteTaskInput) (*gqlgen.CompleteTaskPayload, error) {
	if err := validation.Struct(input); err != nil {
		return &gqlgen.CompleteTaskPayload{
			Error: gqlgen.BadRequestError{
				Message: err.Error(),
			},
		}, nil
	}

	taskIDInDB, err := nodeid.DecodeByType(input.Id, nodeid.TypeTask)
	if err != nil {
		return &gqlgen.CompleteTaskPayload{
			Error: gqlgen.BadRequestError{
				Message: "invalid id",
			},
		}, nil
	}

	var (
		errTaskNotFound         = oops.Errorf("task not found")
		errTaskAlreadyCompleted = oops.Errorf("task already completed")
	)

	qdb := mysql.New(s.mysqlContainer.App)

	taskInDB, err := qdb.GetTask(ctx, taskIDInDB)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &gqlgen.CompleteTaskPayload{
				Error: gqlgen.BadRequestError{
					Message: errTaskNotFound.Error(),
				},
			}, nil
		}

		return nil, oops.Wrapf(err, "failed to get task")
	}
	if taskInDB.Status == gqlgen.TaskStatusCompleted {
		return &gqlgen.CompleteTaskPayload{
			Error: gqlgen.BadRequestError{
				Message: errTaskAlreadyCompleted.Error(),
			},
		}, nil
	}

	if err := sqlutil.Transact(ctx, s.mysqlContainer.App, func(txnCtx context.Context, txn *sql.Tx) (txnErr error) {
		qtxn := mysql.New(txn)

		if taskInDB, txnErr = qtxn.GetTaskForUpdate(txnCtx, taskInDB.ID); txnErr != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return errTaskNotFound
			}

			return oops.Wrapf(txnErr, "failed to get task for update")
		}
		if taskInDB.Status == gqlgen.TaskStatusCompleted {
			return errTaskAlreadyCompleted
		}

		now := s.clock.Now()

		if txnErr = qtxn.CompleteTask(txnCtx, mysql.CompleteTaskParams{
			ID:        taskInDB.ID,
			UpdatedAt: now,
		}); txnErr != nil {
			return oops.Wrapf(txnErr, "failed to complete task")
		}

		if taskInDB, txnErr = qtxn.GetTask(txnCtx, taskInDB.ID); txnErr != nil {
			return oops.Wrapf(txnErr, "failed to get task")
		}

		return
	}); err != nil {
		if errors.Is(err, errTaskNotFound) || errors.Is(err, errTaskAlreadyCompleted) {
			return &gqlgen.CompleteTaskPayload{
				Error: gqlgen.BadRequestError{
					Message: err.Error(),
				},
			}, nil
		}

		return nil, err
	}

	return &gqlgen.CompleteTaskPayload{
		ClientMutationId: input.ClientMutationId,
		Task:             ConvertIntoTask(taskInDB),
	}, nil
}
