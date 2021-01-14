package main

import (
	"fmt"
	"sync"
)

type PostResult struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
}

var PostResFree = sync.Pool{
	New: func() interface{} {
		return &PostResult{1, "2"}
	},
}

func main() {
	a := PostResFree.Get().(*PostResult)
	fmt.Println(a)
}
