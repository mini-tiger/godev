package main

import "fmt"
import "runtime"

type B struct {
	User uint64
}

type A struct {
	Cpu []*B
}

func main() {
	ps := &A{Cpu: make([]*B, runtime.NumCPU())}
	fmt.Println(*ps)

}
