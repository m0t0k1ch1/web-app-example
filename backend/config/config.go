package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

type AppConfig struct {
	MySQL  MySQLConfig  `yaml:"mysql" validate:"required"`
	Server ServerConfig `yaml:"server" validate:"required"`
}

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

type ServerConfig struct {
	Port int `yaml:"port" validate:"required,gte=1,lte=65535"`
}

func (conf ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", conf.Port)
}
