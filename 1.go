package main

import (
	"fmt"
	// "sync"
	"io"
	"os"
	"reflect"
	// "runtime"
	// "time"
)

// var n sync.WaitGroup
func sss1(ss1 chan<- int) {
	// n.Add(1)
	for i := 0; i < 10000000; i++ {

		if i%2 == 0 && i%3 == 0 {
			ss1 <- i
		}
		// ss2 <- ss1
		// fmt.Println(<-ss2)

	}
	// n.Done()
}
func sss2(ss2 chan<- int) {
	// n.Add(1)
	for i := 0; i < 10000000; i++ {
		if i%2 == 0 && i%3 == 0 {
			// fmt.Println(i)
			ss2 <- i
		}
		// fmt.Println(<-ss1)
	}
	// n.Done()
}
func main() {
	// runtime.GOMAXPROCS(2)
	// fmt.Println(runtime.GOOS, runtime.GOARCH)
	// var s1 chan int = make(chan int)
	// var s2 chan int = make(chan int)
	// // s2 <- "s2"
	// go sss1(s1)
	// // go func(ss1 chan int) {
	// // 	for i := 0; i < 100; i++ {

	// // 		fmt.Println(<-ss1)
	// // 	}
	// // }(s1)
	// go func() {
	// 	for {

	// 		fmt.Println(<-s1, <-s2)
	// 	}
	// }()
	// go sss2(s2)

	// // n.Wait()
	// time.Sleep(5 * time.Second)
	// for a := 0; a < 10; a++ {
	// 	go fmt.Print(0)
	// 	time.Sleep(time.Second * 1)
	// 	fmt.Print(1)
	// }
	t := reflect.TypeOf(3)
	fmt.Println(t, t.String(), t.Size(), t.Kind())

	var (
		w io.Writer = os.Stdout
	)
	ww := reflect.TypeOf(w)
	fmt.Println(ww)
	fmt.Printf("%T\n", os.Stdout)

}
