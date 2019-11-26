package rpc_server

import (
	"errors"
	"rpc_server/server_base"
	"rpc_server/tools"
)

type Request struct {
	msg   			serverbase.IMessage
	Conn  			serverbase.IConnection
	Param 			map[string]string
	belongServer	serverbase.IServer
}

func (r *Request) SetBelongServer(server serverbase.IServer) {
	r.belongServer = server
}

func (r *Request) GetBelongServer() serverbase.IServer {
	return r.belongServer
}

func NewRequest(message serverbase.IMessage, Conn serverbase.IConnection) *Request {
	req := &Request{
		msg:message,
		Conn:Conn,
		Param:make(map[string]string),
	}
	req.InitParam(req.Param, req.msg.GetData())
	return req
}

func (r *Request) InitParam(param map[string]string, data []byte) {
	params:= tools.ByteSliceSplit(data, '&')
	for _, v := range params {
		kv := tools.ByteSliceSplit(v, '=')
		if len(kv) == 2 {
			param[string(kv[0])] = string(kv[1])
		}
	}
}

func (r *Request) GetParam(key string) (string, error) {
	if key, ok := r.Param[key]; ok {
		return key, nil
	}
	return "", errors.New("request Param key not fount")
}

func (r *Request) AddParam(key string, val string) error {
	// todo key, val特殊符号url序列化转义
	if _, ok := r.Param[key]; ok {
		return errors.New("request Param key exist")
	}
	r.Param[key] = val
	return nil
}

func (r *Request) GetMessage() serverbase.IMessage {
	return r.msg
}

func (r *Request) SetMessage(msg serverbase.IMessage) {
	r.msg = msg
}

func (r *Request) GetConn() serverbase.IConnection {
	return r.Conn
}

func (r *Request) SetConn(conn serverbase.IConnection) {
	r.Conn = conn
}