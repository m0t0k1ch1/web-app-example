package config

import (
	"fmt"
)

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
