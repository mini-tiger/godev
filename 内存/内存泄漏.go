package main

import (
	nsema "gitee.com/taojun319/tjtools/control"
	mathrand "math/rand"
	"strconv"
	"time"
)

var HtmlFilesChan chan *S = make(chan *S, 0)
var sema *nsema.Semaphore = nsema.NewSemaphore(2)

type S struct {
	Str string
	Int int
}

func main() {

	go Revice()
	go Push()

	time.Sleep(2 * time.Second)
}

func Push() {

	for i := 0; i < 10; i++ {
		rr := new(S)
		rr.createInt(i)
	}

}
func (s *S) createStr() {
	s.Str = strconv.Itoa(mathrand.Int())

}
func (s *S) createInt(i int) {
	s.Int = i

}

func Revice() {
	for {
		select {
		case sa := <-HtmlFilesChan:
			sema.Acquire()
			sa.createStr()
			sema.Release()
		}
	}
}
