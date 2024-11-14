package e

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"app/gen/gqlgen"
)

func NewGQLError(ctx context.Context, err error, code gqlgen.ErrorCode) *gqlerror.Error {
	return &gqlerror.Error{
		Err:     err,
		Message: err.Error(),
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]any{
			"code": code,
		},
	}
}

func NewUnexpectedGQLError(ctx context.Context, err error) *gqlerror.Error {
	return &gqlerror.Error{
		Err:     err,
		Message: "something went wrong",
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]any{
			"code": gqlgen.ErrorCodeInternalServerError,
		},
	}
}
