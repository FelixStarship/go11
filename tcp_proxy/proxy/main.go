package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// Handle error
	}

	// Continually accept incoming connections
	count := 0
	for {
		conn, err := ln.Accept()
		if err != nil {
			// Handle error
		}

		count++
		fmt.Println("connection ID:", count)

		go func(cid int) {
			for {
				// 模拟tcp粘包的场景
				//1.客户端一段时间内发送包的速度太多，服务端没有全部处理完。于是数据就会积压起来，产生粘包。
				//2.定义的读的buffer不够大，而数据包太大或者由于粘包产生，服务端不能一次全部读完，产生半包。
				//3.字节流数据包没有边界、纯裸tcp包
				readbuf := make([]byte, 512)

				dataLen, err := conn.Read(readbuf)

				if err != nil {
					log.Fatal(err)
				}

				if dataLen > 0 {
					fmt.Println("读客户端消息:", string(readbuf), "客户端连接ID:", cid)
				}

				conn.Write([]byte("写入消息!"))

			}
		}(count)
	}
}
