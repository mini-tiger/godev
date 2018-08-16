package study2

import (
	"fmt"
	"runtime"
)

func init() {
	_, filename, _, _ := runtime.Caller(1)
	fmt.Printf("this is filename: %s \n", filename)
}

const (
	Ip = "192.168.43.11"
	DB = "db:mysql:" + Ip
)
