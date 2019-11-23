package rpc_server

import "fmt"

type Message struct {
	DataLen int32
	MsgType int32
	Data	[]byte
}

func NewMessage(dataLen int32, msgType int32, data []byte) *Message {
	return &Message{
		DataLen:dataLen,
		MsgType:msgType,
		Data:data,
	}
}

func(m *Message) GetDataLen() int32 {
	return m.DataLen
}

func(m *Message) GetMsgType() int32 {
	return m.MsgType
}

func(m *Message) GetData() []byte {
	return m.Data
}

func(m *Message) SetDataLen(dataLen int32) {
	m.DataLen = dataLen
}

func(m *Message) SetMsgType(dataType int32) {
	m.MsgType = dataType
}

func(m *Message) SetData(data []byte) {
	m.Data = data
}

func(m *Message) ShowData() {
	fmt.Printf("len = %d, type = %d ", m.GetDataLen(), m.GetMsgType())
	for _, v := range m.GetData() {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}