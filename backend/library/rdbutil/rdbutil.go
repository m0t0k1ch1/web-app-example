package rdbutil

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

func Transact(ctx context.Context, db *sql.DB, f func(context.Context, *sql.Tx) error) (err error) {
	var tx *sql.Tx
	if tx, err = db.BeginTx(ctx, nil); err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			if err = tx.Commit(); err != nil {
				err = errors.Wrap(err, "failed to commit transaction")
			}
		}
	}()

	err = f(ctx, tx)

	return
}