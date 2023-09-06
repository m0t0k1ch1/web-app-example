//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"app/db"
	"app/env"
)

func InitializeApp(conf Config) (*App, error) {
	wire.Build(
		wire.FieldsOf(new(Config), "DB", "Server"),

		db.NewContainer,
		env.NewContainer,

		NewServer,

		NewApp,
	)

	return &App{}, nil
}
