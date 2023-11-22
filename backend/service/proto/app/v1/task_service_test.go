package appv1_test

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	appv1 "app/gen/buf/app/v1"
	"app/internal/testutil"
	here "app/service/proto/app/v1"
)

func TestTaskService(t *testing.T) {
	setup(t)
	t.Cleanup(func() {
		teardown(t)
	})

	ctx := context.Background()

	s := here.NewTaskService(vldtr, clock, mysqlCtr)

	var (
		task1 *appv1.Task
		task2 *appv1.Task
	)

	t.Run("success: no tasks", func(t *testing.T) {
		{
			resp, err := s.List(ctx, connect.NewRequest(&appv1.TaskServiceListRequest{}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, 0, len(resp.Msg.Tasks))
		}
	})

	t.Run("success: create task1", func(t *testing.T) {
		{
			title := "task1"

			resp, err := s.Create(ctx, connect.NewRequest(&appv1.TaskServiceCreateRequest{
				Title: title,
			}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, title, resp.Msg.Task.Title)
			testutil.Equal(t, appv1.TaskStatus_TASK_STATUS_UNCOMPLETED, resp.Msg.Task.Status)

			task1 = resp.Msg.Task
		}
		{
			resp, err := s.Get(ctx, connect.NewRequest(&appv1.TaskServiceGetRequest{
				Id: task1.Id,
			}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, task1, resp.Msg.Task, cmpopts.IgnoreUnexported(appv1.Task{}))
		}
		{
			resp, err := s.List(ctx, connect.NewRequest(&appv1.TaskServiceListRequest{}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, 1, len(resp.Msg.Tasks))
			testutil.Equal(t, task1, resp.Msg.Tasks[0], cmpopts.IgnoreUnexported(appv1.Task{}))
		}
	})

	t.Run("success: create task2", func(t *testing.T) {
		{
			title := "task2"

			resp, err := s.Create(ctx, connect.NewRequest(&appv1.TaskServiceCreateRequest{
				Title: title,
			}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, title, resp.Msg.Task.Title)
			testutil.Equal(t, appv1.TaskStatus_TASK_STATUS_UNCOMPLETED, resp.Msg.Task.Status)

			task2 = resp.Msg.Task
		}
		{
			resp, err := s.Get(ctx, connect.NewRequest(&appv1.TaskServiceGetRequest{
				Id: task2.Id,
			}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, task2, resp.Msg.Task, cmpopts.IgnoreUnexported(appv1.Task{}))
		}
		{
			resp, err := s.List(ctx, connect.NewRequest(&appv1.TaskServiceListRequest{}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, 2, len(resp.Msg.Tasks))
			testutil.Equal(t, task1, resp.Msg.Tasks[0], cmpopts.IgnoreUnexported(appv1.Task{}))
			testutil.Equal(t, task2, resp.Msg.Tasks[1], cmpopts.IgnoreUnexported(appv1.Task{}))
		}
	})

	t.Run("success: update task1 title", func(t *testing.T) {
		{
			title := "task1_updated"

			fm, err := fieldmaskpb.New(&appv1.Task{}, "id", "title")
			if err != nil {
				t.Fatal(err)
			}

			resp, err := s.Update(ctx, connect.NewRequest(&appv1.TaskServiceUpdateRequest{
				Task: &appv1.TaskServiceUpdateRequest_Fields{
					Id:    task1.Id,
					Title: &title,
				},
				FieldMask: fm,
			}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, task1.Id, resp.Msg.Task.Id)
			testutil.Equal(t, title, resp.Msg.Task.Title)
			testutil.Equal(t, task1.Status, resp.Msg.Task.Status)

			task1 = resp.Msg.Task
		}
		{
			resp, err := s.Get(ctx, connect.NewRequest(&appv1.TaskServiceGetRequest{
				Id: task1.Id,
			}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, task1, resp.Msg.Task, cmpopts.IgnoreUnexported(appv1.Task{}))
		}
	})

	t.Run("success: update task1 status", func(t *testing.T) {
		{
			status := appv1.TaskStatus_TASK_STATUS_COMPLETED

			fm, err := fieldmaskpb.New(&appv1.Task{}, "id", "status")
			if err != nil {
				t.Fatal(err)
			}

			resp, err := s.Update(ctx, connect.NewRequest(&appv1.TaskServiceUpdateRequest{
				Task: &appv1.TaskServiceUpdateRequest_Fields{
					Id:     task1.Id,
					Status: &status,
				},
				FieldMask: fm,
			}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, task1.Id, resp.Msg.Task.Id)
			testutil.Equal(t, task1.Title, resp.Msg.Task.Title)
			testutil.Equal(t, status, resp.Msg.Task.Status)

			task1 = resp.Msg.Task
		}
		{
			resp, err := s.Get(ctx, connect.NewRequest(&appv1.TaskServiceGetRequest{
				Id: task1.Id,
			}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, task1, resp.Msg.Task, cmpopts.IgnoreUnexported(appv1.Task{}))
		}
	})

	t.Run("success: delete task1", func(t *testing.T) {
		{
			if _, err := s.Delete(ctx, connect.NewRequest(&appv1.TaskServiceDeleteRequest{
				Id: task1.Id,
			})); err != nil {
				t.Fatal(err)
			}

			task1 = nil
		}
		{
			resp, err := s.List(ctx, connect.NewRequest(&appv1.TaskServiceListRequest{}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, 1, len(resp.Msg.Tasks))
			testutil.Equal(t, task2, resp.Msg.Tasks[0], cmpopts.IgnoreUnexported(appv1.Task{}))
		}
	})
}
