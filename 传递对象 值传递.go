package main

import "fmt"

func main() {
	var i int = 1
	a := &i
	// xxx a指针也是变量, 变量有存储的内容 和 内存地址
	fmt.Printf("内存地址 %#v,%p\n", &a, &a) // (**int)(0xc0000b8018),0xc0000b8018
	copyint(a)                          // 传递到函数后的地址被复制
	fmt.Println(i)
}

func copyint(p *int) {
	// xx p指针也是变量, eee a,p两个指针变量（内容中的内存地址） 是一样的

	fmt.Printf("内存地址 %#v,%p\n", &p, &p) // (**int)(0xc0000b8028),0xc0000b8028
	*p = 2                              //修改指针变量中 对应内存地址的值
	fmt.Println(*p)
}
