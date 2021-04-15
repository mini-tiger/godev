package main

import (
	"fmt"
	"sync"
)

var num int

//var mtx sync.Mutex
var wg sync.WaitGroup

type Num struct {
	sync.Mutex
	N int
}

func (n *Num) add() {
	n.Lock()
	defer n.Unlock()
	defer wg.Done()
	num += 1
	fmt.Println(num)
}

func main() {
	var num *Num = &Num{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		if i%2 == 1 {
			go num.add()
		} else {
			go num.add()
		}

	}
	wg.Wait()

}
