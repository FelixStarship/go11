package znet

import (
	"fmt"
	"github.com/FelixStarship/go11/leo/leov2.0/ziface"
	"net"
)

type Connection struct {
	Conn         *net.TCPConn
	ConnID       uint32
	isClosed     bool
	ExitBuffChan chan bool
	Route        ziface.IRoute
}

func NewConnection(conn *net.TCPConn, connID uint32, route ziface.IRoute) ziface.IConnection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		ExitBuffChan: make(chan bool, 1),
		Route:        route,
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running ")
	defer fmt.Println(c.RemoteAddr().String(), "conn reader exit！ ")
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err:", err)
			c.ExitBuffChan <- true
			continue
		}

		go func(req ziface.IRequest) {
			c.Route.PreHandler(req)
			c.Route.Handler(req)
			c.Route.PostHandler(req)
		}(&Request{
			conn: c,
			data: buf,
		})
	}
}
func (c *Connection) Start() {
	go c.StartReader()
	for {
		select {
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	err := c.Conn.Close()
	if err != nil {
		return
	}
	c.ExitBuffChan <- true
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
