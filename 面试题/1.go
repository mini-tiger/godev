package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//defer_call()  // defer panic 顺序
	//pase_student() //循环赋值指针
	//three() // 随机性和闭包，变量作用域
	//foure() // chan 无缓冲阻塞,select 随机 触发case
	//five() // defer  嵌套函数 执行顺序
	//six() // 满足接口的方法，golang的方法集仅仅影响接口实现和方法表达式转化，与通过实例或者指针调用方法无关
	//seven() // defer 在return 之前执行，变量作用域
	eight() // const  iota
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
	// todo 考点 打印  顺序  panic应该在 defer 后面打印
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu // todo &stu 会变成 最后一次循环的 指针
	}
	// 正确
	//for i:=0;i<len(stus);i++  {
	//	m[stus[i].Name] = &stus[i]
	//}

	for key, value := range m {
		fmt.Printf("%v,%v\n", key, value)
	}

}

func three() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	// 输出顺序不定， 第一个for 每次都会输出 9. 第二个For 输出0-9
}
func foure() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1 // // 如果上面定义chan无缓冲，这里会阻塞报错，必须要先有 从<-chan的线程
	string_chan <- "hello"
	select {
	case value := <-int_chan: // 随机 选择一个case执行
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func five() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	/*
	与下面写法有区别
	虽然 第一个defer后执行，但 calc("1", a, calc("10", a, b))，里面的calc有先执行
	注释中的写法，func中的 calc("1", a, calc("10", a, b))，会整体后执行
	todo  执行到defer后，  会准备好defer所需要的参数，由于第三个参数是  函数，所以先执行
	*/
	//defer func() {
	//	defer calc("1", a, calc("10", a, b))
	//}()

	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

type People interface {
	Speak(string) string
}

type Stduent1 struct{}

func (stu *Stduent1) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func six() {
	var peo People = &Stduent1{} //  这里必须是 内存地址&,如果func (stu Stduent1)不用指针 ，这里可以不用地址
	//receiver 都是 value receiver，执行代码可以看到无论是 pointer 还是 value 都可以正确执行。
	/*
	如果是按 pointer 调用，go 会自动进行转换，因为有了指针总是能得到指针指向的值是什么，
	如果是 value 调用，go 将无从得知 value 的原始值是什么，因为 value 是份拷贝。go 会把指针进行隐式转换得到 value，但反过来则不行。
	*/
	think := "bitch"
	fmt.Println(peo.Speak(think))
}



func seven() {

	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}
/* 在return 分为两个部分
1. return 返回 已经声明的变量， DeferFunc1(i int) (t int) 声明了返回值t
2. 在defer 后进先出
*/
func DeferFunc1(i int) (t int) { // t 作用域整个函数
	t = i
	defer func() { // 返回 t=1 后，在执行这里 t=1+3
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() { // t不是作用于整个函数
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {   // 在return 之后执行 ,t=2,在执行t=1+2
		t += i
	}()
	return 2
}

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

func eight()  {
	fmt.Println(x,y,z,k,p) // 0 1 zz zz 4
	// &x 错误，不能 取地址
}