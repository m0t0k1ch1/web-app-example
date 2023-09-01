package appv1

import (
	"backend/handler"
)

type TaskServiceHandler struct {
	*handler.Base
}

func NewTaskServiceHandler(base *handler.Base) *TaskServiceHandler {
	return &TaskServiceHandler{
		Base: base,
	}
}
