package handler

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"backend/core"
	"backend/gen/sqlc/mysql"
	"backend/library/idutil"
)

type HandlerBase struct {
	Env *core.Env
}

func NewHandlerBase(env *core.Env) *HandlerBase {
	return &HandlerBase{
		Env: env,
	}
}

func (h *HandlerBase) MustGetTask(ctx context.Context, id idutil.ID) (mysql.Task, error) {
	task, err := mysql.New(h.Env.DB).GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mysql.Task{}, NewNotFoundError(errors.Wrap(err, "task not found"))
		}

		return mysql.Task{}, NewUnknownError(errors.Wrap(err, "failed to get task"))
	}

	return task, nil
}
