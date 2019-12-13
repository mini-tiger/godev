package main

import (
	"fmt"
	"log"
	"time"
	"unsafe"
	"sync/atomic"
	"sync"
	"math/rand"
)

var data []string

// get data atomically
func Data() string {
	p := (*string)(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&data)),
	))
	if p == nil {
		return ""
	} else {
		return *p
	}
}

// set data atomically
func SetData(d string) {
	atomic.StorePointer(
		(*unsafe.Pointer)(unsafe.Pointer(&data)),
		unsafe.Pointer(&d),
	)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(200)

	for range [100]struct{}{} {
		go func() {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)

			log.Println(Data()) //xxx 提取数据
			wg.Done()
		}()
	}

	for i := range [100]struct{}{} {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)
			s := fmt.Sprint("#", i)
			log.Println("====", s)

			SetData(s) // xxx 写入数据，在没有写完前 其它纯种提取不到数据
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("final data = %v\n ", data)
}