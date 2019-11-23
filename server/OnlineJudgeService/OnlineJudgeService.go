package OnlineJudgeService

import (
	"rpc_server/conf"
	"rpc_server/server_base"
)

func LoadOnlineJudgeService(server serverbase.IServer) {
	_, _ = server.AddRouter(conf.CppJudgeLXCService, &CppJudgeLXCService{})
}
