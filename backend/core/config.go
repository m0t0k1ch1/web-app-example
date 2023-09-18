package core

import (
	"fmt"
)

type ConfigPath string

func (confPath ConfigPath) String() string {
	return string(confPath)
}

type AppConfig struct {
	MySQL  MySQLConfig  `yaml:"mysql" validate:"required"`
	Server ServerConfig `yaml:"server" validate:"required"`
}

type MySQLConfig struct {
	Host     string `yaml:"host" validate:"required,hostname_rfc1123"`
	Port     int    `yaml:"port" validate:"required,gte=1,lte=65535"`
	User     string `yaml:"user" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	DBName   string `yaml:"db_name" validate:"required"`
}

func (conf MySQLConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DBName,
	)
}

type ServerConfig struct {
	Port int `yaml:"port" validate:"required,gte=1,lte=65535"`
}

func (conf ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", conf.Port)
}
