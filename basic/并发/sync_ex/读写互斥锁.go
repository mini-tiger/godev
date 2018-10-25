package main

//!+
import "sync"
import (
	"fmt"
	"time"
)

var (
	mu      *sync.RWMutex = new(sync.RWMutex)// guards balance
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


//todo RWMutex是一个读写锁，该锁可以加多个读锁或者一个写锁，其经常用于读次数远远多于写次数的场景．
/*
func (rw *RWMutex) Lock()　　写锁，如果在添加写锁之前已经有其他的读锁和写锁，则lock就会阻塞直到该锁可用，为确保该锁最终可用，已阻塞的 Lock 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定
func (rw *RWMutex) Unlock()　写锁解锁，如果没有进行写锁定，则就会引起一个运行时错误

func (rw *RWMutex) RLock() 读锁，当有写锁时，无法加载读锁，当只有读锁或者没有锁时，可以加载读锁，读锁可以加载多个，所以适用于＂读多写少＂的场景

func (rw *RWMutex)RUnlock()　读锁解锁，RUnlock 撤销单次RLock 调用，它对于其它同时存在的读取器则没有效果。若 rw 并没有为读取而锁定，调用 RUnlock 就会引发一个运行时错误(注：这种说法在go1.3版本中是不对的，例如下面这个例子)。

读写锁的写锁只能锁定一次，解锁前不能多次锁定，读锁可以多次，但读解锁次数最多只能比读锁次数多一次，一般情况下我们不建议读解锁次数多余读锁次数
*/

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
