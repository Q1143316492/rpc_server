package serverbase

type IMessage interface {

	GetDataLen() int32

	GetMsgType() int32

	GetData() []byte

	SetDataLen(dataLen  int32)

	SetMsgType(dataType int32)

	SetData([]byte)

	ShowData()
}
