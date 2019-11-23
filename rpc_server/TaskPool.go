package rpc_server

import (
	"errors"
	"fmt"
	"math/rand"
	"rpc_server/conf"
	"rpc_server/server_base"
)

type TaskPool struct {
	TaskQueue 			[]chan serverbase.IRequest
	TaskQueueSize 		int
	TaskQueueBuffSize	int
	BelongServer		serverbase.IServer
}

func NewTaskPool(server *Server) *TaskPool {
	p := &TaskPool{
		BelongServer:server,
	}
	p.TaskQueueSize 	= p.BelongServer.LoadIntConf("TASK_QUEUE_SIZE",   conf.TASK_QUEUE_SIZE)
	p.TaskQueueBuffSize = p.BelongServer.LoadIntConf("TASK_QUEUE_BUFFER", conf.TASK_QUEUE_BUFFER)
	p.TaskQueue = make([]chan serverbase.IRequest, p.TaskQueueSize)
	return p
}

func (t *TaskPool) AddTask(request serverbase.IRequest) {
	id := rand.Int() % t.TaskQueueSize
	t.TaskQueue[id] <- request
}

func (t *TaskPool) StartOneWorker(id int, worker chan serverbase.IRequest) {
	fmt.Printf("start one worker id = %d start\n", id)
	for {
		select {
		case req := <-worker:
			t.DoHandler(req)
		}
	}
}

func (t *TaskPool) DoHandler(req serverbase.IRequest) {
	router, err := t.BelongServer.GetMsgManager().GetRouter(req.GetMessage().GetMsgType())
	if err != nil {
		fmt.Println("DoHandler fail. request msg type not found")
		return
	}
	res := NewResponse(nil, nil)

	router.PreHandler(req, res)
	router.Handler(req, res)
	router.AfterHandler(req, res)

	res.SetConn(req.GetConn())
	res.SetMessage(NewMessage(0, req.GetMessage().GetMsgType(), []byte{}))
	res.GetMessage().SetData(res.InitParam())
	res.GetMessage().SetDataLen(int32(len(res.GetMessage().GetData())))

	fmt.Printf("request:")
	req.GetMessage().ShowData()
	fmt.Printf("response:")
	res.GetMessage().ShowData()

	if err = t.DoResponse(res); err != nil {
		fmt.Println("response error", err)
	}
}

func (t *TaskPool) DoResponse(response serverbase.IResponse) error {
	if response.GetMessage() == nil || response.GetMessage().GetDataLen() == 0{
		return errors.New(fmt.Sprintln("response message empty"))
	}
	ser := NewSerializable()

	data, err := ser.Serialize(response.GetMessage())
	if err != nil {
		return errors.New(fmt.Sprintln("response serialize fail. err = ", err))
	}
	response.GetConn().SendByteToClientBuffer(data)
	return nil
}

func (t *TaskPool) StartWorker() {
	for i := 0; i < t.TaskQueueSize; i++ {
		t.TaskQueue[i] = make(chan serverbase.IRequest, t.TaskQueueBuffSize)
		go t.StartOneWorker(i, t.TaskQueue[i])
	}
}