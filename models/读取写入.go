package main

import (
	"bufio"

	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const file = "D:\\work\\project-dev\\src\\godev\\models\\111"

func main() {

	flag.Set("alsologtostderr", "true") // 日志写入文件的同时，输出到stderr
	//flag.Set("log_dir", "./log")        // 日志文件保存目录
	flag.Set("v", "5") // 配置V输出的等级。
	flag.Parse()

	//openfile1() //使用OS, 也具有read方法 可以使用bufio
	openfile2() //ioutil,  ioutil  要使用bytes.NewReader,
	//openfile3()//bufio,  os.openfile, 查询字符串，字符
}

func check(e error) {
	if e != nil {
		glog.V(0).Infoln("error", e)
	}
}

func openfile3() {
	glog.V(2).Infof("level 1 %s", "openfile3")
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755) //读写模式,返回满足io.Reader 接口 有read方法
	check(err)
	bs := bufio.NewReader(f) //创建 Reader 结构体实例，可以使用 绑定在这个结构体的方法,包括read方法
	for {
		//rb, err := bs.ReadByte() //一次读一个byte
		//line,err:=bs.ReadBytes('\n')//ReadBytes读取直到第一次遇到''\n字节，返回一个包含已读取的数据和'\n'字节的切片
		line, err := bs.ReadString('\n') //注意 单引号是byte，双引号是字符串, 返回字符串，查到'\n', 返回前面的字符串
		if err != nil || err == io.EOF {
			glog.V(0).Infoln("error", err)
			break
		}
		fmt.Printf("%s", line) //line 中 带有换行
	}

}

func openfile2() {
	glog.V(2).Infof("level 1 %s", "openfile2")
	b, err := ioutil.ReadFile(file) //Readfile 不会将读取返回的EOF视为应报告的错误
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("类型是%T\n", b)
	for i, v := range b {
		fmt.Printf("第%d个元素，byes是%v,字符串内容是%c\n", i, v, v) // %c byte到字符串
	}

	bs := strings.NewReader("hellow")
	bb, err := ioutil.ReadAll(bs) // bs 需要符合一个接口，要有Read 方法
	fmt.Printf("%T,%v,%c", bb, bb, bb)

	str1 := "aaaaa,bbbbbbbb,cccccccccc\n"
	ws := []byte(str1)
	for _, _ = range []byte{1, 2, 3} {
		ws = append(ws, ws...)
	}
	err = ioutil.WriteFile(file, ws, 0644) //先清空在写入
	check(err)

	//bss :=bytes.NewReader(bb)
	//fmt.Println(bss)
}

func openfile1() {
	glog.V(2).Infof("level 1 %s", "openfile1")

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755) //读写模式
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//ff,_:=f.Stat()
	//fmt.Println(ff.Name())

	//bs:=bufio.NewReader(f)
	//for {
	//	line,err:=bs.ReadString('\n')
	//	if err !=nil||err==io.EOF{
	//		break
	//	}
	//	fmt.Println(line)
	//}
	rs := make([]byte, 20) //读20个字节
	f.Read(rs)
	fmt.Println(string(rs))
	for i, v := range rs {
		fmt.Printf("第%d个元素，byes是%v,字符串内容是%c\n", i, v, v) // %c byte到字符串
	}

	str1 := "aaaaa,bbbbbbbb,cccccccccc"
	ws := []byte(str1)
	wf, err := f.Write(ws) //追加写
	fmt.Println(wf, err)

	f.Close()

}
