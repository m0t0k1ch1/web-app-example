package main

import (
	kayac_config "github.com/kayac/go-config"

	"backend/config"
)

type Config struct {
	Server config.ServerConfig `yaml:"server"`
	MySQL  config.MySQLConfig  `yaml:"mysql"`
}

func LoadConfig(path string) (Config, error) {
	kayac_config.Delims("<%", "%>")

	var conf Config
	if err := kayac_config.LoadWithEnv(&conf, path); err != nil {
		return Config{}, err
	}

	return conf, nil
}
