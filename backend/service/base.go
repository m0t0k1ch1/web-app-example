package service

import (
	"app/env"
)

type Base struct {
	Env *env.Container
}

func NewBase(envCtr *env.Container) *Base {
	return &Base{
		Env: envCtr,
	}
}
