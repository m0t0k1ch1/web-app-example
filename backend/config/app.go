package config

type AppConfig struct {
	MySQL  MySQLConfig  `yaml:"mysql" validate:"required"`
	Server ServerConfig `yaml:"server" validate:"required"`
}
