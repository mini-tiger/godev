package main

import (
	"fmt"
	"sync"
)

type Config struct {
	Port int
}

func GetConfig() *Config {
	return &Config{80}
}

var cc *Config
var clock sync.Mutex

func GetConfig1() *Config {
	clock.Lock() // 线程安全
	defer clock.Unlock()
	if cc == nil {
		cc = &Config{80}
	}
	return cc
}

var so = sync.Once{}

func GetConfig2() *Config {
	so.Do(func() {
		cc = &Config{80}
	})
	return cc
}

func main() {
	c1 := GetConfig()
	c2 := GetConfig()

	fmt.Println("mem eq:", c1 == c2) // 不相等

	// xxx 方法一
	c1 = GetConfig1()
	c2 = GetConfig1()

	fmt.Println("mem eq:", c1 == c2) // 相等

	cc = nil
	// xxx 方法二
	c1 = GetConfig2()
	c2 = GetConfig2()

	fmt.Println("mem eq:", c1 == c2) // 相等

}
