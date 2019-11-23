package OnlineJudgeService

import (
	"fmt"
	"rpc_server/rpc_server"
	"rpc_server/server_base"
)

type CppJudgeLXCService struct {
	rpc_server.RouterBase
}

func (r *CppJudgeLXCService) Handler(req serverbase.IRequest, res serverbase.IResponse) {
	fmt.Println("CppJudgeLXCService Handler test")

	//_ = res.AddParam("key", "val")
}
