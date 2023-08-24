package idutil

import (
	"github.com/pkg/errors"
	"github.com/sqids/sqids-go"
)

var (
	s *sqids.Sqids
)

func init() {
	s, _ = sqids.New()
}

type ID uint64

func (id ID) Uint64() uint64 {
	return uint64(id)
}

func (id ID) Encode() string {
	encoded, _ := s.Encode([]uint64{id.Uint64()})

	return encoded
}

func Decode(encoded string) (ID, error) {
	ids := s.Decode(encoded)
	if len(ids) > 1 {
		return 0, errors.New("out of range")
	}

	return ID(ids[0]), nil
}
