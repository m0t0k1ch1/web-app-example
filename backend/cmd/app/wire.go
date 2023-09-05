//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"

	"app/db"
	"app/env"
	"app/service"
	appv1 "app/service/app/v1"
)

func InitializeApp(ctx context.Context, conf Config) (*App, error) {
	wire.Build(
		wire.FieldsOf(new(Config), "DB", "Server"),

		db.NewContainer,
		env.NewContainer,

		service.NewBase,
		appv1.NewTaskService,

		NewServer,

		NewApp,
	)

	return &App{}, nil
}
