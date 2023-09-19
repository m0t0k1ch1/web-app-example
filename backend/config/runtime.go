package config

type RuntimeConfig struct {
	Env string `yaml:"env" validate:"required"`
}
