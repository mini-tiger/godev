package main

import "fmt"

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  cost ioia
 * @Version: 1.0.0
 * @Date: 2021/5/14 下午3:45
 */

const (
	a = iota
	_
	b
	c = "zz"
	d
	f = iota
)

const (
	aa = iota
	bb
)

func main() {
	fmt.Println(a, b, c, d, f) // 0 2 zz zz 5
	fmt.Println(aa, bb)        // 0,1  cost 关键字出现 iota清0
}

/*
iota初始值为0，所以a为0，_表示不赋值，但是iota是从上往下加1的，所以b是2，c是“zz”,d和上⾯⼀个
同值也是“zz”,f是iota,从上0开始数他是5

*/
