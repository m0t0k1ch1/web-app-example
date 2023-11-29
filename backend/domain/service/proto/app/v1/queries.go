package appv1

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"app/domain/service/proto"
	"app/gen/sqlc/mysql"
	"app/library/idutil"
)

func GetTaskOrError(ctx context.Context, db mysql.DBTX, encodedID string) (mysql.Task, error) {
	resourceName, id, err := idutil.Decode(encodedID)
	if err != nil {
		return mysql.Task{}, proto.NewInvalidArgumentError(errors.Wrap(err, "failed to decode id"))
	}
	if resourceName != ResourceNameTask {
		return mysql.Task{}, proto.NewInvalidArgumentError(errors.Errorf("unexpected resource name: %s", resourceName))
	}

	task, err := mysql.New(db).GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mysql.Task{}, proto.NewNotFoundError(errors.New("task not found"))
		}

		return mysql.Task{}, proto.NewUnknownError(errors.Wrap(err, "failed to get task"))
	}

	return task, nil
}
