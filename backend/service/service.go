package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"app/env"
	"app/gen/sqlc/mysql"
	"app/library/idutil"
)

type Base struct {
	Env *env.Container
}

func NewBase(env *env.Container) *Base {
	return &Base{
		Env: env,
	}
}

func GetTaskOrError(ctx context.Context, db mysql.DBTX, id idutil.ID) (mysql.Task, error) {
	task, err := mysql.New(db).GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mysql.Task{}, NewNotFoundError(errors.Wrap(err, "task not found"))
		}

		return mysql.Task{}, NewUnknownError(errors.Wrap(err, "failed to get task"))
	}

	return task, nil
}
