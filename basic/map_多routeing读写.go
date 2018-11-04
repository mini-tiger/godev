package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type MM struct {
	sync.RWMutex //这样不用 单独给锁变量名， 如果加上变量名应该使用指针
	M            map[string]struct{}
}

func (s *MM) set(key string) {
	s.Lock()
	defer s.Unlock()
	s.M[key] = struct{}{}
}

func (s *MM) clear() {
	s.Lock()
	defer s.Unlock()
	s.M = make(map[string]struct{})
}

var m MM = MM{M: make(map[string]struct{}, 0)}

/*
RWMutex 是单写多读锁，该锁可以加多个读锁或者一个写锁
读锁占用的情况下会阻止写，不会阻止读，多个 goroutine 可以同时获取读锁
写锁会阻止其他 goroutine（无论读和写）进来，整个锁由该 goroutine 独占
适用于读多写少的场景
*/

func main() {

	go func() {
		for {
			m.RLock()
			for k, _ := range m.M {
				fmt.Printf("key:%s\n", k)

			}
			m.RUnlock()
			time.Sleep(time.Duration(1) * time.Second)
		}

	}()
	go func() {
		for {
			m.RLock()
			fmt.Printf("key:%+v\n", m.M)
			m.RUnlock()
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()

	go func() {
		a := 0
		for {
			m.set(strconv.Itoa(a))
			a = a + 1
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()

	go func() {
		for {
			//m.Lock()
			//m.M =make(map[string]struct{})
			m.clear()
			//m.Unlock()
			time.Sleep(time.Duration(2) * time.Second)
		}
	}()
	select {}
}
