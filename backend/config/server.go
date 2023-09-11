package config

import (
	"fmt"
)

type ServerConfig struct {
	Port int `yaml:"port" validate:"required,gte=1,lte=65535"`
}

func (conf ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", conf.Port)
}
