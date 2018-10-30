package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
)

type A struct {
	B int
}


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
	aa:=1
	bb:=1
	a1:=A{}
	a1.B=aa
	a2:=A{}
	a2.B=bb
	//fmt.Println(*a1==*a2)
	bbbccc(&a1,&a2)
	string1:="1"
	s:=strings.Split(string1,".")
	int, err := strconv.Atoi(s[0])
	fmt.Println(int,err)
	fmt.Printf("%f,%T\n",math.Floor(float64(100)/5/5),math.Floor(float64(100)/5/5))
	fmt.Println(int64(math.Floor(float64(100)/5/5)))

	}


func bbbccc(a,b *A)  {
	fmt.Println(a==b)
}