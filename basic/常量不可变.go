package main

import (
	"fmt"
	"unsafe"
)

/*
http://www.runoob.com/go/go-constants.html

*/

func init() { // 最先执行的 函数
	const a int = 1

	const (
		aa = 1
		bb //这里不定义，与上面 数据一样
	)

	const x, y, z = "xx", 1.1, false

	fmt.Printf("a: %d \n", a)
	fmt.Printf("aa: %d ,bb: %d \n", aa, bb)
	fmt.Printf("x: %s ,y: %6.2f ,z: %t \n", x, y, z)
}

func main() {
	//第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1
	const (
		a = iota //0
		b        //1 #这里与上面 延续上面的类型
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		m = 1 << iota
		n = 3 << iota
		o
		p
	)
	/*
		m=1：左移 0 位,不变仍为 1;
		n=3：左移 1 位,变为二进制 110, 即 6;
		o=3：左移 2 位,变为二进制 1100, 即 12;
		p=3：左移 2 位,变为二进制 11000,即 24。
	*/
	fmt.Println("m=", m)
	fmt.Println("n=", n)
	fmt.Println("o=", o)
	fmt.Println("p=", p)

	const (
		aa = "abc"
		bb = len(aa) //3 个字符

		cc = unsafe.Sizeof(aa)
		//一部分是指向字符串起始地址的指针，另一部分是字符串的长度，两部分各是8字节，所以一共16字节
		dd = unsafe.Sizeof(bb)
		//一部分是指向整形起始地址的指针，另一部分是整形的长度，两部分各是4字节，所以一共8字节
		ee = unsafe.Sizeof(int(178))
	)
	println(aa, bb, cc, dd, ee)
}
