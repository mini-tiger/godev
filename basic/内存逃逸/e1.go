package main

type S struct{}

func main() {
	var x S
	ref(x)
}
func ref(z S) *S {
	return &z
}

//
//go都是值传递，ref函数copy了x的值，传给z，返回z的指针，
//然后在函数外被引用，说明z这个变量在函数內声明，可能会被函数外的其他程序访问。xxx 所以z逃逸了，分配在堆上
