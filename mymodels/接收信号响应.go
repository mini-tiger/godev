package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)



/*




*/


func main()  {
	sigs := make(chan os.Signal,1)
	done :=make(chan bool,1)

	signal.Notify(sigs,syscall.SIGUSR1,syscall.SIGTERM) // 第二个参数  接收的信号类型， 第三个参数  信号的动作
	go func() {
		sig:=<-sigs
		fmt.Println(11111)
		fmt.Println(sig)
		done<-true
	}()
	<-done
}