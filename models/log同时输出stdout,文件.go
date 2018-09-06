package main

import (
	"os"
	"fmt"
	"log"
	"sync"
	"time"
	"runtime"
	"os/signal"
	"syscall"
)

type Log_m struct {
	Log_file    string
	Lock        sync.Mutex
	Logfile_obj *os.File
}

func (l Log_m) Write(p []byte) (n int, err error) {
	os.Stdout.Write(p)
	//fmt.Printf("%c\n",p)

	n, e := l.Write_file(p)
	return n, e
}

//todo 如果不绑定指针，l变量在方法内 会被copy一份
func (l *Log_m) Write_file(p []byte) (n int, err error) {
	//p=append(p,'\n')
	l.Lock.Lock()
	n, e := l.Logfile_obj.Write(p)
	defer l.Lock.Unlock()
	return n, e
}

func (l *Log_m) Createlogfile() {
	f, err := os.OpenFile(l.Log_file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Errorf("err:%s\n", err)
	}
	l.Logfile_obj = f
}

var log1 *log.Logger
var logfile_instance *Log_m
func initlog() {
	logfile_instance = &Log_m{Log_file: "C:\\work\\go-dev\\src\\godev\\log.txt"}
	logfile_instance.Createlogfile()
	log1 = log.New(*logfile_instance, "", log.Ldate|log.Ltime|log.Lshortfile)
}
func sync_close()  {
	f:=*(logfile_instance).Logfile_obj
	f.Sync()
	}

func main() {
	//var i io.Writer
	initlog()
	//log1.Printf("%d ",11111)

	runtime.GOMAXPROCS(runtime.NumCPU())
	exitchan := make(chan struct{}, 0)

	time_start := time.Now()
	go func(i int) { //线程一写日志
		for a := 0; a <= i; a++ {
			log1.Printf("this is goroute1 No.%d write log", a)
		}
		exitchan <- struct{}{}
		fmt.Printf("线程一使用了:%s \n",time.Since(time_start))
	}(100)

	go func(i int) { //线程二写日志
		for a := 0; a <= i; a++ {
			log1.Printf("this is goroute2 No.%d write log", a)
		}
		exitchan <- struct{}{}
		fmt.Printf("线程二使用了:%s \n",time.Since(time_start))
	}(100)

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	go func() { // 按ctrl+c 退出
		<-sign
		sync_close() //内存中 写入文件
		fmt.Println("exit")
		fmt.Printf("结束 使用了:%s \n",time.Since(time_start))
		os.Exit(0)
	}()
	go func() {// 写入日志 执行完成退出
		<-exitchan //两个线程中的第一个
		if _, ok := <-exitchan; ok { //两个线程中的第二个
			//fmt.Println(k, ok)
			sync_close() //将内存中数据 写入文件
			fmt.Printf("结束 使用了:%s \n",time.Since(time_start))
			os.Exit(0)
		}
	}()


	select {}
}
