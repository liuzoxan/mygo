package main

import (
	"flag"
	"fmt"
	"hellogo/service"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
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

	//  增加kill信号处理
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt, os.Kill, syscall.SIGTERM)
		sig := <-sigchan
		fmt.Println("get shutdown signal:", sig)
		server.Shutdown(nil)

		//util.WriteDisk()
		//logger.Flush()
		os.Exit(0)
	}()

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("start http server failed: ", err.Error())
		goto EXIT1
	}

	fmt.Println("Should not go here!!!")

EXIT1:
	//util.WriteDisk()
	//logger.Flush()
	os.Exit(1)
}
