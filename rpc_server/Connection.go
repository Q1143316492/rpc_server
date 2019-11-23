package rpc_server

import (
	"errors"
	"fmt"
	"net"
	"rpc_server/server_base"
)

type Connection struct {

	BelongServer		serverbase.IServer

	Conn				*net.TCPConn

	ConnID				int64

	isClosed			bool

	ExitBuffChan 		chan bool

	msgChan				chan []byte
}

func NewConnection(server serverbase.IServer, conn *net.TCPConn, connID int64) *Connection {
	c := &Connection{
		Conn:conn,
		ConnID:connID,
		isClosed:false,
		BelongServer:server,
		ExitBuffChan:make(chan bool, 1),
		msgChan:make(chan []byte),
	}
	c.BelongServer.GetConnManager().AddConn(c)
	return c
}

func (c *Connection) Start() {
	go c.StartReader()
	go c.StartWriter()
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	fmt.Printf("connect id = %d close\n", c.ConnID)
	c.isClosed = true
	c.SendExitSign()
	close(c.msgChan)
	close(c.ExitBuffChan)
	c.BelongServer.GetConnManager().DelConn(c.GetConnID())
}

func (c *Connection) StartReader() {
	defer c.Stop()
	for {
		ser := NewSerializable()
		msg, err := ser.GetSerializeMsg(c.Conn)
		if err != nil {
			fmt.Println("get a serialize msg error:", err)
			break
		}
		request := NewRequest(msg, c)
		c.BelongServer.GetTaskPool().AddTask(request)
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("start writer")
	for {
		select {
		case data := <- c.msgChan:
			// todo 扩展更多消息行为
			fmt.Println("writer get data:", data)
			_ = c.SendByteToClient(data)
		case _ = <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) GetConnID() int64 {
	return c.ConnID
}

func (c *Connection) SendByteToClientBuffer(data []byte) {
	c.msgChan <- data
}

func (c *Connection) SendByteToClient(data []byte) error {
	if _, err := c.Conn.Write(data); err != nil {
		return errors.New(fmt.Sprintf("send data:%s to client error:%s", string(data), err))
	}
	return nil
}

func (c *Connection) SendExitSign() {
	c.ExitBuffChan <- true
}