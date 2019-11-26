package OnlineJudgeService

import (
	"fmt"
	"rpc_server/conf"
	"rpc_server/server_base"
)

func LoadOnlineJudgeService(server serverbase.IServer) {
	_, err := server.AddRouter(conf.CppJudgeLXCService, &CppJudgeLXCService{})
	if err != nil {
		fmt.Println("fail to load service CppJudgeLXCService")
	} else {
		fmt.Println("success to load service CppJudgeLXCService")
	}
}
