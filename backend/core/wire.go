//go:build wireinject
// +build wireinject

package core

import (
	"context"

	"github.com/google/wire"
)

func InitializeApp(ctx context.Context, confPath ConfigPath) (*App, error) {
	wire.Build(ProviderSet)

	return &App{}, nil
}
