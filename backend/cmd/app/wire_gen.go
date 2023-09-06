// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"app/db"
	"app/env"
	"context"
)

// Injectors from wire.go:

func InitializeApp(ctx context.Context, conf Config) (*App, error) {
	serverConfig := conf.Server
	config := conf.DB
	container, err := db.NewContainer(config)
	if err != nil {
		return nil, err
	}
	envContainer := env.NewContainer(container)
	server := NewServer(serverConfig, envContainer)
	app := NewApp(server)
	return app, nil
}
