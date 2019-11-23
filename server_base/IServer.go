package serverbase

type IServer interface {

	Start()

	Run()

	Stop()

	GetConnManager() IConnManager

	GetTaskPool() ITaskPool

	GetMsgManager() IMsgManager

	GetLogManager() ILogManager

	// 配置文件
	LoadStrConf(key string, defaultVal string) string

	LoadIntConf(key string, defaultVal int) int

	// 路由
	AddRouter(id int, router IRouter) (bool, error)

	DelRouter(id int) (bool, error)
}
