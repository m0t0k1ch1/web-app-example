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
	"app/domain/model"
	"app/domain/validation"
	"app/gen/gqlgen"
	"app/gen/sqlc/mysql"
	"app/library/gqlerrutil"
	"app/library/idutil"
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

	idInDBInAfter *uint64
}

func (in *TaskServiceListInput) Validate() error {
	if err := validation.Struct(in); err != nil {
		return err
	}

	var idInDBInAfter *uint64
	{
		if in.After != nil {
			cursor, err := model.DecodePaginationCursor(*in.After)
			if err != nil {
				return oops.Errorf("invalid after")
			}
			if !cmp.Equal(cursor.TaskStatus, in.Status) {
				return oops.Errorf("invalid after")
			}

			idInDB, err := idutil.DecodeTaskID(cursor.ID)
			if err != nil {
				return oops.Errorf("invalid after")
			}

			idInDBInAfter = &idInDB
		}
	}

	in.idInDBInAfter = idInDBInAfter

	return nil
}

type TaskServiceListOutput struct {
	TaskConnection *gqlgen.TaskConnection
}

func (s *TaskService) List(ctx context.Context, in TaskServiceListInput) (TaskServiceListOutput, error) {
	if err := in.Validate(); err != nil {
		return TaskServiceListOutput{}, gqlerrutil.NewBadUserInputError(ctx, err)
	}

	qdb := mysql.New(s.mysqlContainer.App)

	var (
		edges       []*gqlgen.TaskEdge
		nodes       []*gqlgen.Task
		hasNextPage bool
	)
	{
		var (
			taskInDBs []mysql.Task
			err       error
		)
		{
			limit := in.First + 1

			if in.Status != nil {
				if in.After != nil {
					if taskInDBs, err = qdb.ListFirstTasksAfterCursorByStatus(ctx, mysql.ListFirstTasksAfterCursorByStatusParams{
						Status: *in.Status,
						After:  *in.idInDBInAfter,
						Limit:  limit,
					}); err != nil {
						return TaskServiceListOutput{}, gqlerrutil.NewInternalServerError(ctx, oops.Wrapf(err, "failed to list first tasks after cursor by status"))
					}
				} else {
					if taskInDBs, err = qdb.ListFirstTasksByStatus(ctx, mysql.ListFirstTasksByStatusParams{
						Status: *in.Status,
						Limit:  limit,
					}); err != nil {
						return TaskServiceListOutput{}, gqlerrutil.NewInternalServerError(ctx, oops.Wrapf(err, "failed to list first tasks by status"))
					}
				}
			} else {
				if in.After != nil {
					if taskInDBs, err = qdb.ListFirstTasksAfterCursor(ctx, mysql.ListFirstTasksAfterCursorParams{
						After: *in.idInDBInAfter,
						Limit: limit,
					}); err != nil {
						return TaskServiceListOutput{}, gqlerrutil.NewInternalServerError(ctx, oops.Wrapf(err, "failed to list first tasks after cursor"))
					}
				} else {
					if taskInDBs, err = qdb.ListFirstTasks(ctx, limit); err != nil {
						return TaskServiceListOutput{}, gqlerrutil.NewInternalServerError(ctx, oops.Wrapf(err, "failed to list first tasks"))
					}
				}
			}
		}

		hasNextPage = len(taskInDBs) > int(in.First)
		if hasNextPage {
			taskInDBs = taskInDBs[:in.First]
		}

		edges, nodes, err = convertIntoTaskEdgesAndTasks(taskInDBs, in.Status)
		if err != nil {
			return TaskServiceListOutput{}, gqlerrutil.NewInternalServerError(ctx, oops.Wrapf(err, "failed to convert into task edges and tasks"))
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

	var totalCnt int64
	{
		var err error

		if in.Status != nil {
			if totalCnt, err = qdb.CountTasksByStatus(ctx, *in.Status); err != nil {
				return TaskServiceListOutput{}, gqlerrutil.NewInternalServerError(ctx, oops.Wrapf(err, "failed to count tasks by status"))
			}
		} else {
			if totalCnt, err = qdb.CountAllTasks(ctx); err != nil {
				return TaskServiceListOutput{}, gqlerrutil.NewInternalServerError(ctx, oops.Wrapf(err, "failed to count tasks"))
			}
		}
	}

	return TaskServiceListOutput{
		TaskConnection: &gqlgen.TaskConnection{
			Edges: edges,
			Nodes: nodes,
			PageInfo: &gqlgen.PageInfo{
				EndCursor:   endCursor,
				HasNextPage: hasNextPage,
			},
			TotalCount: gqlutil.Int64(totalCnt),
		},
	}, nil
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
		return TaskServiceCreateOutput{}, gqlerrutil.NewBadUserInputError(ctx, err)
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
				return gqlerrutil.NewInternalServerError(txnCtx, oops.Wrapf(txnErr, "failed to create task"))
			}

			if taskInDB, txnErr = qtxn.GetTask(txnCtx, uint64(id)); txnErr != nil {
				return gqlerrutil.NewInternalServerError(txnCtx, oops.Wrapf(txnErr, "failed to get task"))
			}

			return
		}); err != nil {
			return TaskServiceCreateOutput{}, err
		}
	}

	return TaskServiceCreateOutput{
		Task: convertIntoTask(taskInDB),
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

	idInDB, err := idutil.DecodeTaskID(in.ID)
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
		return TaskServiceCompleteOutput{}, gqlerrutil.NewBadUserInputError(ctx, err)
	}

	qdb := mysql.New(s.mysqlContainer.App)

	taskInDB, err := qdb.GetTask(ctx, in.idInDB)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TaskServiceCompleteOutput{}, nil
		}

		return TaskServiceCompleteOutput{}, gqlerrutil.NewInternalServerError(ctx, oops.Wrapf(err, "failed to get task"))
	}
	if taskInDB.Status == gqlgen.TaskStatusCompleted {
		return TaskServiceCompleteOutput{}, gqlerrutil.NewBadUserInputError(ctx, oops.Errorf("task already completed"))
	}

	if err := sqlutil.Transact(ctx, s.mysqlContainer.App, func(txnCtx context.Context, txn *sql.Tx) (txnErr error) {
		qtxn := mysql.New(txn)

		if taskInDB, txnErr = qtxn.GetTaskForUpdate(txnCtx, taskInDB.ID); txnErr != nil {
			if errors.Is(err, sql.ErrNoRows) {
				taskInDB = mysql.Task{}

				return nil
			}

			return gqlerrutil.NewInternalServerError(txnCtx, oops.Wrapf(txnErr, "failed to get task for update"))
		}
		if taskInDB.Status == gqlgen.TaskStatusCompleted {
			return gqlerrutil.NewBadUserInputError(ctx, oops.Errorf("task already completed"))
		}

		now := s.clock.Now()

		if txnErr = qtxn.CompleteTask(txnCtx, mysql.CompleteTaskParams{
			ID:        taskInDB.ID,
			UpdatedAt: now,
		}); txnErr != nil {
			return gqlerrutil.NewInternalServerError(txnCtx, oops.Wrapf(txnErr, "failed to complete task"))
		}

		if taskInDB, txnErr = qtxn.GetTask(txnCtx, taskInDB.ID); txnErr != nil {
			return gqlerrutil.NewInternalServerError(txnCtx, oops.Wrapf(txnErr, "failed to get task"))
		}

		return
	}); err != nil {
		return TaskServiceCompleteOutput{}, err
	}

	if taskInDB.ID == 0 {
		return TaskServiceCompleteOutput{}, nil
	}

	return TaskServiceCompleteOutput{
		Task: convertIntoTask(taskInDB),
	}, nil
}
