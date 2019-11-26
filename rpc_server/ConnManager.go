package rpc_server

import (
	"errors"
	"fmt"
	"rpc_server/server_base"
	"sync"
)

type ConnManager struct {
	BelongServer	serverbase.IServer
	connections 	map[int64]serverbase.IConnection
	connLock    	sync.RWMutex
}

func NewConnManager(server serverbase.IServer) *ConnManager {
	return &ConnManager{
		connections:make(map[int64]serverbase.IConnection),
		BelongServer:server,
	}
}

func (c *ConnManager) AddConn(conn serverbase.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn
	fmt.Printf("delete connection connID = %d\n", conn.GetConnID())
}

func (c *ConnManager) DelConn(connId int64) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connections, connId)
	fmt.Printf("delete connection connID = %d\n", connId)
}

func (c *ConnManager) GetConn(connId int64) (serverbase.IConnection, error) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	if conn, ok := c.connections[connId]; ok {
		return conn, nil
	} else {
		return nil, errors.New(fmt.Sprintf("connID = %d not found", connId))
	}
}

func (c *ConnManager) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	for connID, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connID)
	}
}

func (c *ConnManager) GetConnNumber() int {
	return len(c.connections)
}