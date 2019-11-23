package rpc_server

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"rpc_server/conf"
	"rpc_server/server_base"
)

type Serializable struct {}

func NewSerializable() *Serializable {
	return &Serializable{}
}

func (s *Serializable) GetSerializeMsg(conn *net.TCPConn) (serverbase.IMessage, error) {
	headData  := make([]byte, s.GetHeadLen())
	if _, err := io.ReadFull(conn, headData); err != nil {
		return nil, errors.New(fmt.Sprintln("read msg head error", err))
	}
	msg, err := s.deSerializationHead(headData)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("De serialize error ", err))
	}
	var data []byte
	if msg.GetDataLen() > 0 && msg.GetDataLen() < conf.MAX_MESSAGE_LEN {
		data = make([]byte, msg.GetDataLen())
		if _, err := io.ReadFull(conn, data); err != nil {
			return nil, errors.New(fmt.Sprintln("read msg body error ", err))
		}
	}
	msg.SetData(data)
	return msg, nil
}

func (s *Serializable) Serialize(msg serverbase.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetMsgType()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetData()); err != nil {
		return nil ,err
	}
	return dataBuff.Bytes(), nil
}

func (s *Serializable) deSerializationHead(binaryData []byte) (serverbase.IMessage, error) {
	dataBuff := bytes.NewReader(binaryData)
	msg := &Message{}
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.MsgType); err != nil {
		return nil, err
	}
	return msg, nil
}

func (s *Serializable) GetHeadLen() int32 {
	return 8
}