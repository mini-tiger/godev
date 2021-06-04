package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	cmd := exec.Command("/bin/sh", "-c", "watch date > date.txt")
	// Go会将PGID设置成与PID相同的值
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	start := time.Now()
	time.AfterFunc(3*time.Second, func() { syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL) })
	err := cmd.Run()
	fmt.Printf("pid=%d duration=%s err=%s\n", cmd.Process.Pid, time.Since(start), err)
}
