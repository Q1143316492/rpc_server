package OnlineJudgeService

import (
	"fmt"
	"rpc_server/rpc_server"
	"rpc_server/server_base"
	"rpc_server/serverimpl/OnlineJudgeService"
)

type CppJudgeLXCService struct {
	rpc_server.RouterBase
}

func (r *CppJudgeLXCService) Handler(req serverbase.IRequest, res serverbase.IResponse) {
	fmt.Println("CppJudgeLXCService Handler test")

	// req.GetMessage().ShowData()

	err := OnlineJudgeService.CreateSandBoxEnv(req.GetBelongServer());
	if err != nil {
		msg := fmt.Sprintln("create sandbox env fail. err :", err)
		_ = res.AddParam("msg", msg)
		return
	}
	//res.AddParam("test", "test")
	//res.GetMessage().ShowData()
}
