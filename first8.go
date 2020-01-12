package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

func ClearMap() {
	tmpLock := make(chan struct{}, 0)
	_ = cap(tmpLock)
}

func main() {
	go FindHtml() // 生产者
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:7783", nil))
	}()
	select {}

}

var OnceHtmlFiles = sync.Pool{
	New: func() interface{} {
		return new(sync.WaitGroup)
	},
}

func FindHtml() { // xxx 这里必须要放在for 里，不能routine， 等待执行结束在重新扫描
	//go ddd()


	go RunStatus() // xxx 定时打印运行状态,GC
	select {}
}

type HtmlFiles struct {
	Files   []string
	Success bool
	Err     error
	d       *sync.WaitGroup
}

var hs *sync.WaitGroup = new(sync.WaitGroup)

func ddd() {
	for {
		hs.Add(1)
		h, _ := GetFileList("/home/go/GoDevEach/works/haifei/syncHtml/htmlData") // 遍历目录 包括子目录
		hs.Done()
		log.Println(fmt.Sprintf("%p,%p", h, hs))

		time.Sleep(time.Duration(1 * time.Second))
	}
}

var filesPool = sync.Pool{
	New: func() interface{} {
		return make([]string, 0)
	},
}

func GetFileList(path string) ([]string, error) {
	//files := make([]string, 0)
	files := filesPool.Get().([]string)
	f, err := os.Stat(path)
	if err != nil {
		return files, errors.New(fmt.Sprintf("path: %s ,Err:%s", path, err))
	}
	if !f.IsDir() {
		return files, errors.New(fmt.Sprintf("path: %s ,Not Dir", path))
	}

	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return files, errors.New(fmt.Sprintf("path: %s ,err:%s", path, err))
	}
	return files, nil
}

var startRunTime = time.Now().Unix()

type RunStat struct {
	mm                 runtime.MemStats
	d, h, m, s, sinter int64
	tm                 time.Time
}

var OnceRunStat = sync.Pool{
	New: func() interface{} {
		return &RunStat{}
	},
}

func RunStatus() {
	//gct := time.NewTicker(time.Duration(60) * time.Second)
	RunStatt := time.NewTicker(time.Duration(30) * time.Second)
	//go func() {
	//	for {
	//		time.Sleep(time.Duration(g.GetConfig().Timeinter) * 6 * time.Second)
	//		log.Printf("手动GC开始\n")
	//		runtime.GC()
	//		log.Printf("手动GC结束\n")
	//	}
	//}()
	//
	//for {
	//	time.Sleep(time.Duration(60) * time.Second)
	//	go RunStatusWheel()
	//	//debug.FreeOSMemory()
	//
	//}
	for {
		select {
		//case <-gct.C:
		//	log.Printf("手动GC开始\n")
		//	runtime.GC()
		//	log.Printf("手动GC结束\n")
		case <-RunStatt.C:
			go RunStatusWheel()

		}
	}

}
//var m = &RunStat{}
func RunStatusWheel() {

	//var mm runtime.MemStats
	//var d, h, m, s int64

	m := OnceRunStat.Get().(*RunStat)
	runtime.ReadMemStats(&(m.mm))
	log.Printf("Sys(从系统获取过的内存,虚拟内存):%+v kb\n", m.mm.Sys/1024)
	log.Printf("Alloc(golang语言框架堆空间分配的内存,go虚拟机分配的内存,同HeapAlloc):%+v kb\n", m.mm.Alloc/1024)
	//log.Printf("HeapAlloc(堆上目前分配的内存):%+v kb\n", mm.HeapAlloc/1024)
	log.Printf("HeapReleased(回收到OS的内存):%+v kb\n", m.mm.HeapReleased/1024)
	log.Printf("HeapInuse(堆上目前使用的内存):%+v kb\n", m.mm.HeapInuse/1024)
	log.Printf("HeapIdle(堆上目前未使用的内存):%+v kb\n", m.mm.HeapIdle/1024)
	log.Printf("HeapSys(系统获得的堆内存,包含缓存空间):%+v kb\n", m.mm.HeapSys/1024)
	log.Printf("StackInuse(栈正在使用的内存):%+v kb\n", m.mm.StackInuse/1024)
	log.Printf("StackSys(系统分配给运行栈的内存):%+v kb\n", m.mm.StackSys/1024)
	log.Printf("GCSys(GC垃圾回收标记元信息使用的内存):%+v kb\n", m.mm.GCSys/1024)
	log.Printf("NumGC:%+v \n", m.mm.NumGC)
	log.Printf("OtherSys(golang架构占用的额外内存):%+v kb\n", m.mm.OtherSys/1024)
	log.Printf("BySize:%+v kb\n", m.mm.BySize[60])
	//log.Printf("GC:%+v \n", mm.DebugGC)
	//log.Printf("Eneable GC:%+v \n", mm.EnableGC) 72144

	if m.mm.LastGC > 1000 {
		m.tm = time.Unix(int64(m.mm.LastGC/1000/1000/1000), 0)
		//log.Printf("Last GC:%+v \n", int64(mm.LastGC/1000/1000/1000))
		log.Printf("Last GC(GC最后一次执行的时间):%+v \n", m.tm.Format("2006-01-02 15:04:05"))
	}

	log.Printf("Next GC(保证HeapAlloc<=NextGc,计算下个周期的目标):%+v kb\n", m.mm.NextGC/1024)
	m.sinter = time.Now().Unix() - startRunTime
	m.d, m.h, m.m, m.s = GetTime(m.sinter)
	log.Printf("开始运行时间 %s,已经运行%d天%d小时%d分钟%d秒\n", time.Unix(startRunTime, 0).Format("2006-01-02 15:04:05"), m.d, m.h, m.m, m.s)

}

func GetTime(sinter int64) (d, h, m, s int64) {
	if day, interval := getBasic(sinter, 86400); day > 0 {
		d = day
		h, m = oneDay(interval)
	} else {
		h, m = oneDay(interval)
	}
	s = sinter - (d * 86400) - (h * 3600) - (m * 60)
	return
}

func oneDay(interval int64) (h, m int64) { // 一天内的小时分钟
	h = interval / 3600
	if h == 0 {
		m = interval / 60
		return
	} else {
		interval = interval - (3600 * h)
		m = interval / 60
		return
	}
	return
}

func getBasic(sinter, tt int64) (num, interval int64) { // 是否不足1天, 有几天, 减去天数后的时间差
	if num = sinter / tt; num > 0 {
		return num, sinter - (num * tt)
	} else {
		return 0, sinter
	}
	return
}
