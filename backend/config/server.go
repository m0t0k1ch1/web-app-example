package config

import (
	"fmt"
)

type Server struct {
	Port uint16 `yaml:"port" validate:"required"`
}

func (conf Server) Addr() string {
	return fmt.Sprintf(":%d", conf.Port)
}
