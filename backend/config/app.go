package config

type AppConfig struct {
	MySQL   MySQLConfig   `yaml:"mysql" validate:"required"`
	Runtime RuntimeConfig `yaml:"runtime" validate:"required"`
	Sentry  SentryConfig  `yaml:"sentry"`
	Server  ServerConfig  `yaml:"server" validate:"required"`
}
