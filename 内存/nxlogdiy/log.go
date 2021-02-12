package nxlogdiy

import (
	logDiy "gitee.com/taojun319/tjtools/logDiyNew"
	nxlog "github.com/ccpaging/nxlog4go"
	"sync"
)

var (
	lock  = new(sync.RWMutex)
	logge *nxlog.Logger
)

func InitLog() *nxlog.Logger {
	// 初始化 日志

	logge = logDiy.InitLog1("/home/go/src/godev/内存/logs", 3, true, "DEBUG", true)
	return logge

}

func GetLog() *nxlog.Logger {
	lock.RLock()
	defer lock.RUnlock()
	return logge
}
