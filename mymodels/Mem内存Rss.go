package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"os"
)

var ps *process.Process

func mem(n int)  {
	if ps == nil{
		p,err:=process.NewProcess(int32(os.Getpid()))
		if err!=nil{
			panic(err)
		}
		ps=p
	}
	mem,_:=ps.MemoryInfoEx()
	fmt.Printf("rss:%d,vss:%d\n",mem.RSS >> 20,mem.VMS>>20)

}
func main()  {

	mem(1)
	data:=new([10][1024*1024]byte)
	mem(2)
	for i := range data{
		for x,n:=0,len(data[i]);x<n;x++{
			data[i][x]=1
		}
		mem(3)
	}


}