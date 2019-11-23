package serverbase

type IRouter interface {

	PreHandler(IRequest, IResponse)

	Handler(IRequest, IResponse)

	AfterHandler(IRequest, IResponse)
}
