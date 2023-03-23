package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	ep := "/data/work/go/GoDevEach/程序自更新/main1"
	dir := filepath.Dir(ep)
	fullpath, err := filepath.Rel(dir, ep)
	if err != nil {
		log.Fatalf("filepath.Rel: %v\n", err)
	}
	logf, err := os.OpenFile("/tmp/22", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		panic(err)
	}
	//bashpath, err := exec.LookPath("nohup")
	//if err != nil {
	//	panic(err)
	//}
	cmd := exec.Command("./main1")
	cmd.Dir = dir
	cmd.Path = fullpath
	cmd.Env = os.Environ()
	cmd.Stdout = logf
	cmd.Stderr = logf
	//cmd.ExtraFiles = []*os.File{logf}

	//cmd.Args = []string{"> ", "/tmp/22 2>&1 ", "&"}

	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true} // xxx 拥有自己的进程组,子进程 独立于 父进程

	err = cmd.Run()
	if err != nil {
		log.Println("err:", err)
	}
	//cmd.Wait()
	time.Sleep(1 * time.Second)
	//fmt.Printf(" current process pid: %v\n", cmd.ProcessState.Pid())
}
