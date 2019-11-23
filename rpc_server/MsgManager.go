package rpc_server

import (
	"errors"
	"fmt"
	"rpc_server/server_base"
)

type MsgManager struct {
	routers 		map[int32]serverbase.IRouter
	BelongServer	serverbase.IServer
}

func NewMsgManager(server *Server) *MsgManager {
	return &MsgManager{
		routers:make(map[int32]serverbase.IRouter),
		BelongServer:server,
	}
}

func (m *MsgManager) AddRouter(id int32, router serverbase.IRouter) (bool, error) {
	if _, ok := m.routers[id]; ok {
		fmt.Println("Router id = ", id, "is registered")
		return false, errors.New("AddRouter fail. id is registered")
	}
	m.routers[id] = router
	return true, nil
}

func (m *MsgManager) DelRouter(id int32) (bool, error) {
	delete(m.routers, id)
	return true, nil
}

func (m *MsgManager) GetRouter(id int32) (serverbase.IRouter, error) {
	if _, ok := m.routers[id]; ok == false {
		fmt.Println("Router id = ", id, "not found")
		return nil, errors.New(fmt.Sprintln("Router id = ", id, "not found"))
	}
	return m.routers[id], nil
}
