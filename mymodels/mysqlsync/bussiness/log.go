package bussiness

import (
	log "github.com/ccpaging/nxlog4go"
	"io"
	"os"
)

func Returnlog() *log.Logger {
	return Log
}

var Log *log.Logger

func Initlog() {

	log.FileFlushDefault = 5 // 修改默认写入硬盘时间
	//log.LogCallerDepth = 3 //runtime.caller(3)  日志触发上报的层级
	rfw := log.NewRotateFileWriter(Config.Logfile).SetDaily(true).SetMaxBackup(Config.LogMaxDays)
	ww := io.MultiWriter(os.Stdout, rfw) //todo 同时输出到
	// Get a new logger instance
	// todo FINEST 级别最低
	var level log.Level
	switch Config.LogLevel {
	case "FINEST":
			level = log.FINEST
	case "debug":
			level = log.DEBUG
	case "info":
			level = log.INFO
	}
	Log = log.New(level).SetOutput(ww).SetPattern("[%Y %T] [%L] (%s) %M\n")
}
