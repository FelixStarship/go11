package tcp_epoller

import (
	"golang.org/x/sys/unix"
	"net"
	"sync"
)

type epoll struct {
	fd          int
	connections map[int]net.Conn
	lock        *sync.RWMutex
}

func MKEpoll() (*epoll, error) {
	fd, err := unix.EpollCreate1(0)
}
