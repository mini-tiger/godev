package main

//!+
import "sync"
import (
	"fmt"
	"time"
)

var (
	mu      sync.RWMutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.RLock() //RLock方法将rw锁定为读取状态，禁止其他线程写入，但不禁止读取。

	balance = balance + amount
	time.Sleep(200 * time.Millisecond)
	defer mu.RUnlock()
}

func Balance() int {
	// mu.Lock()
	// time.Sleep(2000 * time.Millisecond)
	// b := balance
	// mu.Unlock()
	return balance
}

func main() {
	// Deposit [1..1000] concurrently.

	var n sync.WaitGroup
	for i := 1; i <= 100; i++ {
		n.Add(1)
		go func() {
			fmt.Println(Balance())
		}()

		go func(amount int) {
			//
			Deposit(amount)
			n.Done()
		}(i)

	}
	n.Wait()

	fmt.Println(Balance())
}
