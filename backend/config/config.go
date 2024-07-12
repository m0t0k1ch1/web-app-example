package config

import (
	"fmt"
)

type AppConfig struct {
	MySQL  MySQLConfig  `yaml:"mysql" validate:"required" en:"mysql"`
	Server ServerConfig `yaml:"server" validate:"required" en:"server"`
}

type MySQLConfig struct {
	App MySQLDBConfig `yaml:"app" validate:"required" en:"app"`
}

type MySQLDBConfig struct {
	Host     string `yaml:"host" validate:"required,hostname_rfc1123" en:"host"`
	Port     int    `yaml:"port" validate:"gte=1,lte=65535" en:"port"`
	User     string `yaml:"user" validate:"required" en:"user"`
	Password string `yaml:"password" validate:"required" en:"password"`
	Name     string `yaml:"name" validate:"required" en:"name"`
}

func (conf MySQLDBConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name,
	)
}

type ServerConfig struct {
	Port           int  `yaml:"port" validate:"gte=1,lte=65535" en:"port"`
	WithPlayground bool `yaml:"with_playground" en:"with_playground"`
}

func (conf ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", conf.Port)
}
