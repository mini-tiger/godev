package main

func foo() *int { // 外部要引用， xxx 逃逸到heap上
	var x int
	return &x
}

func bar() int { // 建立一个内存地址，直接值复制，外部没有引用，不会逃逸
	x := new(int)
	*x = 1
	return *x
}

func foo1(i *int) { // 没有返回 不会逃逸
	*i = 1

}

func foo2(i *int) *int { // leaking param: i to result ~r1 level=0,  流式变量，输入 与 输出同一个变量，没有逃逸
	*i = 1
	return i

}

func main() {
	var ii int
	ii = 2
	foo1(&ii)

	var ii2 int
	ii2 = 2
	ii2 = *foo2(&ii2)
}
