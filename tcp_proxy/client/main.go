package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		// Handle error
	}

	for {

		time.Sleep(time.Second * 10)

		conn.Write([]byte("tcp"))
		readbuf := make([]byte, 512)
		dataLen, err := conn.Read(readbuf)
		if err != nil {
			log.Fatal(err)
		}

		if dataLen > 0 {
			fmt.Println(string(readbuf))
		}
	}

}
