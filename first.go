package main

import (
	"io"
	"os"
	"path/filepath"

	log "github.com/ccpaging/nxlog4go"
	"time"
	"fmt"
	"os/signal"
	"syscall"
)

const (
	filename = "_rfw.log"
	backups = "_rfw.*"
)

//// Print what was logged to the file (yes, I know I'm skipping error checking)
//func PrintFile(fn string) {
//	fd, _ := os.Open(fn)
//	in := bufio.NewReader(fd)
//	fmt.Print("Messages logged to file were: (line numbers not included)\n")
//	for lineno := 1; ; lineno++ {
//		line, err := in.ReadString('\n')
//		if err == io.EOF {
//			break
//		}
//		fmt.Printf("%3d:\t%s", lineno, line)
//	}
//	fd.Close()
//}


// todo 输出格式 layout.go
// todo level nxlog4go.go


var Log *log.Logger

func _filemaxsqlit()  {
	// Can also specify manually via the following: (these are the defaults)
	// 一个最大5K，最多10个
	rfw := log.NewRotateFileWriter(filepath.Join("c:\\work\\go-dev\\",filename)).SetMaxSize(1024 * 5).SetMaxBackup(10)
	ww := io.MultiWriter(os.Stdout, rfw) //todo 同时输出到
	// Get a new logger instance
	// todo FINEST 级别最低
	Log = log.New(log.FINEST).SetOutput(ww).SetPattern("[%Y %T] [%L] (%s) %M\n")
	// Log some experimental messages
	for j := 0; j < 15; j++ {
		for i := 0; i < 400 / (j+1); i++ {
			Log.Finest("Everything is created now (notice that I will not be printing to the file)")
			Log.Info("%d. The time is now: %s", j, time.Now().Format("15:04:05 MST 2006/01/02"))
			Log.Critical("Time to close out!")
		}
	}
	rfw.Close()
}
func rm_filemaxsqlit_file()  {
	// contains a list of all files in the current directory
	files, _ := filepath.Glob(backups)
	fmt.Printf("%d files match %s\n", len(files), backups)
	for _, f := range files {
		fmt.Printf("Remove %s\n", f)
		os.Remove(f)
	}
}

func main() {
	// Enable internal logger
	log.GetLogLog().SetLevel(log.TRACE) //全局 级别
	// todo 默认输出到os.stderr

	//_filemaxsqlit() //按文件最大分割
	//rm_filemaxsqlit_file() //删除文件

	_days_split()

	c:=make(chan os.Signal,1)
	signal.Notify(c,syscall.SIGINT|syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()
	select{}
}


func _days_split()  {
	log.FileFlushDefault = 2 // 修改默认写入硬盘时间

	rfw := log.NewRotateFileWriter(filepath.Join("c:\\work\\go-dev\\",filename)).SetDaily(true).SetMaxBackup(7)
	ww := io.MultiWriter(os.Stdout, rfw) //todo 同时输出到
	// Get a new logger instance
	// todo FINEST 级别最低
	// todo %p prefix, %N 行号
	Log = log.New(log.FINEST).SetOutput(ww).SetPattern("%P [%Y %T] [%L] (%s LineNo:%N) %M\n")
	Log.SetPrefix("11111")
	Log.SetLevel(0)

	// Log some experimental messages
	for j := 0; j < 15; j++ {
		for i := 0; i < 400 / (j+1); i++ {
			Log.Finest("Everything is created now (notice that I will not be printing to the file)")
			Log.Info("%d. The time is now: %s", j, time.Now().Format("15:04:05 MST 2006/01/02"))
			Log.Critical("Time to close out!")
			time.Sleep(1*time.Second)
		}
	}
	rfw.Close()
}
