package rpc_server

import (
	"errors"
	"rpc_server/server_base"
)

type Response struct {
	msg 			serverbase.IMessage
	Conn			serverbase.IConnection
	Param 			map[string]string
	belongServer	serverbase.IServer
}

func (r *Response) SetBelongServer(server serverbase.IServer) {
	r.belongServer = server
}

func (r *Response) GetBelongServer() serverbase.IServer {
	return r.belongServer
}

func NewResponse(message serverbase.IMessage, conn serverbase.IConnection) *Response {
	return &Response{
		msg:message,
		Conn:conn,
		Param:make(map[string]string),
	}
}

func (r *Response) InitParam() []byte {
	var data []byte
	for k, v := range r.Param {
		data = append(data, byte('&'))
		data = append(data, []byte(k)...)
		data = append(data, byte('='))
		data = append(data, []byte(v)...)
	}
	if len(data) > 0 && data[0] == '&' {
		return data[1:]
	}
	return data
}

func (r *Response) GetParam(key string) (string, error) {
	if key, ok := r.Param[key]; ok {
		return key, nil
	}
	return "", errors.New("request Param key not fount")
}

func (r *Response) AddParam(key string, val string) error {
	if _, ok := r.Param[key]; ok {
		return errors.New("request Param key exist")
	}
	r.Param[key] = val
	return nil
}

func (r *Response) GetMessage() serverbase.IMessage {
	return r.msg
}

func (r *Response) SetMessage(msg serverbase.IMessage) {
	r.msg = msg
}

func (r *Response) GetConn() serverbase.IConnection {
	return r.Conn
}

func (r *Response) SetConn(conn serverbase.IConnection) {
	r.Conn = conn
}