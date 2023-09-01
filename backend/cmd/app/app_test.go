package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pkg/errors"

	"backend/config"
	appv1 "backend/gen/buf/app/v1"
	"backend/internal/testutil"
)

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

	t.Run("success: no tasks", func(t *testing.T) {
		{
			var resp appv1.TaskServiceListResponse
			statusCode, err := c.DoAPI(ctx,
				http.MethodPost,
				"/grpc/app.v1.TaskService/List",
				appv1.TaskServiceListRequest{},
				&resp,
			)
			if err != nil {
				t.Fatal(errors.Wrap(err, "failed to list tasks"))
			}
			testutil.Equal(t, http.StatusOK, statusCode)
			testutil.Equal(t, []*appv1.Task{}, resp.Tasks)
		}
	})
}
