package rpc_server

import (
	"fmt"
	"net"
	"os"
	"rpc_server/conf"
	"rpc_server/server_base"
	"rpc_server/tools"
	"strconv"
	"time"
)

type Server struct {

	ServerName 	 string

	IPVersion	 string

	IP 		 	 string

	Port	 	 string

	msgManager   serverbase.IMsgManager

	taskPool     serverbase.ITaskPool

	connManager	 serverbase.IConnManager

	confManager  serverbase.IConfManager

	logManager 	 *tools.LogManager

}

func NewServer() (*Server, error) {
	s := &Server{}
	var err error
	s.confManager, err = NewConfManager()
	if err != nil {
		fmt.Println("new server fail. err = ", err)
		return nil, err
	}
	s.logManager  = tools.NewLogManager()
	s.msgManager  = NewMsgManager(s)
	s.connManager = NewConnManager(s)
	s.taskPool    = NewTaskPool(s)

	if s.LoadStrConf("LOG_LEVEL", "INFO") == "DEBUG" {
		fmt.Println("set log level debug")
		s.GetLogManager().SetLogLevel(conf.LOG_LEVEL_DEBUG)
	}
	s.ServerName = s.LoadStrConf("SERVER_VERSION", conf.SERVER_VERSION)
	s.IPVersion  = s.LoadStrConf("SERVER_IP_VERSION", conf.SERVER_IP_VERSION)
	s.IP		 = s.LoadStrConf("SERVER_LISTEN_IP", conf.SERVER_LISTEN_IP)
	s.Port		 = s.LoadStrConf("SERVER_LISTEN_PORT", conf.SERVER_LISTEN_PORT)
	return s, nil
}

func (s *Server) LoadStrConf(key string, defaultVal string) string {
	val, err := s.confManager.GetConf(key)
	if err != nil {
		fmt.Println("load conf key = ", key, "not found. use default")
		return defaultVal
	}
	return  val
}

func (s *Server) LoadIntConf(key string, defaultVal int) int {
	val, err := s.confManager.GetConf(key)
	if err != nil {
		fmt.Println("load conf key = ", key, "not found. use default")
		return defaultVal
	}
	var v int
	if v, err = strconv.Atoi(val); err != nil {
		fmt.Println("conf not int")
		return defaultVal
	}
	return v
}

func (s *Server) GetConfManager() serverbase.IConfManager {
	return s.confManager
}

func (s *Server) SetConfManager(manager serverbase.IConfManager) {
	s.confManager = manager
}

func (s *Server) GetMsgManager() serverbase.IMsgManager {
	return s.msgManager
}

func (s *Server) GetTaskPool() serverbase.ITaskPool {
	return s.taskPool
}

func (s *Server) GetConnManager() serverbase.IConnManager {
	return s.connManager
}

func (s *Server) GetLogManager() serverbase.ILogManager {
	return s.logManager
}

func (s *Server) Start() {
	fmt.Printf("server %s start... IP = %s, port = %s\n", s.ServerName, s.IP, s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%s", s.IP, s.Port))
		if err != nil {
			fmt.Println("init socket error:", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("listen ip:%s, port:%s fail err:%s\n", s.IP, s.Port, err)
			return
		}

		s.taskPool.StartWorker()

		var id int64 = 0
		maxConnLimit := s.LoadIntConf("MAX_CONNECT_LIMIT", conf.MAX_CONNECT_LIMIT)
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept err = ", err)
				continue
			}
			if s.connManager.GetConnNumber() > maxConnLimit {
				fmt.Println("to many connections")
				_ = conn.Close()
				continue
			}
			workerConn := NewConnection(s, conn, id)
			go workerConn.Start()
			id++

		}
	}()
}

func (s *Server) Run() {
	if os.Getuid() != 0 {
		fmt.Println("server need root to run.")
		os.Exit(0);
	}
	s.Start()
	for {
		time.Sleep(10 * time.Second)
	}
}

func (s *Server) Stop() {
	s.connManager.ClearConn()
	fmt.Println("server stop..")
}

func (s *Server) AddRouter(id int, router serverbase.IRouter) (bool, error) {
	ret, err := s.msgManager.AddRouter(int32(id), router)
	return ret, err
}

func (s *Server) DelRouter(id int) (bool, error) {
	ret, err := s.msgManager.DelRouter(int32(id))
	return ret, err
}