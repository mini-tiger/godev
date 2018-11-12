package main

import (
	"fmt"
	"time"
)

func main()  {
	m:=make(map[string]string,0)
	if len(m) == 0{
		fmt.Println(1)
	}
	fmt.Println(len(m))
	m1:=[]map[string]string{}
	fmt.Println(len(m1))


	a:= func() {
		for   {
			fmt.Println(111222)
			time.Sleep(time.Duration(1)*time.Second)
		}

	}
	b:= func() {
		go a()
		for   {
			fmt.Println(222333)
			time.Sleep(time.Duration(1)*time.Second)
		}

	}
	go b()

	select {}


}
