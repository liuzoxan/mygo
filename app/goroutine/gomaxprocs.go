package main

import (
	"fmt"
	"runtime"
)

func main()  {
	runtime.GOMAXPROCS(1)
	go func() {
		fmt.Println("hello world!")
		// 如果GOMAXPROCS设置为1，不会输出
	}()

	go func() {
		for {
			// 该死循环会一直占用一个Processor
		}
	}()
	select {

	}
}
