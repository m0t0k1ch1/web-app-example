package appv1

import (
	"app/service"
)

type TaskService struct {
	*service.Base
}

func NewTaskService(base *service.Base) *TaskService {
	return &TaskService{
		Base: base,
	}
}
