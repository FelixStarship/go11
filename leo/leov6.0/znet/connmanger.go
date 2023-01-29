package znet

import (
	"errors"
	"fmt"
	"github.com/FelixStarship/go11/leo/leov6.0/ziface"
	"sync"
)

type ConnManger struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func NewConnManger() *ConnManger {
	return &ConnManger{
		connections: make(map[uint32]ziface.IConnection),
	}
}

func (c *ConnManger) Add(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn

	fmt.Println("connection add to ConnManger successful:conn num=", c.Len())
}

func (c *ConnManger) Remove(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	delete(c.connections, conn.GetConnID())

	fmt.Println("connection Remove ConnID=", conn.GetConnID(), " successful:conn num=", c.Len())
}

func (c *ConnManger) Len() int {
	return len(c.connections)
}

func (c *ConnManger) Get(connID uint32) (ziface.IConnection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	if conn, ok := c.connections[connID]; ok {
		return conn, nil
	}
	return nil, errors.New("connection not found")
}

func (c *ConnManger) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	for connID, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connID)
	}

	fmt.Println("Clear ALL Connections successful:conn num=", c.Len())
}
