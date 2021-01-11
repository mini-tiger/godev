package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type Logstruct struct {
	Type string
	Str  string
}

var logchan chan Logstruct = make(chan Logstruct, 0)

func main() {
	//print(runtime.GOOS,"\n")
	//print(runtime.GOARCH)
	//print(runtime.NumGoroutine())
	fmt.Println(runtime.Caller(0))

	go func() {
		for {
			select {
			case l, ok := <-logchan:
				if !ok {
					os.Exit(0)
				}
				fmt.Printf("[%s] info:%s\n", l.Type, l.Str)
			}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			logchan <- Logstruct{Type: "debug", Str: strconv.Itoa(i)}
		}
		close(logchan)
	}()

	select {}

}
