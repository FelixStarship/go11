package main

import (
	"io"
	"net"
)

func main() {
	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// Handle error
	}

	// Continually accept incoming connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			// Handle error
		}

		// Create outgoing connection
		out, err := net.Dial("tcp", "www.example.com:80")
		if err != nil {
			// Handle error
		}

		// Use goroutines to proxy data between connections
		go io.Copy(out, conn)
		go io.Copy(conn, out)
	}
}
