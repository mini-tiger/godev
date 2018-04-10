package bao_demo

import (
	. "fmt"
)

const (
	AA = 1
	BB = 2 //这里不定义，与上面 数据一样
)

func T1(xx int) string {

	const x, y, z = "xx", 1.1, false
	return x + Sprint(xx)
}
