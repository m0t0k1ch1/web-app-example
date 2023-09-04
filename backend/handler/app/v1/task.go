package appv1

import (
	"app/handler"
)

type TaskService struct {
	*handler.Base
}

func NewTaskService(base *handler.Base) *TaskService {
	return &TaskService{
		Base: base,
	}
}
