package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pkg/errors"

	"app/config"
	appv1 "app/gen/buf/app/v1"
	"app/internal/testutil"
	"app/library/timeutil"
)

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func TestMain(m *testing.M) {
	os.Exit(testutil.Run(m))
}

func TestApp(t *testing.T) {
	ctx := context.Background()

	conf := config.App{}
	{
		mysqlConf, teardown, err := testutil.SetupMySQL(ctx, "test")
		if err != nil {
			t.Fatal(errors.Wrap(err, "failed to setup mysql"))
		}
		defer teardown()

		conf.MySQL = mysqlConf
	}

	app, err := NewApp(ctx, conf)
	if err != nil {
		t.Fatal(errors.Wrap(err, "failed to initialize app"))
	}

	srv := httptest.NewServer(app)
	defer srv.Close()

	c, err := testutil.NewAPIClient(srv.URL)
	if err != nil {
		t.Fatal(errors.Wrap(err, "failed to initialize api client"))
	}

	var (
		task1 *Task
		task2 *Task
	)

	t.Run("success: no tasks", func(t *testing.T) {
		{
			var resp struct {
				Tasks []*Task `json:"tasks"`
			}
			statusCode, err := c.DoAPI(ctx,
				http.MethodPost,
				"/grpc/app.v1.TaskService/List",
				struct{}{},
				&resp,
			)
			if err != nil {
				t.Fatal(errors.Wrap(err, "failed to list tasks"))
			}

			testutil.Equal(t, http.StatusOK, statusCode)
			testutil.Equal(t, 0, len(resp.Tasks))
		}
	})

	t.Run("success: create task1", func(t *testing.T) {
		now := timeutil.Now()
		timeutil.Lock(now)
		defer timeutil.Unlock()

		{
			title := "task1"

			var resp struct {
				Task *Task `json:"task"`
			}
			statusCode, err := c.DoAPI(ctx,
				http.MethodPost,
				"/grpc/app.v1.TaskService/Create",
				struct {
					Title string `json:"title"`
				}{
					Title: title,
				},
				&resp,
			)
			if err != nil {
				t.Fatal(errors.Wrap(err, "failed to create task"))
			}

			testutil.Equal(t, http.StatusOK, statusCode)
			testutil.Equal(t, title, resp.Task.Title)
			testutil.Equal(t, appv1.TaskStatus_TASK_STATUS_UNCOMPLETED.String(), resp.Task.Status)
			testutil.Equal(t, now.String(), resp.Task.UpdatedAt)
			testutil.Equal(t, now.String(), resp.Task.CreatedAt)

			task1 = resp.Task
		}
		{
			var resp struct {
				Task *Task `json:"task"`
			}
			statusCode, err := c.DoAPI(ctx,
				http.MethodPost,
				"/grpc/app.v1.TaskService/Get",
				struct {
					ID string `json:"id"`
				}{
					ID: task1.ID,
				},
				&resp,
			)
			if err != nil {
				t.Fatal(errors.Wrap(err, "failed to get task"))
			}

			testutil.Equal(t, http.StatusOK, statusCode)
			testutil.Equal(t, task1.ID, resp.Task.ID)
			testutil.Equal(t, task1.Title, resp.Task.Title)
			testutil.Equal(t, task1.Status, resp.Task.Status)
			testutil.Equal(t, task1.UpdatedAt, resp.Task.UpdatedAt)
			testutil.Equal(t, task1.CreatedAt, resp.Task.CreatedAt)
		}
		{
			var resp struct {
				Tasks []*Task `json:"tasks"`
			}
			statusCode, err := c.DoAPI(ctx,
				http.MethodPost,
				"/grpc/app.v1.TaskService/List",
				struct{}{},
				&resp,
			)
			if err != nil {
				t.Fatal(errors.Wrap(err, "failed to list tasks"))
			}

			testutil.Equal(t, http.StatusOK, statusCode)
			testutil.Equal(t, 1, len(resp.Tasks))
			testutil.Equal(t, task1.ID, resp.Tasks[0].ID)
			testutil.Equal(t, task1.Title, resp.Tasks[0].Title)
			testutil.Equal(t, task1.Status, resp.Tasks[0].Status)
			testutil.Equal(t, task1.UpdatedAt, resp.Tasks[0].UpdatedAt)
			testutil.Equal(t, task1.CreatedAt, resp.Tasks[0].CreatedAt)
		}
	})

	t.Run("success: create task2", func(t *testing.T) {
		now := timeutil.Now()
		timeutil.Lock(now)
		defer timeutil.Unlock()

		{
			title := "task2"

			var resp struct {
				Task *Task `json:"task"`
			}
			statusCode, err := c.DoAPI(ctx,
				http.MethodPost,
				"/grpc/app.v1.TaskService/Create",
				struct {
					Title string `json:"title"`
				}{
					Title: title,
				},
				&resp,
			)
			if err != nil {
				t.Fatal(errors.Wrap(err, "failed to create task"))
			}

			testutil.Equal(t, http.StatusOK, statusCode)
			testutil.Equal(t, title, resp.Task.Title)
			testutil.Equal(t, appv1.TaskStatus_TASK_STATUS_UNCOMPLETED.String(), resp.Task.Status)
			testutil.Equal(t, now.String(), resp.Task.UpdatedAt)
			testutil.Equal(t, now.String(), resp.Task.CreatedAt)

			task2 = resp.Task
		}
		{
			var resp struct {
				Task *Task `json:"task"`
			}
			statusCode, err := c.DoAPI(ctx,
				http.MethodPost,
				"/grpc/app.v1.TaskService/Get",
				struct {
					ID string `json:"id"`
				}{
					ID: task2.ID,
				},
				&resp,
			)
			if err != nil {
				t.Fatal(errors.Wrap(err, "failed to get task"))
			}

			testutil.Equal(t, http.StatusOK, statusCode)
			testutil.Equal(t, task2.ID, resp.Task.ID)
			testutil.Equal(t, task2.Title, resp.Task.Title)
			testutil.Equal(t, task2.Status, resp.Task.Status)
			testutil.Equal(t, task2.UpdatedAt, resp.Task.UpdatedAt)
			testutil.Equal(t, task2.CreatedAt, resp.Task.CreatedAt)
		}
		{
			var resp struct {
				Tasks []*Task `json:"tasks"`
			}
			statusCode, err := c.DoAPI(ctx,
				http.MethodPost,
				"/grpc/app.v1.TaskService/List",
				struct{}{},
				&resp,
			)
			if err != nil {
				t.Fatal(errors.Wrap(err, "failed to list tasks"))
			}

			testutil.Equal(t, http.StatusOK, statusCode)
			testutil.Equal(t, 2, len(resp.Tasks))
			testutil.Equal(t, task2.ID, resp.Tasks[0].ID)
			testutil.Equal(t, task2.Title, resp.Tasks[0].Title)
			testutil.Equal(t, task2.Status, resp.Tasks[0].Status)
			testutil.Equal(t, task2.UpdatedAt, resp.Tasks[0].UpdatedAt)
			testutil.Equal(t, task2.CreatedAt, resp.Tasks[0].CreatedAt)
			testutil.Equal(t, task1.ID, resp.Tasks[1].ID)
			testutil.Equal(t, task1.Title, resp.Tasks[1].Title)
			testutil.Equal(t, task1.Status, resp.Tasks[1].Status)
			testutil.Equal(t, task1.UpdatedAt, resp.Tasks[1].UpdatedAt)
			testutil.Equal(t, task1.CreatedAt, resp.Tasks[1].CreatedAt)
		}
	})

	t.Run("success: update task1", func(t *testing.T) {
		now := timeutil.Now()
		timeutil.Lock(now)
		defer timeutil.Unlock()

		{
			title := "task1_updated"
			status := appv1.TaskStatus_TASK_STATUS_COMPLETED.String()

			var resp struct {
				Task *Task `json:"task"`
			}
			statusCode, err := c.DoAPI(ctx,
				http.MethodPost,
				"/grpc/app.v1.TaskService/Update",
				struct {
					ID     string `json:"id"`
					Title  string `json:"title"`
					Status string `json:"status"`
				}{
					ID:     task1.ID,
					Title:  title,
					Status: status,
				},
				&resp,
			)
			if err != nil {
				t.Fatal(errors.Wrap(err, "failed to update task"))
			}

			testutil.Equal(t, http.StatusOK, statusCode)
			testutil.Equal(t, task1.ID, resp.Task.ID)
			testutil.Equal(t, title, resp.Task.Title)
			testutil.Equal(t, status, resp.Task.Status)
			testutil.Equal(t, now.String(), resp.Task.UpdatedAt)
			testutil.Equal(t, task1.CreatedAt, resp.Task.CreatedAt)

			task1 = resp.Task
		}
		{
			var resp struct {
				Task *Task `json:"task"`
			}
			statusCode, err := c.DoAPI(ctx,
				http.MethodPost,
				"/grpc/app.v1.TaskService/Get",
				struct {
					ID string `json:"id"`
				}{
					ID: task1.ID,
				},
				&resp,
			)
			if err != nil {
				t.Fatal(errors.Wrap(err, "failed to get task"))
			}

			testutil.Equal(t, http.StatusOK, statusCode)
			testutil.Equal(t, task1.ID, resp.Task.ID)
			testutil.Equal(t, task1.Title, resp.Task.Title)
			testutil.Equal(t, task1.Status, resp.Task.Status)
			testutil.Equal(t, task1.UpdatedAt, resp.Task.UpdatedAt)
			testutil.Equal(t, task1.CreatedAt, resp.Task.CreatedAt)
		}
	})

	t.Run("success: delete task1", func(t *testing.T) {
		{
			statusCode, err := c.DoAPI(ctx,
				http.MethodPost,
				"/grpc/app.v1.TaskService/Delete",
				struct {
					ID string `json:"id"`
				}{
					ID: task1.ID,
				},
				nil,
			)
			if err != nil {
				t.Fatal(errors.Wrap(err, "failed to delete task"))
			}

			testutil.Equal(t, http.StatusOK, statusCode)

			task1 = nil
		}
		{
			var resp struct {
				Tasks []*Task `json:"tasks"`
			}
			statusCode, err := c.DoAPI(ctx,
				http.MethodPost,
				"/grpc/app.v1.TaskService/List",
				struct{}{},
				&resp,
			)
			if err != nil {
				t.Fatal(errors.Wrap(err, "failed to list tasks"))
			}

			testutil.Equal(t, http.StatusOK, statusCode)
			testutil.Equal(t, 1, len(resp.Tasks))
			testutil.Equal(t, task2.ID, resp.Tasks[0].ID)
			testutil.Equal(t, task2.Title, resp.Tasks[0].Title)
			testutil.Equal(t, task2.Status, resp.Tasks[0].Status)
			testutil.Equal(t, task2.UpdatedAt, resp.Tasks[0].UpdatedAt)
			testutil.Equal(t, task2.CreatedAt, resp.Tasks[0].CreatedAt)
		}
	})
}
