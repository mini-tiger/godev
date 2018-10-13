package main

import (
	"fmt"
)

func main() {
	//cpumhz,err:=nux.CpuMHz()
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(cpumhz)
	//cpucore:=runtime.NumCPU()
	//fmt.Println(cpucore)
	//cmd := exec.Command("/bin/bash", "-c", "grep -i processor /proc/cpuinfo |wc -l")
	//buf, _ := cmd.CombinedOutput()
	//cmd.Run()
	//fmt.Println(string(buf)) // todo 统一返回
	m:=make(map[string]interface{},1)
	m["1"]=1
	fmt.Println(len(m))
	}
