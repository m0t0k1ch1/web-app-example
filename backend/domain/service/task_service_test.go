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
			payload := testTaskServiceCreateSuccess(t, ctx, gqlgen.CreateTaskInput{
				Title: "task1.title",
			})

			task1 = payload.Task
		}
	})

	t.Run("success: create task2", func(t *testing.T) {
		{
			payload := testTaskServiceCreateSuccess(t, ctx, gqlgen.CreateTaskInput{
				Title: "task2.title",
			})

			task2 = payload.Task
		}
	})

	t.Run("success: create task3", func(t *testing.T) {
		{
			payload := testTaskServiceCreateSuccess(t, ctx, gqlgen.CreateTaskInput{
				Title: "task3.title",
			})

			task3 = payload.Task
		}
	})

	t.Run("success: create task4", func(t *testing.T) {
		{
			payload := testTaskServiceCreateSuccess(t, ctx, gqlgen.CreateTaskInput{
				Title: "task4.title",
			})

			task4 = payload.Task
		}
	})

	t.Run("success: create task5", func(t *testing.T) {
		{
			payload := testTaskServiceCreateSuccess(t, ctx, gqlgen.CreateTaskInput{
				Title: "task5.title",
			})

			task5 = payload.Task
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
							ID: task1.Id,
							Params: model.PaginationCursorParams{
								TaskStatus: &status,
							},
						}),
						Node: task1,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task2.Id,
							Params: model.PaginationCursorParams{
								TaskStatus: &status,
							},
						}),
						Node: task2,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task3.Id,
							Params: model.PaginationCursorParams{
								TaskStatus: &status,
							},
						}),
						Node: task3,
					},
				},
				PageInfo: &gqlgen.PageInfo{
					EndCursor: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
						ID: task3.Id,
						Params: model.PaginationCursorParams{
							TaskStatus: &status,
						},
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
					ID: task3.Id,
					Params: model.PaginationCursorParams{
						TaskStatus: &status,
					},
				})),
				First: 3,
			})
			require.Nil(t, err)

			testutil.Equal(t, &gqlgen.TaskConnection{
				Edges: []*gqlgen.TaskEdge{
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task4.Id,
							Params: model.PaginationCursorParams{
								TaskStatus: &status,
							},
						}),
						Node: task4,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task5.Id,
							Params: model.PaginationCursorParams{
								TaskStatus: &status,
							},
						}),
						Node: task5,
					},
				},
				PageInfo: &gqlgen.PageInfo{
					EndCursor: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
						ID: task5.Id,
						Params: model.PaginationCursorParams{
							TaskStatus: &status,
						},
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
			payload := testTaskServiceCompleteSuccess(t, ctx, gqlgen.CompleteTaskInput{
				Id: task2.Id,
			})

			task2 = payload.Task
		}
	})

	t.Run("success: complete task4", func(t *testing.T) {
		{
			payload := testTaskServiceCompleteSuccess(t, ctx, gqlgen.CompleteTaskInput{
				Id: task4.Id,
			})

			task4 = payload.Task
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
							ID: task1.Id,
							Params: model.PaginationCursorParams{
								TaskStatus: &status,
							},
						}),
						Node: task1,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task3.Id,
							Params: model.PaginationCursorParams{
								TaskStatus: &status,
							},
						}),
						Node: task3,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task5.Id,
							Params: model.PaginationCursorParams{
								TaskStatus: &status,
							},
						}),
						Node: task5,
					},
				},
				PageInfo: &gqlgen.PageInfo{
					EndCursor: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
						ID: task5.Id,
						Params: model.PaginationCursorParams{
							TaskStatus: &status,
						},
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
							ID: task2.Id,
							Params: model.PaginationCursorParams{
								TaskStatus: &status,
							},
						}),
						Node: task2,
					},
					{
						Cursor: encodePaginationCursor(t, model.PaginationCursor{
							ID: task4.Id,
							Params: model.PaginationCursorParams{
								TaskStatus: &status,
							},
						}),
						Node: task4,
					},
				},
				PageInfo: &gqlgen.PageInfo{
					EndCursor: coreutil.Ptr(encodePaginationCursor(t, model.PaginationCursor{
						ID: task4.Id,
						Params: model.PaginationCursorParams{
							TaskStatus: &status,
						},
					})),
					HasNextPage: false,
				},
				TotalCount: 2,
			}, out.TaskConnection)
		}
	})
}

func testTaskServiceCreateSuccess(t *testing.T, ctx context.Context, input gqlgen.CreateTaskInput) *gqlgen.CreateTaskPayload {
	t.Helper()

	taskService, _ := setUpTaskService(t, nil)

	payload, err := taskService.Create(ctx, input)
	require.Nil(t, err)

	testutil.Equal(t, input.Title, payload.Task.Title)
	testutil.Equal(t, gqlgen.TaskStatusUncompleted, payload.Task.Status)

	return payload
}

func testTaskServiceCompleteSuccess(t *testing.T, ctx context.Context, input gqlgen.CompleteTaskInput) *gqlgen.CompleteTaskPayload {
	t.Helper()

	taskService, _ := setUpTaskService(t, nil)

	payload, err := taskService.Complete(ctx, input)
	require.Nil(t, err)

	testutil.Equal(t, input.Id, payload.Task.Id)
	testutil.Equal(t, gqlgen.TaskStatusCompleted, payload.Task.Status)

	return payload
}
