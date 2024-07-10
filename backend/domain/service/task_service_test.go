package service_test

import (
	"context"
	"testing"

	"github.com/m0t0k1ch1-go/coreutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"app/domain/model"
	"app/domain/service"
	"app/gen/gqlgen"
	"app/internal/testutil"
)

func setUpTaskService(t *testing.T, _ *gomock.Controller) (*service.TaskService, *Mocks) {
	t.Helper()

	return service.NewTaskService(
		clock,
		mysqlCtr,
	), &Mocks{}
}

func TestTaskService(t *testing.T) {
	setup(t)
	t.Cleanup(func() {
		teardown(t)
	})

	ctx := context.Background()

	mockCtrl := gomock.NewController(t)

	var (
		task1 *gqlgen.Task
		task2 *gqlgen.Task
		task3 *gqlgen.Task
		task4 *gqlgen.Task
		task5 *gqlgen.Task
	)

	t.Run("success: create task1", func(t *testing.T) {
		{
			out := testTaskServiceCreateSuccess(t, ctx, service.TaskServiceCreateInput{
				Title: "task1.title",
			})

			task1 = out.Task
		}
	})

	t.Run("success: create task2", func(t *testing.T) {
		{
			out := testTaskServiceCreateSuccess(t, ctx, service.TaskServiceCreateInput{
				Title: "task2.title",
			})

			task2 = out.Task
		}
	})

	t.Run("success: create task3", func(t *testing.T) {
		{
			out := testTaskServiceCreateSuccess(t, ctx, service.TaskServiceCreateInput{
				Title: "task3.title",
			})

			task3 = out.Task
		}
	})

	t.Run("success: create task4", func(t *testing.T) {
		{
			out := testTaskServiceCreateSuccess(t, ctx, service.TaskServiceCreateInput{
				Title: "task4.title",
			})

			task4 = out.Task
		}
	})

	t.Run("success: create task5", func(t *testing.T) {
		{
			out := testTaskServiceCreateSuccess(t, ctx, service.TaskServiceCreateInput{
				Title: "task5.title",
			})

			task5 = out.Task
		}
	})

	t.Run("success: list tasks", func(t *testing.T) {
		{
			taskService, _ := setUpTaskService(t, mockCtrl)

			out, err := taskService.List(ctx, service.TaskServiceListInput{
				First: 3,
			})
			require.Nil(t, err)

			testutil.Equal(t, &gqlgen.TaskConnection{
				Edges: []*gqlgen.TaskEdge{
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task1.Id,
						}),
						Node: task1,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task2.Id,
						}),
						Node: task2,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task3.Id,
						}),
						Node: task3,
					},
				},
				Nodes: []*gqlgen.Task{
					task1,
					task2,
					task3,
				},
				PageInfo: &gqlgen.PageInfo{
					EndCursor: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
						ID: task3.Id,
					})),
					HasNextPage: true,
				},
				TotalCount: 5,
			}, out.TaskConnection)
		}
		{
			taskService, _ := setUpTaskService(t, mockCtrl)

			out, err := taskService.List(ctx, service.TaskServiceListInput{
				After: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
					ID: task3.Id,
				})),
				First: 3,
			})
			require.Nil(t, err)

			testutil.Equal(t, &gqlgen.TaskConnection{
				Edges: []*gqlgen.TaskEdge{
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task4.Id,
						}),
						Node: task4,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task5.Id,
						}),
						Node: task5,
					},
				},
				Nodes: []*gqlgen.Task{
					task4,
					task5,
				},
				PageInfo: &gqlgen.PageInfo{
					EndCursor: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
						ID: task5.Id,
					})),
					HasNextPage: false,
				},
				TotalCount: 5,
			}, out.TaskConnection)
		}
	})

	t.Run("success: list tasks by status: uncompleted", func(t *testing.T) {
		{
			var (
				status = gqlgen.TaskStatusUncompleted
			)

			taskService, _ := setUpTaskService(t, mockCtrl)

			out, err := taskService.List(ctx, service.TaskServiceListInput{
				Status: &status,
				First:  3,
			})
			require.Nil(t, err)

			testutil.Equal(t, &gqlgen.TaskConnection{
				Edges: []*gqlgen.TaskEdge{
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID:         task1.Id,
							TaskStatus: &status,
						}),
						Node: task1,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID:         task2.Id,
							TaskStatus: &status,
						}),
						Node: task2,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID:         task3.Id,
							TaskStatus: &status,
						}),
						Node: task3,
					},
				},
				Nodes: []*gqlgen.Task{
					task1,
					task2,
					task3,
				},
				PageInfo: &gqlgen.PageInfo{
					EndCursor: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
						ID:         task3.Id,
						TaskStatus: &status,
					})),
					HasNextPage: true,
				},
				TotalCount: 5,
			}, out.TaskConnection)
		}
		{
			var (
				status = gqlgen.TaskStatusUncompleted
			)

			taskService, _ := setUpTaskService(t, mockCtrl)

			out, err := taskService.List(ctx, service.TaskServiceListInput{
				Status: &status,
				After: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
					ID:         task3.Id,
					TaskStatus: &status,
				})),
				First: 3,
			})
			require.Nil(t, err)

			testutil.Equal(t, &gqlgen.TaskConnection{
				Edges: []*gqlgen.TaskEdge{
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID:         task4.Id,
							TaskStatus: &status,
						}),
						Node: task4,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID:         task5.Id,
							TaskStatus: &status,
						}),
						Node: task5,
					},
				},
				Nodes: []*gqlgen.Task{
					task4,
					task5,
				},
				PageInfo: &gqlgen.PageInfo{
					EndCursor: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
						ID:         task5.Id,
						TaskStatus: &status,
					})),
					HasNextPage: false,
				},
				TotalCount: 5,
			}, out.TaskConnection)
		}
	})

	t.Run("success: list tasks by status: completed", func(t *testing.T) {
		{
			var (
				status = gqlgen.TaskStatusCompleted
			)

			taskService, _ := setUpTaskService(t, mockCtrl)

			out, err := taskService.List(ctx, service.TaskServiceListInput{
				Status: &status,
				First:  3,
			})
			require.Nil(t, err)

			testutil.Equal(t, &gqlgen.TaskConnection{
				Edges: []*gqlgen.TaskEdge{},
				Nodes: []*gqlgen.Task{},
				PageInfo: &gqlgen.PageInfo{
					EndCursor:   nil,
					HasNextPage: false,
				},
				TotalCount: 0,
			}, out.TaskConnection)
		}
	})

	t.Run("success: complete task2", func(t *testing.T) {
		{
			out := testTaskServiceCompleteSuccess(t, ctx, service.TaskServiceCompleteInput{
				ID: task2.Id,
			})

			task2 = out.Task
		}
	})

	t.Run("success: complete task4", func(t *testing.T) {
		{
			out := testTaskServiceCompleteSuccess(t, ctx, service.TaskServiceCompleteInput{
				ID: task4.Id,
			})

			task4 = out.Task
		}
	})

	t.Run("success: list tasks", func(t *testing.T) {
		{
			taskService, _ := setUpTaskService(t, mockCtrl)

			out, err := taskService.List(ctx, service.TaskServiceListInput{
				First: 3,
			})
			require.Nil(t, err)

			testutil.Equal(t, &gqlgen.TaskConnection{
				Edges: []*gqlgen.TaskEdge{
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task1.Id,
						}),
						Node: task1,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task2.Id,
						}),
						Node: task2,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task3.Id,
						}),
						Node: task3,
					},
				},
				Nodes: []*gqlgen.Task{
					task1,
					task2,
					task3,
				},
				PageInfo: &gqlgen.PageInfo{
					EndCursor: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
						ID: task3.Id,
					})),
					HasNextPage: true,
				},
				TotalCount: 5,
			}, out.TaskConnection)
		}
		{
			taskService, _ := setUpTaskService(t, mockCtrl)

			out, err := taskService.List(ctx, service.TaskServiceListInput{
				After: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
					ID: task3.Id,
				})),
				First: 3,
			})
			require.Nil(t, err)

			testutil.Equal(t, &gqlgen.TaskConnection{
				Edges: []*gqlgen.TaskEdge{
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task4.Id,
						}),
						Node: task4,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task5.Id,
						}),
						Node: task5,
					},
				},
				Nodes: []*gqlgen.Task{
					task4,
					task5,
				},
				PageInfo: &gqlgen.PageInfo{
					EndCursor: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
						ID: task5.Id,
					})),
					HasNextPage: false,
				},
				TotalCount: 5,
			}, out.TaskConnection)
		}
	})

	t.Run("success: list tasks by status: uncompleted", func(t *testing.T) {
		{
			var (
				status = gqlgen.TaskStatusUncompleted
			)

			taskService, _ := setUpTaskService(t, mockCtrl)

			out, err := taskService.List(ctx, service.TaskServiceListInput{
				Status: &status,
				First:  3,
			})
			require.Nil(t, err)

			testutil.Equal(t, &gqlgen.TaskConnection{
				Edges: []*gqlgen.TaskEdge{
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID:         task1.Id,
							TaskStatus: &status,
						}),
						Node: task1,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID:         task3.Id,
							TaskStatus: &status,
						}),
						Node: task3,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID:         task5.Id,
							TaskStatus: &status,
						}),
						Node: task5,
					},
				},
				Nodes: []*gqlgen.Task{
					task1,
					task3,
					task5,
				},
				PageInfo: &gqlgen.PageInfo{
					EndCursor: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
						ID:         task5.Id,
						TaskStatus: &status,
					})),
					HasNextPage: false,
				},
				TotalCount: 3,
			}, out.TaskConnection)
		}
	})

	t.Run("success: list tasks by status: completed", func(t *testing.T) {
		{
			var (
				status = gqlgen.TaskStatusCompleted
			)

			taskService, _ := setUpTaskService(t, mockCtrl)

			out, err := taskService.List(ctx, service.TaskServiceListInput{
				Status: &status,
				First:  3,
			})
			require.Nil(t, err)

			testutil.Equal(t, &gqlgen.TaskConnection{
				Edges: []*gqlgen.TaskEdge{
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID:         task2.Id,
							TaskStatus: &status,
						}),
						Node: task2,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID:         task4.Id,
							TaskStatus: &status,
						}),
						Node: task4,
					},
				},
				Nodes: []*gqlgen.Task{
					task2,
					task4,
				},
				PageInfo: &gqlgen.PageInfo{
					EndCursor: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
						ID:         task4.Id,
						TaskStatus: &status,
					})),
					HasNextPage: false,
				},
				TotalCount: 2,
			}, out.TaskConnection)
		}
	})
}

func testTaskServiceCreateSuccess(t *testing.T, ctx context.Context, in service.TaskServiceCreateInput) service.TaskServiceCreateOutput {
	t.Helper()

	taskService, _ := setUpTaskService(t, nil)

	out, err := taskService.Create(ctx, in)
	require.Nil(t, err)

	testutil.Equal(t, in.Title, out.Task.Title)
	testutil.Equal(t, gqlgen.TaskStatusUncompleted, out.Task.Status)

	return out
}

func testTaskServiceCompleteSuccess(t *testing.T, ctx context.Context, in service.TaskServiceCompleteInput) service.TaskServiceCompleteOutput {
	t.Helper()

	taskService, _ := setUpTaskService(t, nil)

	out, err := taskService.Complete(ctx, in)
	require.Nil(t, err)

	testutil.Equal(t, in.ID, out.Task.Id)
	testutil.Equal(t, gqlgen.TaskStatusCompleted, out.Task.Status)

	return out
}
