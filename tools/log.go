package tools

import (
	"fmt"
	"log"
	"os"
	"rpc_server/conf"
)

type LogManager struct {
	logLevel 		int
	levelMapper		map[int]string
}

func NewLogManager() *LogManager {
	manager := &LogManager{
		logLevel:conf.G_LOG_LEVEL,
		levelMapper:make(map[int]string),
	}
	manager.levelMapper[conf.LOG_LEVEL_DEBUG] 	   = "DEBUG"
	manager.levelMapper[conf.LOG_LEVEL_INFO] 	   = "INFO"
	manager.levelMapper[conf.LOG_LEVEL_WARN] 	   = "WARN"
	manager.levelMapper[conf.LOG_LEVEL_ERROR] 	   = "ERROR"
	manager.levelMapper[conf.LOG_LEVEL_CORE_ERROR] = "CORE_ERROR"
	return manager
}

func (l *LogManager) SetLogLevel(level int)  {
	l.logLevel = level
	conf.G_LOG_LEVEL = level
}

func (l *LogManager)Println(logs string) {
	// logfile := "logs/test.log"
	file, err := os.OpenFile(conf.PROJECT_PATH + "logs/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	prefix := fmt.Sprintf("[%s]", l.levelMapper[l.logLevel])
	logger := log.New(file, prefix, log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(logs)
}