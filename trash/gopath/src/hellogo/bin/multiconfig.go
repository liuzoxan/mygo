package main

import (
	"fmt"
	"github.com/koding/multiconfig"
)

type APPConfig struct {
	Name         string `required:"true"`
	ServerConfig ServerConfig
	LogConfig    LogConfig
}

type ServerConfig struct {
	Port int `default:"6060"`
}

type LogConfig struct {
	Runtime string `required:"true"`
	Status  string `required:"true"`
}

func main() {
	// Create a new DefaultLoader without or with an initial config file
	//m := multiconfig.New()
	mc := multiconfig.NewWithPath("./conf/multiconfig.toml") // supports TOML, JSON and YAML

	// Get an empty struct for your configuration
	serverConf := new(APPConfig)

	// Populated the serverConf struct
	//err := m.Load(serverConf) // Check for error
	mc.MustLoad(serverConf) // Panic's if there is any error

	// Access now populated fields
	fmt.Println(serverConf.Name, serverConf.ServerConfig.Port, serverConf.LogConfig)
}
