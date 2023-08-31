package appv1

import (
	"backend/handler"
)

type TaskServiceHandler struct {
	*handler.HandlerBase
}

func NewTaskServiceHandler(base *handler.HandlerBase) *TaskServiceHandler {
	return &TaskServiceHandler{
		HandlerBase: base,
	}
}
