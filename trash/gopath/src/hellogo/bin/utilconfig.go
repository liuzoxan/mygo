package main

import (
	"fmt"
	"hellogo/util"
)

func main() {
	Config := util.Config
	Config.LoadConfigFile("./conf/multiconfig.toml")
	fmt.Println(*Config)
}
