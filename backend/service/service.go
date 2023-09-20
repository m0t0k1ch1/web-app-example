package service

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"

	"app/gen/sqlc/mysql"
)

func GetTaskByDisplayIDOrError(ctx context.Context, db mysql.DBTX, displayID string) (mysql.Task, error) {
	task, err := mysql.New(db).GetTaskByDisplayID(ctx, sql.NullString{
		String: displayID,
		Valid:  true,
	})
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
