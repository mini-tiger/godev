package main

import (
	"fmt"
	"os/exec"
	"runtime"

)

func main() {

	cpucore:=runtime.NumCPU()
	fmt.Println(cpucore)
	cmd := exec.Command("/bin/bash", "-c", "grep -i processor /proc/cpuinfo |wc -l")
	buf, _ := cmd.CombinedOutput()
	cmd.Run()
	fmt.Println(string(buf)) // todo 统一返回
}
