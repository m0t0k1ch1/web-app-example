package appv1_test

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/m0t0k1ch1-go/timeutil"

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

	s := here.NewTaskService(clock, mysqlCtr)

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
		now := timeutil.Now()
		clock.Set(now)

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
			testutil.Equal(t, now.Time().Unix(), resp.Msg.Task.UpdatedAt)
			testutil.Equal(t, now.Time().Unix(), resp.Msg.Task.CreatedAt)

			task1 = resp.Msg.Task
		}
		{
			resp, err := s.Get(ctx, connect.NewRequest(&appv1.TaskServiceGetRequest{
				Id: task1.Id,
			}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, task1.Id, resp.Msg.Task.Id)
			testutil.Equal(t, task1.Title, resp.Msg.Task.Title)
			testutil.Equal(t, task1.Status, resp.Msg.Task.Status)
			testutil.Equal(t, task1.UpdatedAt, resp.Msg.Task.UpdatedAt)
			testutil.Equal(t, task1.CreatedAt, resp.Msg.Task.CreatedAt)
		}
		{
			resp, err := s.List(ctx, connect.NewRequest(&appv1.TaskServiceListRequest{}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, 1, len(resp.Msg.Tasks))
			testutil.Equal(t, task1.Id, resp.Msg.Tasks[0].Id)
			testutil.Equal(t, task1.Title, resp.Msg.Tasks[0].Title)
			testutil.Equal(t, task1.Status, resp.Msg.Tasks[0].Status)
			testutil.Equal(t, task1.UpdatedAt, resp.Msg.Tasks[0].UpdatedAt)
			testutil.Equal(t, task1.CreatedAt, resp.Msg.Tasks[0].CreatedAt)
		}
	})

	t.Run("success: create task2", func(t *testing.T) {
		now := timeutil.Now()
		clock.Set(now)

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
			testutil.Equal(t, now.Time().Unix(), resp.Msg.Task.UpdatedAt)
			testutil.Equal(t, now.Time().Unix(), resp.Msg.Task.CreatedAt)

			task2 = resp.Msg.Task
		}
		{
			resp, err := s.Get(ctx, connect.NewRequest(&appv1.TaskServiceGetRequest{
				Id: task2.Id,
			}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, task2.Id, resp.Msg.Task.Id)
			testutil.Equal(t, task2.Title, resp.Msg.Task.Title)
			testutil.Equal(t, task2.Status, resp.Msg.Task.Status)
			testutil.Equal(t, task2.UpdatedAt, resp.Msg.Task.UpdatedAt)
			testutil.Equal(t, task2.CreatedAt, resp.Msg.Task.CreatedAt)
		}
		{
			resp, err := s.List(ctx, connect.NewRequest(&appv1.TaskServiceListRequest{}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, 2, len(resp.Msg.Tasks))
			testutil.Equal(t, task1.Id, resp.Msg.Tasks[0].Id)
			testutil.Equal(t, task1.Title, resp.Msg.Tasks[0].Title)
			testutil.Equal(t, task1.Status, resp.Msg.Tasks[0].Status)
			testutil.Equal(t, task1.UpdatedAt, resp.Msg.Tasks[0].UpdatedAt)
			testutil.Equal(t, task1.CreatedAt, resp.Msg.Tasks[0].CreatedAt)
			testutil.Equal(t, task2.Id, resp.Msg.Tasks[1].Id)
			testutil.Equal(t, task2.Title, resp.Msg.Tasks[1].Title)
			testutil.Equal(t, task2.Status, resp.Msg.Tasks[1].Status)
			testutil.Equal(t, task2.UpdatedAt, resp.Msg.Tasks[1].UpdatedAt)
			testutil.Equal(t, task2.CreatedAt, resp.Msg.Tasks[1].CreatedAt)
		}
	})

	t.Run("success: update task1", func(t *testing.T) {
		now := timeutil.Now()
		clock.Set(now)

		{
			title := "task1_updated"
			status := appv1.TaskStatus_TASK_STATUS_COMPLETED

			resp, err := s.Update(ctx, connect.NewRequest(&appv1.TaskServiceUpdateRequest{
				Id:     task1.Id,
				Title:  title,
				Status: status,
			}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, task1.Id, resp.Msg.Task.Id)
			testutil.Equal(t, title, resp.Msg.Task.Title)
			testutil.Equal(t, status, resp.Msg.Task.Status)
			testutil.Equal(t, now.Time().Unix(), resp.Msg.Task.UpdatedAt)
			testutil.Equal(t, task1.CreatedAt, resp.Msg.Task.CreatedAt)

			task1 = resp.Msg.Task
		}
		{
			resp, err := s.Get(ctx, connect.NewRequest(&appv1.TaskServiceGetRequest{
				Id: task1.Id,
			}))
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, task1.Id, resp.Msg.Task.Id)
			testutil.Equal(t, task1.Title, resp.Msg.Task.Title)
			testutil.Equal(t, task1.Status, resp.Msg.Task.Status)
			testutil.Equal(t, task1.UpdatedAt, resp.Msg.Task.UpdatedAt)
			testutil.Equal(t, task1.CreatedAt, resp.Msg.Task.CreatedAt)
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
			testutil.Equal(t, task2.Id, resp.Msg.Tasks[0].Id)
			testutil.Equal(t, task2.Title, resp.Msg.Tasks[0].Title)
			testutil.Equal(t, task2.Status, resp.Msg.Tasks[0].Status)
			testutil.Equal(t, task2.UpdatedAt, resp.Msg.Tasks[0].UpdatedAt)
			testutil.Equal(t, task2.CreatedAt, resp.Msg.Tasks[0].CreatedAt)
		}
	})
}
