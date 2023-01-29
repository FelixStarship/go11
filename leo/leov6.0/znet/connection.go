package znet

import (
	"fmt"
	"io"
	"net"

	"github.com/FelixStarship/go11/leo/leov6.0/ziface"
	"github.com/pkg/errors"
)

type Connection struct {
	Conn         *net.TCPConn
	ConnID       uint32
	isClosed     bool
	ExitBuffChan chan bool
	MsgHandler   ziface.IMsgHandle
	msgChan      chan []byte
	TcpServer    ziface.IServer
}

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) ziface.IConnection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		ExitBuffChan: make(chan bool, 1),
		MsgHandler:   msgHandler,
		msgChan:      make(chan []byte),
		TcpServer:    server,
	}
	c.TcpServer.GetConnMgr().Add(c)
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running ")
	defer fmt.Println(c.RemoteAddr().String(), "conn reader exit！ ")
	defer c.Stop()

	for {
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			c.ExitBuffChan <- true
			continue
		}

		msg, err := dp.UnPack(headData)
		if err != nil {
			c.ExitBuffChan <- true
			continue
		}

		data := make([]byte, msg.GetDataLen())
		if msg.GetDataLen() > 0 {
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)

		req := &Request{
			conn: c,
			msg:  msg,
		}

		//go c.MsgHandler.DoMsgHandler(req)
		// 启动工作池机制
		c.MsgHandler.SendMsgToTaskQueue(req)
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("Writer Goroutine is running ")
	defer fmt.Println(c.RemoteAddr().String(), "Conn Writer exit!")
	defer c.Stop()

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send Data error:,", err, " Conn Writer exit")
				return
			}
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start() ... ConnID =", c.GetConnID())
	// 读客户端数据
	go c.StartReader()
	// 写客户端数据
	go c.StartWriter()
	for {
		select {
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop ... ConnID=", c.GetConnID())
	if c.isClosed {
		return
	}
	c.isClosed = true
	err := c.Conn.Close()
	if err != nil {
		return
	}
	c.ExitBuffChan <- true

	c.TcpServer.GetConnMgr().Remove(c)
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
	//数据写回客户端
	c.msgChan <- msg
	return err
}
