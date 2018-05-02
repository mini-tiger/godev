package main

import (
	"fmt"
	// "time"
	"io/ioutil"
	"log"
)

type name interface {
	Print()
}

type sll struct {
	i int
}

func (ss *sll) Print() {
	fmt.Println(ss.i * ss.i)
}

func put(sa chan<- []int) {
	s := []int{}
	//加入一个接口调用
	var n name
	var sn sll
	//
	for i := 0; i < 3; i++ {
		for x := 0; x < 100; x++ {
			//加入一个接口调用
			sn.i = x
			n = &sn
			n.Print()
			//
			s = append(s, x)

		}
		sa <- s
		s = []int{} //清空
	}
	close(sa)
}

func jisuan(sa <-chan []int, sa1 chan<- int) {
	for x := range sa { //遍历channel
		s := 0
		for _, v := range x { //遍历切片
			s += v
		}
		sa1 <- s
	}
	close(sa1)
}
func tt() {
	var (
		sa  = make(chan []int)
		sa1 = make(chan int)
	)

	go put(sa)
	go jisuan(sa, sa1)
	for x := range sa1 {
		fmt.Println(x)

	}

}

func main() {
	// tt()
	// c := time.NewTicker(3 * time.Second)
	// fmt.Printf("%d\n", 1)
	// fmt.Println(<-c.C)
	files, err := ioutil.ReadDir("c:/godev")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			_, b := ioutil.ReadDir(file.Name())
			fmt.Println()
		} else {
			fmt.Printf("%s\n", file.Name())
		}
	}
	// fmt.Println(a)
}
