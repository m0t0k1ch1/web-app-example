package sqlutil

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type TxStarter interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

func Transact(ctx context.Context, starter TxStarter, f func(context.Context, *sql.Tx) error) (err error) {
	var tx *sql.Tx
	if tx, err = starter.BeginTx(ctx, nil); err != nil {
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
