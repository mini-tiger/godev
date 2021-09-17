package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  first
 * @Version: 1.0.0
 * @Date: 2021/9/10 下午2:56
 */

type A struct {
	AA int
}
type B struct {
	BB string
}

func (a *A) Data() {
	fmt.Println(a.AA)
}

func (b *B) Data() {
	fmt.Println(b.BB)
}

type Process interface {
	Data()
}
type Pros struct {
	MP Process
}

func main() {
	var a1 Process = &A{AA: 1}
	var b1 Process = &B{BB: "1"}

	a1.Data()
	b1.Data()

}
