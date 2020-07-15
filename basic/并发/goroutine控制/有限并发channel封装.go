package main

import (
	"fmt"
	"time"
)

type Semaphore struct {
	bufSize int
	channel chan int8
}

func NewSemaphore(concurrencyNum int) *Semaphore {
	return &Semaphore{channel: make(chan int8, concurrencyNum), bufSize: concurrencyNum}
}

func (this *Semaphore) TryAcquire() bool {
	select {
	case this.channel <- int8(0):
		return true
	default:
		return false
	}
}

func (this *Semaphore) Acquire() {
	this.channel <- int8(0)
}

func (this *Semaphore) Release() {
	<-this.channel
}

func (this *Semaphore) AvailablePermits() int {
	return this.bufSize - len(this.channel)
}

func main() {
	se := NewSemaphore(5) // 有限的并发
	for i := 0; i < 10000; i++ {
		fmt.Printf("循环第%d次\n", i)
		se.Acquire() // 没有可用这里会阻塞
		go func(index int) {
			defer se.Release()
			fmt.Printf("可用并发数%d,循环%d次\n", se.AvailablePermits(),index)
			time.Sleep(time.Duration(300000)*time.Millisecond)
		}(i)
		time.Sleep(time.Duration(500)*time.Millisecond)

	}
}
