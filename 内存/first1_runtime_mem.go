package main

import (
	"github.com/shirou/gopsutil/process"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

var RunStatFree = sync.Pool{
	New: func() interface{} {
		return &runtime.MemStats{}
	},
}

func mem() {
	//var ps *process.Process

	ps, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		panic(err)
	}
	//ps = p

	meminfo, _ := ps.MemoryInfoEx()
	log.Printf("内存监控 RSS占用:%d Kb, VSS占用:%d Kb \n", meminfo.RSS>>10, meminfo.VMS>>10)

}

var RunStatt = time.NewTicker(time.Duration(60) * time.Second)
var gg bool = false

func main() {

	for {
		select {
		//case <-gct.C:
		//	g.GetLog().Info("手动GC开始\n")
		//	runtime.GC()
		//	//debug.FreeOSMemory()
		//	g.GetLog().Info("手动GC结束\n")
		case <-RunStatt.C:
			mem()
			g()
		}
	}

}

func g() {
	if gg {
		var mm = RunStatFree.Get().(*runtime.MemStats)
		defer func() {
			RunStatFree.Put(mm)
		}()
		runtime.ReadMemStats(mm)
		log.Printf("Sys(从系统获取过的内存,虚拟内存):%+v kb\n", mm.Sys/1024)
		log.Printf("Alloc(golang语言框架堆空间分配的内存,go虚拟机分配的内存,同HeapAlloc):%+v kb\n", mm.Alloc/1024)
		log.Printf("HeapReleased(回收到OS的内存):%+v kb\n", mm.HeapReleased/1024)
		log.Printf("HeapInuse(堆上目前使用的内存):%+v kb\n", mm.HeapInuse/1024)
		log.Printf("HeapIdle(堆上目前未使用的内存):%+v kb\n", mm.HeapIdle/1024)
		log.Printf("HeapSys(系统获得的堆内存,包含缓存空间):%+v kb\n", mm.HeapSys/1024)
		log.Printf("StackInuse(栈正在使用的内存):%+v kb\n", mm.StackInuse/1024)
		log.Printf("StackSys(系统分配给运行栈的内存):%+v kb\n", mm.StackSys/1024)
		log.Printf("GCSys(GC垃圾回收标记元信息使用的内存):%+v kb\n", mm.GCSys/1024)
		log.Printf("NumGC:%+v \n", mm.NumGC)
		log.Printf("OtherSys(golang架构占用的额外内存):%+v kb\n", mm.OtherSys/1024)
		log.Printf("BySize:%+v kb\n", mm.BySize[60])
		log.Printf("MSpanSys:%+v kb\n", mm.MSpanSys/1024)
		log.Printf("MCacheInuse:%+v kb\n", mm.MCacheInuse/1024)
		log.Printf("Mallocs:%+v \n", mm.Mallocs)
		log.Printf("StackInuse:%+v kb\n", mm.StackInuse/1024)
		log.Printf("StackSys:%+v kb\n", mm.StackSys/1024)
	}
}
