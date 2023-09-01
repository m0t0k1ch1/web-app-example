package main

import (
	"backend/config"

	"github.com/go-playground/validator/v10"
	kayac_config "github.com/kayac/go-config"
	"github.com/pkg/errors"
)

func LoadConfig(path string) (config.App, error) {
	kayac_config.Delims("<%", "%>")

	var conf config.App
	if err := kayac_config.LoadWithEnv(&conf, path); err != nil {
		return config.App{}, err
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	if err := v.Struct(conf); err != nil {
		return config.App{}, errors.Wrap(err, "invalid config")
	}

	return conf, nil
}
