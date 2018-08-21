package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"math/rand"
	"time"
)

var totle_run int =1

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	class string
}

func create_instance() *Human {
	return &Human{}
}

func (h *Human) create_humandata() *Human {
	h.age = rand.Intn(100)
	h.name = "Alex"
	return h
}

func (s *Student) String() string {
	return fmt.Sprintf("myname is %s,age: %d,class:%s\n", s.name, s.age, s.class)
}

var c chan int = make(chan int)

func main() {
	var ip = flag.BoolP("version", "v", false, "版本")
	flag.Parse()
	if *ip {
		fmt.Println("0.2.1")
		return
	}
	go crond1()
	go crond()
	select {}

}

func crond() {
	for {
		select {
		case t:=<-c:
			time.Sleep(3*time.Second)
			fmt.Printf("this is No. %d run,waited 3s",t)

			backend()
		}

		//time.Sleep(2*time.Second)
	}
}

func crond1() {
	d := time.Duration(2 * time.Second)
	tick := time.Tick(d)
	for {
		select {
		case out := <-tick:
			fmt.Printf("this is crond1 %s ,this is No. %d run\n", out,totle_run)
			c <- totle_run
			totle_run+=1
		}
	}
}

func backend() {
	h1 := create_instance()
	h1.create_humandata()
	var s1 *Student = &Student{*h1, "3年二班"}
	fmt.Println(s1.String())
}
