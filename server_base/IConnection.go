package serverbase

type IConnection interface {

	Start()

	Stop()

	StartReader()

	StartWriter()

	GetConnID() int64

	SendByteToClientBuffer(data []byte)

	SendByteToClient(data []byte) error

	SendExitSign()
}
