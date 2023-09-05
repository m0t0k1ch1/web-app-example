package config

import (
	"fmt"
)

type Server struct {
	Port int `yaml:"port" validate:"required,gte=0,lte=65535"`
}

func (conf Server) Addr() string {
	return fmt.Sprintf(":%d", conf.Port)
}
