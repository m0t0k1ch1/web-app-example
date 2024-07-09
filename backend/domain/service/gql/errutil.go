package gql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"app/gql/errcode"
)

func NewBadUserInputError(ctx context.Context, err error) error {
	return newErrorWithStatusCode(ctx, err, errcode.BadUserInput)
}

func NewInternalServerError(ctx context.Context, err error) error {
	return newErrorWithStatusCode(ctx, err, errcode.InternalServerError)
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
