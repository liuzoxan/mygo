package main

import (
	"fmt"
	"hellogo/util"
	"os"
	"time"
)

func main() {
	logger := util.Logger
	if err := logger.Init("conf/seelog_runtime.xml", "conf/seelog_status.xml"); err != nil {
		fmt.Println("init log failed: ", err.Error())
		os.Exit(0)
	}

	util.RunWatch()
	util.AddWatchValue("SomeData", 1)
	go func() {
		for {
			select {
			case <-time.Tick(time.Second):
				logger.Runtime.Debug("Heartbeat")
			}
		}
	}()

	time.Sleep(time.Second * 8)
	logger.Flush()
}
