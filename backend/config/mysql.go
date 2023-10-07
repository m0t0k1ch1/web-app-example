package config

import (
	"fmt"
)

type MySQLConfig struct {
	App MySQLDBConfig `yaml:"app" validate:"required"`
}

type MySQLDBConfig struct {
	Host     string `yaml:"host" validate:"required,hostname_rfc1123"`
	Port     int    `yaml:"port" validate:"required,gte=1,lte=65535"`
	User     string `yaml:"user" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Name     string `yaml:"name" validate:"required"`
}

func (conf MySQLDBConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name,
	)
}
