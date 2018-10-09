package bussiness

import (
	log "github.com/ccpaging/nxlog4go"
	"io"
	"os"
	"path/filepath"
)

func Returnlog() *log.Logger {
	return Log
}

var Log *log.Logger

func init() {
	log.FileFlushDefault = 5 // 修改默认写入硬盘时间
	//log.LogCallerDepth = 3 //runtime.caller(3)  日志触发上报的层级
	rfw := log.NewRotateFileWriter(filepath.Join("mysqlsync.log")).SetDaily(true).SetMaxBackup(7)
	ww := io.MultiWriter(os.Stdout, rfw) //todo 同时输出到
	// Get a new logger instance
	// todo FINEST 级别最低
	Log = log.New(log.FINEST).SetOutput(ww).SetPattern("[%Y %T] [%L] (%s) %M\n")
}
