package serverbase

type IResponse interface {

	GetMessage() IMessage

	SetMessage(IMessage)

	GetConn() IConnection

	SetConn(IConnection)

	GetParam(string) (string, error)

	AddParam(string, string) error

	InitParam() []byte
}
