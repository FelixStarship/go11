package main

import (
	"crypto/tls"
	"time"
)

func main() {

	// Functional Options编程模式
	NewServer("127.0.0.1", 3306, WithProtocol("https"), WithTimeout(10))

}

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

type Option func(*Server)

func WithProtocol(p string) Option {
	return func(server *Server) {
		server.Protocol = p
	}
}

func WithTimeout(t time.Duration) Option {
	return func(server *Server) {
		server.Timeout = t
	}
}

func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {
	srv := &Server{
		Addr: addr,
		Port: port,
	}
	for _, option := range options {
		option(srv)
	}
	return srv, nil
}
