package config

type ConfigPath string

func (confPath ConfigPath) String() string {
	return string(confPath)
}
