package gql

import (
	appv1 "app/domain/service/gql/app/v1"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	taskService *appv1.TaskService
	nodeService *appv1.NodeService
}

func NewResolver(
	taskService *appv1.TaskService,
	nodeService *appv1.NodeService,
) *Resolver {
	return &Resolver{
		taskService: taskService,
		nodeService: nodeService,
	}
}
