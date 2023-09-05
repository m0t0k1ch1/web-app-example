package service

import (
	"app/env"
)

type Base struct {
	Env *env.Container
}

func NewBase(env *env.Container) *Base {
	return &Base{
		Env: env,
	}
}
