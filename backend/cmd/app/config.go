package main

import (
	"github.com/go-playground/validator/v10"
	kayac_config "github.com/kayac/go-config"
	"github.com/pkg/errors"

	"backend/config"
)

type Config struct {
	Server config.ServerConfig `yaml:"server" validate:"required"`
	MySQL  config.MySQLConfig  `yaml:"mysql" validate:"required"`
}

func LoadConfig(path string) (Config, error) {
	kayac_config.Delims("<%", "%>")

	var conf Config
	if err := kayac_config.LoadWithEnv(&conf, path); err != nil {
		return Config{}, err
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	if err := v.Struct(conf); err != nil {
		return Config{}, errors.Wrap(err, "invalid config")
	}

	return conf, nil
}
