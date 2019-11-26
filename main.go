package main

import (
	"fmt"
	"log"
	"rpc_server/conf"
	"rpc_server/rpc_server"
	"rpc_server/server/OnlineJudgeService"
)

func init() {
	var err error
	conf.PROJECT_PATH, err = rpc_server.GetProjectPath()
	if err != nil {
		log.Fatalln("get project path fail")
	}
	conf.G_LOG_LEVEL = conf.LOG_LEVEL_INFO
}

func ServerRun() {
	s, err:= rpc_server.NewServer()
	if err != nil {
		fmt.Println("server create fail. err ", err)
		return
	}

	// 加载服务
	OnlineJudgeService.LoadOnlineJudgeService(s)

	// 运行
	s.Run()
}

func main() {

	ServerRun()
}
