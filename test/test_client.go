package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:7736")
	if err != nil{
		fmt.Println("client dial err:", err)
		return
	}
	ser := cnet.NewSerializable()
	str := "name=cwl&id=1"
	msg := cnet.NewMessage(int32(len(str)),1001,[]byte(str))
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

	msg1, err := ser.GetSerializeMsg(conn.(*net.TCPConn))
	if err != nil {
		fmt.Println("get msg from net error")
		return
	}
	msg1.ShowData()
}