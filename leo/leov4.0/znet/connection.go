package znet

import (
	"fmt"
	"net"

	"github.com/FelixStarship/go11/leo/leov4.0/ziface"
	"github.com/pkg/errors"
)

type Connection struct {
	Conn         *net.TCPConn
	ConnID       uint32
	isClosed     bool
	ExitBuffChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32) ziface.IConnection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		ExitBuffChan: make(chan bool, 1),
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
	c.isClosed = true
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

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send msg")
	}
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		return errors.Wrap(err, "Pack error msg")
	}
	if _, err := c.Conn.Write(msg); err != nil {
		c.ExitBuffChan <- true
		return errors.Wrap(err, "conn Write err")
	}
	return err
}
