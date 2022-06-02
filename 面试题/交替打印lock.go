package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//var mtx sync.Mutex
var wg sync.WaitGroup

type Num struct {
	sync.Mutex
	N uint32
}

func (n *Num) add(s string) {
	n.Lock()
	defer n.Unlock()
	defer wg.Done()
	atomic.AddUint32(&n.N, 1)
	fmt.Println(s, n.N)
}

func main() {
	var num *Num = &Num{}

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		if i%2 == 1 {
			go num.add("dan")
		} else {
			go num.add("shuang")
		}

	}
	wg.Wait()

}
