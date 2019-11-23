package serverbase

type ITaskPool interface {

	AddTask(request IRequest)

	StartWorker()

	StartOneWorker(int, chan IRequest)

	DoHandler(req IRequest)

	DoResponse(response IResponse) error
}
