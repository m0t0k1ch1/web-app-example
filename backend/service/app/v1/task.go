package appv1

import (
	"database/sql"
)

type TaskService struct {
	mysql *sql.DB
}

func NewTaskService(mysql *sql.DB) *TaskService {
	return &TaskService{
		mysql: mysql,
	}
}
