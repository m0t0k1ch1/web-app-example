// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package core

import (
	"app/config"
	"context"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeApp(ctx context.Context, conf config.AppConfig) (*App, error) {
	clock := provideClock()
	mySQLContainer, err := provideMySQLContainer(conf)
	if err != nil {
		return nil, err
	}
	taskService := provideTaskService(clock, mySQLContainer)
	nodeService := provideNodeService(mySQLContainer)
	resolver := provideGQLResolver(taskService, nodeService)
	server := provideServer(conf, resolver)
	app := provideApp(conf, server)
	return app, nil
}
