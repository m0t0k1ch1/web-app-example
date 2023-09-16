package service

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"

	"app/gen/sqlc/mysql"
	"app/library/idutil"
)

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

func NewUnknownError(err error) error {
	return connect.NewError(connect.CodeUnknown, err)
}

func NewInvalidArgumentError(err error) error {
	return connect.NewError(connect.CodeInvalidArgument, err)
}

func NewNotFoundError(err error) error {
	return connect.NewError(connect.CodeNotFound, err)
}
