package rpc_server

import (
	"errors"
	iface "rpc_server/server_base"
	"rpc_server/tools"
)

type Request struct {
	msg   iface.IMessage
	Conn  iface.IConnection
	Param map[string]string
}

func NewRequest(message iface.IMessage, Conn iface.IConnection) *Request {
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

func (r *Request) GetMessage() iface.IMessage {
	return r.msg
}

func (r *Request) SetMessage(msg iface.IMessage) {
	r.msg = msg
}

func (r *Request) GetConn() iface.IConnection {
	return r.Conn
}

func (r *Request) SetConn(conn iface.IConnection) {
	r.Conn = conn
}