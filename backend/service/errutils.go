package service

import (
	"connectrpc.com/connect"
)

func NewUnknownError(err error) error {
	return connect.NewError(connect.CodeUnknown, err)
}

func NewInvalidArgumentError(err error) error {
	return connect.NewError(connect.CodeInvalidArgument, err)
}

func NewNotFoundError(err error) error {
	return connect.NewError(connect.CodeNotFound, err)
}
