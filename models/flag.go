package main

import (
	"flag"
	"fmt"
	"os"
)

//http://blog.studygolang.com/2013/02/%E6%A0%87%E5%87%86%E5%BA%93-%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%8F%82%E6%95%B0%E8%A7%A3%E6%9E%90flag/

var (
	levelFlag = flag.Int("level", 0, "级别")
	bnFlag    int
	ss        string
)

func init() {
	// 方式一
	// 四个参数
	// 1. 传入一个类型指针
	// 后面三个与方式二 一样
	flag.IntVar(&bnFlag, "bn", 3, "份数")
	flag.StringVar(&ss, "sn", "abc", "测试")
}

// 方式二
// 定义 三个参数
// 1. 参数名
// 2. 默认值
// 3. 帮助提示

func main() {

	count := len(os.Args)
	fmt.Println("参数总个数:", count)

	// Parse函数读取所有的命令行参数，即os.Args[1:]，并传入FlagSet的Parse方法

	flag.Parse()

	flag.Set("s", "abc")                 // 在已注册后，设置 参数的值

	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("flag 参数:", flag.Args()) //[a b c]

	fmt.Println("参数详情:")
	for i := 0; i < count; i++ {
		fmt.Println(i, ":", os.Args[i])
	}
	fmt.Println("bnFlag=", bnFlag)
	fmt.Println("ss=", ss)
	fmt.Println("levelFlag=", *levelFlag)
}
