package main

import _ "net"

func main() {
	//通道尚未初始化，即零值通道
	var ch chan struct{}
	close(ch)
}
