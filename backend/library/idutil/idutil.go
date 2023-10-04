package idutil

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	separator = ":"
)

var (
	enc = base64.RawURLEncoding
)

func Encode(prefix string, id uint64) string {
	return enc.EncodeToString([]byte(prefix + separator + strconv.FormatUint(id, 10)))
}

func Decode(s string) (string, uint64, error) {
	decoded, err := enc.DecodeString(s)
	if err != nil {
		return "", 0, err
	}

	parts := strings.Split(string(decoded), separator)
	if len(parts) != 2 {
		return "", 0, errors.New("invalid format")
	}

	id, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return "", 0, errors.New("invalid format")
	}

	return parts[0], id, nil
}
