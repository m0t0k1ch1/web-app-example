package config

type App struct {
	Server Server `yaml:"server" validate:"required"`
	MySQL  MySQL  `yaml:"mysql" validate:"required"`
}
