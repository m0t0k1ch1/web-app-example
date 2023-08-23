package config

import (
	"fmt"
	"time"
)

type ServerConfig struct {
	Port            uint16        `json:"port"`
	ShutdownTimeout time.Duration `json:"shutdownTimeout"`
}

func (conf ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", conf.Port)
}
