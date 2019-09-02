package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)

type Aa struct {
	A1   func(int) int
	args int
}

func abc(a int) int {
	return 1 + a
}

func bcd(a int) int {
	return 2 + a
}

var C1 chan *Aa = make(chan *Aa, 1)

func t() {
	for {
		select {
		case a := <-C1:
			fmt.Println(a.A1(a.args))

		}
	}

}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		//db.DB.Close()
		os.Exit(0)
	}()


	go t()
	for i := 1; i < 100; i++ {
		C1 <- &Aa{abc, i}
		C1 <- &Aa{bcd, i}
		time.Sleep(time.Duration(1 * time.Second))
	}


}
