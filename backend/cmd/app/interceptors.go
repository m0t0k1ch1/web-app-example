package main

import (
	"context"

	"connectrpc.com/connect"
)

var ValidationInterceptor = connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		v, ok := req.Any().(interface {
			ValidateAll() error
		})
		if ok {
			if err := v.ValidateAll(); err != nil {
				return nil, connect.NewError(connect.CodeInvalidArgument, err)
			}
		}

		return next(ctx, req)
	})
})
