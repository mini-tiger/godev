package main

import (
	"fmt"
	"unsafe"
)

func main() {
	u := new(user)
	fmt.Println(*u)

	pName := (*string)(unsafe.Pointer(u))
	*pName = "张三"

	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge = 20

	fmt.Println(*u)

	f := (*func(int) string)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.f)))

	*f = func(a int) string {
		return "1"
	}

	fmt.Println(*u, u.f(1))
}

type user struct {
	name string
	age  int
	f    func(int) string
}
