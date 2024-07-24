package gqlerrutil

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	CodeBadUserInput        = "BAD_USER_INPUT"
	CodeInternalServerError = "INTERNAL_SERVER_ERROR"
)

func NewBadUserInputError(ctx context.Context, err error) error {
	return newErrorWithStatusCode(ctx, err, CodeBadUserInput)
}

func newErrorWithStatusCode(ctx context.Context, err error, code string) error {
	return &gqlerror.Error{
		Err:     err,
		Message: err.Error(),
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]any{
			"code": code,
		},
	}
}
