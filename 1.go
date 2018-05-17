package main

import (
	"fmt"
	// "sync"
	// "io"
	// "os"
	// "reflect"
	// "runtime"
	// "strconv"
	// "time"
	"math"
	"unicode/utf8"
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
	// t := reflect.TypeOf(3)
	// fmt.Println(t, t.String(), t.Size(), t.Kind())
	// fmt.Printf("%T\n", t)
	// fmt.Printf("%T\n", reflect.ValueOf(3))
	// var (
	// 	w io.Writer = os.Stdout
	// )
	// ww := reflect.TypeOf(w)
	// fmt.Println(ww)
	// fmt.Printf("%T\n", os.Stdout)

	// s:=[]int{1,2,3,4,5,6,7,8,9}
	// for i := 1; i < 10; i++ {
	// 	for j := 1; j < i+1; j++ {
	// 		fmt.Printf("%d*%d=%d ", j, i, i*j)

	// 	}
	// 	fmt.Println("")

	// }

	// var c chan int = make(chan int)
	// func() {

	// 	c <- 2
	// }()
	// go func() {

	// 	for i := 0; i < 10; i++ {
	// 		// fmt.Printf("%d ", i)
	// 		fmt.Printf("%d ", <-c)
	// 	}
	// }()

	// time.Sleep(1 * time.Second)
	type text string
	// var t text = "1"
	// fmt.Println(t)
	// a := int(t)
	// a, _ := strconv.ParseInt(t, 10, 0)
	// fmt.Println(a)
	var t int = 65
	fmt.Printf("%s\n", string(t))
	fmt.Println(aaa, bbb, ccc, ddd, eee)

	a := "1111111111111"
	b :=
		`
------------
1111111111
1222222222222
23333333333`
	c := 'a'
	fmt.Println(a, b, c)
	fmt.Printf("%T,%T,%T\n", a, b, c)
	fmt.Println(math.MaxInt16, math.MaxFloat32)

	s := "Go编程"
	fmt.Println(len(s))

	fmt.Println(len([]byte{1, 2}))

	ss1 := []byte(s)
	ss2 := []rune(s)
	ss3 := utf8.RuneCount(ss1)
	ss4 := utf8.RuneCountInString(s)

	fmt.Println(ss1)
	fmt.Println(len(ss2))
	fmt.Println(ss3)
	fmt.Println(ss4)

}

const (
	aaa = 'A'
	bbb
	ccc = iota
	ddd
)

const (
	eee = iota
)
