//go:build wireinject
// +build wireinject

package core

import (
	"context"

	"github.com/google/wire"

	"app/config"
)

func InitializeApp(ctx context.Context, confPath config.ConfigPath) (*App, error) {
	wire.Build(ProviderSet)

	return &App{}, nil
}
