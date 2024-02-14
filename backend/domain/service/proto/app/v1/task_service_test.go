package appv1_test

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/m0t0k1ch1-go/coreutil"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	here "app/domain/service/proto/app/v1"
	appv1 "app/gen/buf/app/v1"
	"app/internal/testutil"
)

func TestTaskService(t *testing.T) {
	setup(t)
	t.Cleanup(func() {
		teardown(t)
	})

	ctx := context.Background()

	taskService := here.NewTaskService(
		clock,
		mysqlCtr,
	)

	var (
		task1 *appv1.Task
		task2 *appv1.Task
	)

	t.Run("success: no tasks", func(t *testing.T) {
		{
			resp, err := taskService.List(ctx, connect.NewRequest(&appv1.TaskServiceListRequest{}))
			require.Nil(t, err)

			testutil.Equal(t, 0, len(resp.Msg.Tasks))
		}
	})

	t.Run("success: create task1", func(t *testing.T) {
		{
			title := "task1"

			resp, err := taskService.Create(ctx, connect.NewRequest(&appv1.TaskServiceCreateRequest{
				Title: title,
			}))
			require.Nil(t, err)

			testutil.Equal(t, title, resp.Msg.Task.Title)
			testutil.Equal(t, appv1.TaskStatus_TASK_STATUS_UNCOMPLETED, resp.Msg.Task.Status)

			task1 = resp.Msg.Task
		}
		{
			resp, err := taskService.Get(ctx, connect.NewRequest(&appv1.TaskServiceGetRequest{
				Id: task1.Id,
			}))
			require.Nil(t, err)

			testutil.Equal(t, task1, resp.Msg.Task, cmpopts.IgnoreUnexported(appv1.Task{}))
		}
		{
			resp, err := taskService.List(ctx, connect.NewRequest(&appv1.TaskServiceListRequest{}))
			require.Nil(t, err)

			testutil.Equal(t, 1, len(resp.Msg.Tasks))
			testutil.Equal(t, task1, resp.Msg.Tasks[0], cmpopts.IgnoreUnexported(appv1.Task{}))
		}
	})

	t.Run("success: create task2", func(t *testing.T) {
		{
			title := "task2"

			resp, err := taskService.Create(ctx, connect.NewRequest(&appv1.TaskServiceCreateRequest{
				Title: title,
			}))
			require.Nil(t, err)

			testutil.Equal(t, title, resp.Msg.Task.Title)
			testutil.Equal(t, appv1.TaskStatus_TASK_STATUS_UNCOMPLETED, resp.Msg.Task.Status)

			task2 = resp.Msg.Task
		}
		{
			resp, err := taskService.Get(ctx, connect.NewRequest(&appv1.TaskServiceGetRequest{
				Id: task2.Id,
			}))
			require.Nil(t, err)

			testutil.Equal(t, task2, resp.Msg.Task, cmpopts.IgnoreUnexported(appv1.Task{}))
		}
		{
			resp, err := taskService.List(ctx, connect.NewRequest(&appv1.TaskServiceListRequest{}))
			require.Nil(t, err)

			testutil.Equal(t, 2, len(resp.Msg.Tasks))
			testutil.Equal(t, task1, resp.Msg.Tasks[0], cmpopts.IgnoreUnexported(appv1.Task{}))
			testutil.Equal(t, task2, resp.Msg.Tasks[1], cmpopts.IgnoreUnexported(appv1.Task{}))
		}
	})

	t.Run("failure: update task1 title: title required", func(t *testing.T) {
		{
			fm, err := fieldmaskpb.New(&appv1.Task{}, "id", "title")
			require.Nil(t, err)

			_, err = taskService.Update(ctx, connect.NewRequest(&appv1.TaskServiceUpdateRequest{
				Task: &appv1.TaskServiceUpdateRequest_Fields{
					Id: task1.Id,
				},
				FieldMask: fm,
			}))

			var connectErr *connect.Error
			require.ErrorAs(t, err, &connectErr)

			testutil.Equal(t, connect.CodeInvalidArgument, connectErr.Code())
			testutil.Equal(t, "title required", connectErr.Message())
		}
	})

	t.Run("success: update task1 title", func(t *testing.T) {
		{
			task1.Title = "task1_updated"

			fm, err := fieldmaskpb.New(&appv1.Task{}, "id", "title")
			require.Nil(t, err)

			resp, err := taskService.Update(ctx, connect.NewRequest(&appv1.TaskServiceUpdateRequest{
				Task: &appv1.TaskServiceUpdateRequest_Fields{
					Id:    task1.Id,
					Title: &task1.Title,
				},
				FieldMask: fm,
			}))
			require.Nil(t, err)

			testutil.Equal(t, task1, resp.Msg.Task, cmpopts.IgnoreUnexported(appv1.Task{}))
		}
		{
			resp, err := taskService.Get(ctx, connect.NewRequest(&appv1.TaskServiceGetRequest{
				Id: task1.Id,
			}))
			require.Nil(t, err)

			testutil.Equal(t, task1, resp.Msg.Task, cmpopts.IgnoreUnexported(appv1.Task{}))
		}
	})

	t.Run("failure: update task1 status: status required", func(t *testing.T) {
		{
			fm, err := fieldmaskpb.New(&appv1.Task{}, "id", "status")
			require.Nil(t, err)

			_, err = taskService.Update(ctx, connect.NewRequest(&appv1.TaskServiceUpdateRequest{
				Task: &appv1.TaskServiceUpdateRequest_Fields{
					Id: task1.Id,
				},
				FieldMask: fm,
			}))

			var connectErr *connect.Error
			require.ErrorAs(t, err, &connectErr)

			testutil.Equal(t, connect.CodeInvalidArgument, connectErr.Code())
			testutil.Equal(t, "status required", connectErr.Message())
		}
	})

	t.Run("success: update task1 status", func(t *testing.T) {
		{
			task1.Status = appv1.TaskStatus_TASK_STATUS_COMPLETED

			fm, err := fieldmaskpb.New(&appv1.Task{}, "id", "status")
			require.Nil(t, err)

			resp, err := taskService.Update(ctx, connect.NewRequest(&appv1.TaskServiceUpdateRequest{
				Task: &appv1.TaskServiceUpdateRequest_Fields{
					Id:     task1.Id,
					Status: &task1.Status,
				},
				FieldMask: fm,
			}))
			require.Nil(t, err)

			testutil.Equal(t, task1, resp.Msg.Task, cmpopts.IgnoreUnexported(appv1.Task{}))
		}
		{
			resp, err := taskService.Get(ctx, connect.NewRequest(&appv1.TaskServiceGetRequest{
				Id: task1.Id,
			}))
			require.Nil(t, err)

			testutil.Equal(t, task1, resp.Msg.Task, cmpopts.IgnoreUnexported(appv1.Task{}))
		}
		{
			resp, err := taskService.List(ctx, connect.NewRequest(&appv1.TaskServiceListRequest{
				Status: coreutil.Ptr(appv1.TaskStatus_TASK_STATUS_UNCOMPLETED),
			}))
			require.Nil(t, err)

			testutil.Equal(t, 1, len(resp.Msg.Tasks))
			testutil.Equal(t, task2, resp.Msg.Tasks[0], cmpopts.IgnoreUnexported(appv1.Task{}))
		}
	})

	t.Run("success: delete task1", func(t *testing.T) {
		{
			_, err := taskService.Delete(ctx, connect.NewRequest(&appv1.TaskServiceDeleteRequest{
				Id: task1.Id,
			}))
			require.Nil(t, err)

			task1 = nil
		}
		{
			resp, err := taskService.List(ctx, connect.NewRequest(&appv1.TaskServiceListRequest{}))
			require.Nil(t, err)

			testutil.Equal(t, 1, len(resp.Msg.Tasks))
			testutil.Equal(t, task2, resp.Msg.Tasks[0], cmpopts.IgnoreUnexported(appv1.Task{}))
		}
	})
}
