package config

import ()

type Config struct {
	ProjectName string
}

const (
	VERSION = "0.1"
)

var Cfg *Config = &Config{}