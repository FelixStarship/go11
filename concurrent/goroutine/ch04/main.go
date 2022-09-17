package main

import (
	"fmt"
	_ "net"
	"runtime"
)

func main() {
	//通道尚未初始化，即零值通道
	var ch chan struct{}
	for c := range ch {
		fmt.Println(c)
	}
	fmt.Println(runtime.NumGoroutine())
	<-ch
	close(ch)
}
