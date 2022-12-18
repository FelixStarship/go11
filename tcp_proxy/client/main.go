package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		// Handle error
	}

	conn.Write([]byte("testettetete"))

	io.Copy(os.Stdout, conn)
}
