package main

import (
	"fmt"
	"sync"
	"time"
)

type Myint int

func (m *Myint) Batch() {
	defer s.Done()
	time.Sleep(time.Duration(Myint(*m)) * time.Second)
	fmt.Printf("%v,%v\n", int(*m), time.Now().Format(time.RFC3339))
}

var c chan *Myint = make(chan *Myint, 0)
var s sync.WaitGroup

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			ii := Myint(i)
			c <- &ii
			time.Sleep(1 * time.Second)
		}
		close(c)
	}()

	for v := range c {
		s.Add(1)
		go v.Batch()

	}
	s.Wait()

	//fmt.Println(Abc())
	//time.Sleep(3*time.Second)
}

func Abc() (i int) {
	defer fmt.Println(i)
	fmt.Println(1)
	panic(i)
}
