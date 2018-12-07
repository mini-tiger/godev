package main

import (
	"time"
	"log"
)

var nt1 *time.Ticker
var nt2 *time.Ticker

func main() {
	nt1 = time.NewTicker(time.Duration(2) * time.Second)
	nt2 = time.NewTicker(time.Duration(2) * time.Second)
	f1 := func() {
		for {
			select {
			case <-nt1.C:
				log.Println("this is english")
			}
		}
	}

	f2 := func() {
		for {
			select {
			case <-nt2.C:
				log.Println("这是中文")
			}
		}
	}
	run([]func(){f1, f2})
	select {}
}

func run(f []func()) {
	for _, fn := range f {
		go fn()
	}
}
