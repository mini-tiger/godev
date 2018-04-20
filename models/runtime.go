package main

import "runtime"
import "log"

func main() {
	test()
}

func test() {
	test2()
}

func test2() {
	runtime.GOMAXPROCS(2)

	pc, file, line, ok := runtime.Caller(2)
	log.Println(pc)
	log.Println(file) //当前文件名
	log.Println(line) //调用行号
	log.Println(ok)   //
	f := runtime.FuncForPC(pc)
	log.Println(f.Name()) //调用的函数

	pc, file, line, ok = runtime.Caller(0)
	log.Println(pc)
	log.Println(file)
	log.Println(line)
	log.Println(ok)
	f = runtime.FuncForPC(pc)
	log.Println(f.Name())

	pc, file, line, ok = runtime.Caller(1)
	log.Println(pc)
	log.Println(file)
	log.Println(line)
	log.Println(ok)
	f = runtime.FuncForPC(pc)
	log.Println(f.Name())
}
