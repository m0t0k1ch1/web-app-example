//go:build wireinject
// +build wireinject

package core

import (
	"context"

	"github.com/google/wire"

	"app/config"
)

func InitializeApp(ctx context.Context, conf config.AppConfig) (*App, error) {
	wire.Build(ProviderSet)

	return &App{}, nil
}
