package nodeid

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/samber/oops"
)

const (
	TypeTask = "Task"

	separator = ":"
)

var (
	enc = base64.URLEncoding
)

type Type string

func (t Type) String() string {
	return string(t)
}

func Encode(id uint64, t Type) string {
	return enc.EncodeToString([]byte(t.String() + separator + strconv.FormatUint(id, 10)))
}

func Decode(encodedID string) (uint64, Type, error) {
	b, err := enc.DecodeString(encodedID)
	if err != nil {
		return 0, "", err
	}

	parts := strings.Split(string(b), separator)
	if len(parts) != 2 {
		return 0, "", oops.Errorf("invalid format")
	}

	id, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return 0, "", oops.Errorf("invalid format")
	}

	return id, Type(parts[0]), nil
}

func DecodeByType(encodedID string, expectedType Type) (uint64, error) {
	id, t, err := Decode(encodedID)
	if err != nil {
		return 0, err
	}
	if t != expectedType {
		return 0, oops.Errorf("unexpected type: %s", t)
	}

	return id, nil
}
