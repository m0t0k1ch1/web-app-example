package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	configloader "github.com/kayac/go-config"
	"github.com/pkg/errors"

	"app/db"
)

type Config struct {
	DB     db.Config    `yaml:"db" validate:"required"`
	Server ServerConfig `yaml:"server" validate:"required"`
}

type ServerConfig struct {
	Port int `yaml:"port" validate:"required,gte=1,lte=65535"`
}

func (conf ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", conf.Port)
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
