package main

import (
	"fmt"
	"sync"
	"time"
)

type test_struct struct {
	lock *sync.Mutex
	Name string
	Num  int
}

type test_struct_bind struct {
	sync.Mutex
	Num int
}

func (this *test_struct_bind) Call() {
	for {
		this.Lock()
		defer this.Unlock()
		fmt.Printf("current loop_run_struct_bind %d \n", this.Num)
		this.Num++
		time.Sleep(1 * time.Second)

	}

}

func main() {
	m := new(sync.Mutex)
	go loop_run(m) //单独加锁
	//////////////////////////////////

	tt := test_struct{}
	tt.lock = new(sync.Mutex)

	go loop_struct(&tt) // struct中有 一项是sync.Mutex
	///////////////////////////////////////
	var ss test_struct_bind = test_struct_bind{} //方法 ss := test_struct_bind{}
	go ss.Call()

	select {} //无限等待
}

func loop_run(mm *sync.Mutex) {
	num := 0
	for {
		(*mm).Lock()
		fmt.Printf("Current loop_run %d \n", num)
		num++
		time.Sleep(1 * time.Second)
		(*mm).Unlock()
	}

}

func loop_struct(t *test_struct) {
	for {
		t.lock.Lock()
		fmt.Printf("current loop_run_struct %d \n", t.Num)
		t.Num++
		time.Sleep(1 * time.Second)
		t.lock.Unlock()

	}

}
