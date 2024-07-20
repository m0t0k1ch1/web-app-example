package model

import (
	"encoding/base64"
	"encoding/json"

	"app/gen/gqlgen"
)

type PaginationCursor struct {
	ID     string                 `json:"id"`
	Params PaginationCursorParams `json:"params"`
}

type PaginationCursorParams struct {
	TaskStatus *gqlgen.TaskStatus `json:"taskStatus,omitempty"`
}

func DecodePaginationCursor(encoded string) (PaginationCursor, error) {
	b, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return PaginationCursor{}, err
	}

	var cursor PaginationCursor
	{
		if err := json.Unmarshal(b, &cursor); err != nil {
			return PaginationCursor{}, err
		}
	}

	return cursor, nil
}

func (cursor PaginationCursor) Encode() (string, error) {
	b, err := json.Marshal(cursor)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
