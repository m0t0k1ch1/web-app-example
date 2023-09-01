package config

import (
	"fmt"
)

type MySQL struct {
	Host     string `yaml:"host" validate:"required,hostname_rfc1123"`
	Port     uint16 `yaml:"port" validate:"required"`
	User     string `yaml:"user" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	DBName   string `yaml:"db_name" validate:"required"`
}

func (conf MySQL) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DBName,
	)
}
