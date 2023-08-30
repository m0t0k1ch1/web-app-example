package config

import (
	"fmt"
)

type ServerConfig struct {
	Port uint16 `yaml:"port" validate:"required"`
}

func (conf ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", conf.Port)
}
