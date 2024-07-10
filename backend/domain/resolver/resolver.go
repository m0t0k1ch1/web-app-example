package resolver

import "app/domain/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	taskService *service.TaskService
	nodeService *service.NodeService
}

func NewResolver(
	taskService *service.TaskService,
	nodeService *service.NodeService,
) *Resolver {
	return &Resolver{
		taskService: taskService,
		nodeService: nodeService,
	}
}
