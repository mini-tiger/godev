package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
)

func main() {

	//go run glog.go -logtostderr=true -alsologtostderr=false -log_backtrace_at=main.go:2
	//
	//-logtostderr=true 写入标准错误，flase是文件

	fmt.Println(flag.CommandLine)
	flag.Parse()

	fmt.Println(flag.NFlag())

	glog.Info("Testing glog.")
	glog.Flush()

	p, err := os.Getwd()
	if err != nil {
		glog.Info("Getwd: ", err)
	} else {
		glog.Info("Getwd: ", p)
	}

	glog.Info("Prepare to repel boarders")
	glog.Info("222222222222---log_backtrace_at")
	glog.Info("333333333333")

	glog.V(1).Infoln("Processed1", "nItems1", "elements1")

	glog.V(2).Infoln("Processed2", "nItems2", "elements2")

	glog.V(3).Infoln("Processed3", "nItems3", "elements3")

	glog.V(4).Infoln("Processed4", "nItems4", "elements4")

	glog.V(5).Infoln("Processed5", "nItems5", "elements5")

	glog.Error("errrrr")

	/*
		if glog.V(2) {
			glog.Info("Starting transaction...")
		}
		glog.V(2).Infoln("Processed", "nItems", "elements")
		ch := make(chan int)
		go func() {
			for i := 0; i < 100; i++ {
				glog.Info("info:", i)
			}
			ch <- 1
		}()
		<-ch
		glog.Fatalf("Initialization failed: %s", errors.New("test info"))
	*/

	exit()
}

func exit() {
	glog.Flush()
}
