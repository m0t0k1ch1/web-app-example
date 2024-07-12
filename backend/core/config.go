package core

import (
	configloader "github.com/kayac/go-config"
	"github.com/samber/oops"

	"app/config"
	"app/domain/validation"
)

func init() {
	configloader.Delims("<%", "%>")
}

func LoadAppConfig(confPath string) (config.AppConfig, error) {
	var conf config.AppConfig
	if err := configloader.LoadWithEnv(&conf, confPath); err != nil {
		return config.AppConfig{}, oops.Wrapf(err, "failed to load config")
	}

	if err := validation.Struct(conf); err != nil {
		return config.AppConfig{}, oops.Wrapf(err, "invalid config")
	}

	return conf, nil
}
