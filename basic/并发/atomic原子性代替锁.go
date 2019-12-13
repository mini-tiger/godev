package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

/*
fixme 原子性操作由底层硬件支持，锁是由操作系统API实现，前者效率更高
https://www.kancloud.cn/digest/batu-go/153537
*/

var data *string
var m map[int]interface{}
var arr []int

// get data atomically
func Data() (string,*string) {
	p := (*string)(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&data)),
	))
	if p == nil {
		log.Printf("%v,%p\n",nil,p)
		return "",p
	} else {
		log.Printf("%v,%p\n",*p,p)
		return *p,p
	}

}

func mData()  {
	p := (*map[int]interface{})(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&m)),
	))
	if p == nil {
		log.Printf("%v,%p\n",nil,p)

	} else {
		log.Printf("%v,%p\n",*p,p)

	}

}

func aData()  {
	p := (*[]int)(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&arr)),
	))
	if p == nil {
		log.Printf("%v,%p\n",nil,p)

	} else {
		log.Printf("%v,%p\n",*p,p)
	}

}

// set data atomically
func SetData(d string, tm map[int]interface{},ta []int) {
	atomic.StorePointer(
		(*unsafe.Pointer)(unsafe.Pointer(&data)),
		unsafe.Pointer(&d),
	)
	atomic.StorePointer(
		(*unsafe.Pointer)(unsafe.Pointer(&m)),
		unsafe.Pointer(&tm),
	)
	atomic.StorePointer(
		(*unsafe.Pointer)(unsafe.Pointer(&arr)),
		unsafe.Pointer(&ta),
	)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(200)

	for range [100]struct{}{} {
		go func() {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)

			Data() //xxx 提取数据
			mData() //xxx 提取数据
			aData() //xxx 提取数据
			wg.Done()
		}()
	}

	for i := range [100]struct{}{} {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)
			s := fmt.Sprint("#", i)
			//log.Println("====", s)

			SetData(s, map[int]interface{}{i: i},[]int{i}) // xxx 写入数据，在没有写完前 其它线程提取不到数据
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("final data = ", *data)

	// xxx 对于整形的加减 原子操作,在同一块内存上修改
	var i32 int32
	fmt.Printf("i32 value:%d，ptr:%p\n", i32, &i32)
	atomic.AddInt32(&i32, 3)
	fmt.Printf("i32 value:%d，ptr:%p\n", i32, &i32)

	/*
		xxx CAS 比较并交换,
		1. 判断第一1上参数addr指向的被操作值与 第二个参数old的值是否相等，
		2. 相等的情况 下，将第三个参数赋值给第一个参数的地址， 并返回true
	*/
	var newValue int32 = 11
	fmt.Printf("i32 eq :%v\n", atomic.CompareAndSwapInt32(&i32, newValue, 11)) // 这里i32 是3  ，应该不相等
	fmt.Printf("i32 value:%d，ptr:%p\n", i32, &i32)
	newValue = 3
	fmt.Printf("i32 eq :%v\n", atomic.CompareAndSwapInt32(&i32, newValue, 11))
	fmt.Printf("i32 value:%d，ptr:%p\n", i32, &i32)

}
