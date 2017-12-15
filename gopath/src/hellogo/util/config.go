package util

import "github.com/koding/multiconfig"

type APPConfig struct {
	Name         string
	ServerConfig ServerConfig
	LogConfig    LogConfig
}

type ServerConfig struct {
	Port int
}

type LogConfig struct {
	Runtime string
	Status  string
}

var Config *APPConfig

func (config *APPConfig) LoadConfigFile(confpath string) error {
	m := &multiconfig.TOMLLoader{Path: confpath}
	err := m.Load(config)
	return err
}

func init() {
	Config = new(APPConfig)
}
