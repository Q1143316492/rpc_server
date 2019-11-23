package serverbase

type IRequest interface {

	GetMessage() IMessage

	SetMessage(IMessage)

	GetConn() IConnection

	SetConn(IConnection)

	GetParam(string) (string, error)

	AddParam(string, string) error

	InitParam(param map[string]string, data []byte)
}