package appv1

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/m0t0k1ch1-go/timeutil"
	"github.com/pkg/errors"

	appv1 "app/gen/buf/app/v1"
	"app/gen/sqlc/mysql"
	"app/library/sqlutil"
	"app/service/proto"
)

type TaskService struct {
	clock timeutil.Clock
	mysql *sql.DB
}

func NewTaskService(clock timeutil.Clock, mysql *sql.DB) *TaskService {
	return &TaskService{
		clock: clock,
		mysql: mysql,
	}
}

func (s *TaskService) Create(ctx context.Context, req *connect.Request[appv1.TaskServiceCreateRequest]) (*connect.Response[appv1.TaskServiceCreateResponse], error) {
	var task mysql.Task

	if err := sqlutil.Transact(ctx, s.mysql, func(txCtx context.Context, tx *sql.Tx) (txErr error) {
		qtx := mysql.New(tx)

		now := s.clock.Now()

		var id int64
		if id, txErr = qtx.CreateTask(txCtx, mysql.CreateTaskParams{
			Title:     req.Msg.Title,
			UpdatedAt: now,
			CreatedAt: now,
		}); txErr != nil {
			return proto.NewUnknownError(errors.Wrap(txErr, "failed to create task"))
		}

		if task, txErr = qtx.GetTask(txCtx, uint64(id)); txErr != nil {
			return proto.NewUnknownError(errors.Wrap(txErr, "failed to get task"))
		}

		return
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceCreateResponse{
		Task: ConvertTask(task),
	}), nil
}

func (s *TaskService) Get(ctx context.Context, req *connect.Request[appv1.TaskServiceGetRequest]) (*connect.Response[appv1.TaskServiceGetResponse], error) {
	task, err := GetTaskOrError(ctx, s.mysql, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceGetResponse{
		Task: ConvertTask(task),
	}), nil
}

func (s *TaskService) List(ctx context.Context, req *connect.Request[appv1.TaskServiceListRequest]) (*connect.Response[appv1.TaskServiceListResponse], error) {
	qdb := mysql.New(s.mysql)

	tasks, err := qdb.ListTasks(ctx)
	if err != nil {
		return nil, proto.NewUnknownError(errors.Wrap(err, "failed to list tasks"))
	}

	return connect.NewResponse(&appv1.TaskServiceListResponse{
		Tasks: ConvertTasks(tasks),
	}), nil
}

func (s *TaskService) Update(ctx context.Context, req *connect.Request[appv1.TaskServiceUpdateRequest]) (*connect.Response[appv1.TaskServiceUpdateResponse], error) {
	task, err := GetTaskOrError(ctx, s.mysql, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	if err := sqlutil.Transact(ctx, s.mysql, func(txCtx context.Context, tx *sql.Tx) (txErr error) {
		qtx := mysql.New(tx)

		if task, txErr = qtx.GetTaskForUpdate(txCtx, task.ID); txErr != nil {
			if errors.Is(txErr, sql.ErrNoRows) {
				return proto.NewNotFoundError(errors.Wrap(txErr, "task not found"))
			}

			return proto.NewUnknownError(errors.Wrap(txErr, "failed to get task for update"))
		}

		if txErr = qtx.UpdateTask(txCtx, mysql.UpdateTaskParams{
			ID:        task.ID,
			Title:     req.Msg.Title,
			Status:    req.Msg.Status,
			UpdatedAt: s.clock.Now(),
		}); txErr != nil {
			return proto.NewUnknownError(errors.Wrap(txErr, "failed to update task"))
		}

		if task, txErr = qtx.GetTask(txCtx, task.ID); txErr != nil {
			return proto.NewUnknownError(errors.Wrap(txErr, "failed to get task"))
		}

		return
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceUpdateResponse{
		Task: ConvertTask(task),
	}), nil
}

func (s *TaskService) Delete(ctx context.Context, req *connect.Request[appv1.TaskServiceDeleteRequest]) (*connect.Response[appv1.TaskServiceDeleteResponse], error) {
	task, err := GetTaskOrError(ctx, s.mysql, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	if err := sqlutil.Transact(ctx, s.mysql, func(txCtx context.Context, tx *sql.Tx) (txErr error) {
		qtx := mysql.New(tx)

		if task, txErr = qtx.GetTaskForUpdate(txCtx, task.ID); txErr != nil {
			if errors.Is(txErr, sql.ErrNoRows) {
				return proto.NewNotFoundError(errors.Wrap(txErr, "task not found"))
			}

			return proto.NewUnknownError(errors.Wrap(txErr, "failed to get task for update"))
		}

		if txErr = qtx.DeleteTask(txCtx, task.ID); txErr != nil {
			return proto.NewUnknownError(errors.Wrap(txErr, "failed to delete task"))
		}

		return
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&appv1.TaskServiceDeleteResponse{}), nil
}
