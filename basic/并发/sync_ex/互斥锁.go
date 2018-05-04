package main

//!+
import "sync"
import (
	"fmt"
	"time"
)

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock() //Lock方法将rw锁定为写入状态，禁止其他线程读取或者写入。

	balance = balance + amount
	time.Sleep(200 * time.Millisecond)
	defer mu.Unlock()
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

		go func(amount int) { //由于上面 2秒后 解锁，这里等待
			//
			Deposit(amount)
			n.Done()
		}(i)

	}
	n.Wait()

	fmt.Println(Balance())
}
