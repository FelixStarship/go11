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
