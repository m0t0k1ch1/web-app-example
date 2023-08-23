package main

import (
	kayac_config "github.com/kayac/go-config"

	"github.com/m0t0k1ch1/web-app-sample/backend/config"
)

type Config struct {
	Server config.ServerConfig `json:"server"`
}

func LoadConfig(path string) (Config, error) {
	kayac_config.Delims("<%", "%>")

	var conf Config
	if err := kayac_config.LoadWithEnv(&conf, path); err != nil {
		return Config{}, err
	}

	return conf, nil
}
