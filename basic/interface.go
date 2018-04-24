package main

import (
	"fmt"
)

type Phone interface {
	call(ex int)
}

type Phone_instance struct {
	name   string
	money  int
	pinpai string
}

func (n *Phone_instance) call(ex int) {
	fmt.Printf("I am %s,myphone is %s,how much?, %d + %d !\n", n.name, n.pinpai, n.money, ex)
}

func main() {

	var tjphone Phone

	tjphone = &Phone_instance{"tj", 600, "nokia"} //内存地址给n，
	tjphone.call(2)                               //先传到interface
	//_____________________________________________________________
	var ctphone Phone
	var pp Phone_instance

	pp.money = 6000
	pp.name = "ct"
	pp.pinpai = "iphone"
	ctphone = &pp
	ctphone.call(250) //先传到interface

	example1() //使用上面的call ,在加个touch
}

func example1() {
	var tjphone Phone_er                        //局部变量
	tjphone = &Phone_instance{"tj", 1600, "mi"} //内存地址给n， 使用之前结构体
	tjphone.call(2)                             //先到interface
	tjphone.touch()

}

type Phone_er interface {
	call(ex int)
	touch()
}

func (i Phone_instance) touch() {
	fmt.Println("I am iPhone, I can touch you!")
}
