package serverbase

type IConnManager interface {

	AddConn(IConnection)

	DelConn(int64)

	GetConn(int64) (IConnection, error)

	ClearConn()

	GetConnNumber() int
}
