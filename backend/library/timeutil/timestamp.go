package timeutil

import (
	"database/sql/driver"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Timestamp struct {
	time.Time
}

func Now() Timestamp {
	return Timestamp{time.Now()}
}

func (t Timestamp) Value() (driver.Value, error) {
	return t.Unix(), nil
}

func (t *Timestamp) Scan(src any) error {
	switch v := src.(type) {
	case int64:
		t.Time = time.Unix(v, 0).In(time.UTC)

	case []byte:
		i, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return errors.Wrapf(err, "failed to convert %s into type int64", string(v))
		}

		t.Time = time.Unix(i, 0).In(time.UTC)

	default:
		return errors.Errorf("unexpected src type: %T", src)
	}

	return nil
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.Unix(), 10)), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return errors.Wrapf(err, "failed to convert %s into type int64", string(b))
	}

	t.Time = time.Unix(i, 0).In(time.UTC)

	return nil
}
