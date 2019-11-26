package main

import (
	"fmt"
	"net"
	"rpc_server/rpc_server"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:7737")
	if err != nil{
		fmt.Println("client dial err:", err)
		return
	}
	// request
	ser := rpc_server.NewSerializable()
	str := "name=cwl&id=1"
	msg := rpc_server.NewMessage(int32(len(str)),1001,[]byte(str))
	data, err := ser.Serialize(msg)
	if err != nil {
		fmt.Println("Serialize msg fail")
		return
	}
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("connect write data err")
		return
	}

	// response
	msg1, err := ser.GetSerializeMsg(conn.(*net.TCPConn))
	if err != nil {
		fmt.Println("get msg from net error")
		return
	}
	msg1.ShowData()
}