package serverbase

import "net"

type ISerializable interface {

	GetHeadLen() int32

	// 从网络IO中取得一个消息
	GetSerializeMsg(conn *net.TCPConn) (IMessage, error)

	// 把消息转换为二进制byte串
	Serialize(msg IMessage) ([]byte, error)
}
