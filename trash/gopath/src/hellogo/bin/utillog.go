package main

import (
	"fmt"
	"hellogo/util"
	"os"
)

func main() {
	logger := util.Logger
	if err := logger.Init("conf/seelog_runtime.xml", "conf/seelog_status.xml"); err != nil {
		fmt.Println("init log failed: ", err.Error())
		os.Exit(0)
	}

	logger.Runtime.Info("Runtime log")
	logger.StatusLog.Info("Status log")
	logger.Flush()
}
