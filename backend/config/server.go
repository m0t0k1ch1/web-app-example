package config

import (
	"fmt"
)

type ServerConfig struct {
	Port uint16 `json:"port"`
}

func (conf ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", conf.Port)
}
