package main

import (
	"github.com/go-playground/validator/v10"
	configloader "github.com/kayac/go-config"
	"github.com/pkg/errors"

	"app/config"
)

type Config struct {
	Server config.Server `yaml:"server" validate:"required"`
	MySQL  config.MySQL  `yaml:"mysql" validate:"required"`
}

func LoadConfig(path string) (Config, error) {
	configloader.Delims("<%", "%>")

	var conf Config
	if err := configloader.LoadWithEnv(&conf, path); err != nil {
		return Config{}, err
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(conf); err != nil {
		return Config{}, errors.Wrap(err, "invalid config")
	}

	return conf, nil
}
