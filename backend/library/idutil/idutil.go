package idutil

import (
	"github.com/cockroachdb/errors"
	"github.com/sqids/sqids-go"
)

var (
	codec *sqids.Sqids
)

func init() {
	codec, _ = sqids.New()
}

func Encode(i uint64) string {
	encoded, _ := codec.Encode([]uint64{i})

	return encoded
}

func Decode(s string) (uint64, error) {
	is := codec.Decode(s)
	if len(is) == 0 {
		return 0, errors.New("invalid")
	}
	if len(is) > 1 {
		return 0, errors.New("out of range")
	}

	return is[0], nil
}
