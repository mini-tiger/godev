package main

import (
	"fmt"
)

type empty interface {
}

type Connecter interface {
	Connect() //包含一个 方法
}

type USB interface {
	Name() string
	Connecter //嵌入Connecter的方法
}

//定义手机
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string { //PhoneConnecter 绑定
	return pc.name
}

func (pc PhoneConnecter) Connect() { //PhoneConnecter绑定
	fmt.Println("Connect:", pc.name)
}

//定义另一个设备 电视机
type TVConnecter struct {
	name string
}

func (tv TVConnecter) Connect() { //TVConnecter绑定
	fmt.Println("Connect:", tv.name)
}

func (tv TVConnecter) Name() string { //PhoneConnecter 绑定
	return tv.name
}

func main() {
	var a PhoneConnecter = PhoneConnecter{name: "phone"}

	a.Connect() //Connect: phone
	Disconnect(a)
	Disconnect_phone(a) //Disconnect_phone: phone

	var aa Connecter
	aa = Connecter(a) // PhoneConnecter是Connecter的超集，转换为子集 类型,只拥有子集方法Connect(),没有Name()
	aa.Connect()

	b := TVConnecter{"TV"} //Connect: TV
	b.Connect()
	Disconnect(b) //Disconnect: TV

}

func Disconnect_phone(usb USB) { //USB 接口包含的方法，都可以调用，空接口包含所有方法

	fmt.Println("Disconnect_phone:", usb.Name())

}

func Disconnect(usb interface{} /*empty*/) { // 空接口，匹配所有类型,这里要求的是USB类型, 而USB是interface类型
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnect:", v.Name())
	case TVConnecter:
		fmt.Println("Disconnect:", v.Name())
	default:
		fmt.Println("Unknown device.")
	}
}
