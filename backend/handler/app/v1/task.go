package appv1

import (
	"backend/handler"
)

type TaskService struct {
	*handler.Base
}

func NewTaskService(base *handler.Base) *TaskService {
	return &TaskService{
		Base: base,
	}
}
