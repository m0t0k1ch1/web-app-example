package timeutil

import (
	"database/sql/driver"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

var (
	locked Timestamp
)

type Timestamp struct {
	t time.Time
}

func Now() Timestamp {
	if !locked.IsZero() {
		return locked
	}

	return Timestamp{
		t: time.Now(),
	}
}

func (t Timestamp) IsZero() bool {
	return t.t.IsZero()
}

func (t Timestamp) Time() time.Time {
	return t.t
}

func (t Timestamp) Unix() int64 {
	return t.t.Unix()
}

func (t Timestamp) String() string {
	return strconv.FormatInt(t.t.Unix(), 10)
}

func (t Timestamp) Value() (driver.Value, error) {
	return t.t.Unix(), nil
}

func (t *Timestamp) Scan(src any) error {
	switch v := src.(type) {
	case int64:
		t.t = time.Unix(v, 0).In(time.UTC)

	case []byte:
		i, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return errors.Wrapf(err, "failed to convert %s into type int64", string(v))
		}

		t.t = time.Unix(i, 0).In(time.UTC)

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

	t.t = time.Unix(i, 0).In(time.UTC)

	return nil
}

func Lock(t Timestamp) {
	locked = t
}

func Unlock() {
	locked = Timestamp{}
}
