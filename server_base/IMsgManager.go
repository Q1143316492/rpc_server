package serverbase

type IMsgManager interface {

	AddRouter(int32, IRouter) (bool, error)

	DelRouter(int32) (bool, error)

	GetRouter(int32) (IRouter, error)
}