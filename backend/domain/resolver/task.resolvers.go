package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"app/gen/gqlgen"
	"context"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input gqlgen.CreateTaskInput) (*gqlgen.CreateTaskPayload, error) {
	return r.taskService.Create(ctx, input)
}

// CompleteTask is the resolver for the completeTask field.
func (r *mutationResolver) CompleteTask(ctx context.Context, input gqlgen.CompleteTaskInput) (*gqlgen.CompleteTaskPayload, error) {
	return r.taskService.Complete(ctx, input)
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, status *gqlgen.TaskStatus, after *string, first int32) (*gqlgen.TaskConnection, error) {
	return r.taskService.List(ctx, status, after, first)
}
