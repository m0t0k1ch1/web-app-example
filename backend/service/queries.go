package service

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"app/gen/sqlc/mysql"
	"app/library/idutil"
)

func GetTaskOrError(ctx context.Context, db mysql.DBTX, encodedID string) (mysql.Task, error) {
	resourceName, id, err := idutil.Decode(encodedID)
	if err != nil {
		return mysql.Task{}, NewInvalidArgumentError(errors.Wrap(err, "failed to decode id"))
	}
	if resourceName != ResourceNameTask {
		return mysql.Task{}, NewInvalidArgumentError(errors.Errorf("unexpected resource name: %s", resourceName))
	}

	task, err := mysql.New(db).GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mysql.Task{}, NewNotFoundError(errors.Wrap(err, "task not found"))
		}

		return mysql.Task{}, NewUnknownError(errors.Wrap(err, "failed to get task"))
	}

	return task, nil
}
