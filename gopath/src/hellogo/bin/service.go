package main

import (
	"flag"
	"fmt"
	"hellogo/service"
	"hellogo/util"
	"net/http"
	"os"
	"runtime"
	"time"
)

var _VERSION_ = "unknown"

func main() {
	//  设置并发度
	runtime.GOMAXPROCS(runtime.NumCPU() * 4) //number of core

	//  处理命令行参数
	confPath := flag.String("c", "./conf/zeusextractor.conf", "configure file path")
	showVersion := flag.Bool("version", false, "show version")

	flag.Parse()
	if *showVersion {
		fmt.Println(_VERSION_)
		os.Exit(0)
	}

	if false {
		fmt.Println(*confPath)
		os.Exit(0)
	}

	//  启动http服务
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/admin", service.DoAdmin)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " ready to start http server")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("start http server failed: ", err.Error())
		goto EXIT1
	}

	//logger.Flush()
	server.Shutdown(nil)
	os.Exit(0)

EXIT1:
	util.WriteDisk()
	//logger.Runtime.Info("close redis sender.")
	//logger.Flush()
	os.Exit(1)
}
