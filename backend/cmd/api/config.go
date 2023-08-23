package main

import (
	"github.com/m0t0k1ch1/web-app-sample/backend/config"
)

type Config struct {
	Server config.ServerConfig `json:"server"`
}
