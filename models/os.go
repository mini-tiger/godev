package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec" //执行命令行
	"path"
	"time"
)

func exist(file string) bool {
	if d, e := os.Getwd(); e == nil {
		f := path.Join(d, file)
		// fmt.Println(f)

		if _, err := os.Stat(f); err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("文件: %s 不存在\n", f)
				return false
			}
		} else {
			fmt.Printf("文件: %s 存在\n", f)
			return true
		}
	}
	return false
}

/*
C:\godev\models>go run os.go
C:\Users\ADMINI~1\AppData\Local\Temp\go-build492234478\command-line-arguments\_obj\exe\os.exe
C:\Users\ADMINI~1\AppData\Local\Temp\go-build492234478\command-line-arguments\_obj\exe\os.exe <nil>
文件: C:\godev\models/1.txt 存在
remove 1.txt

j
j
k
k

*/

func main() {
	// a,b=os.getwd()
	//fmt.Println(os.Args[0]) //执行文件
	//fmt.Println(os.Executable())
	//var file string = "1.txt"
	//// fmt.Println(file)
	//if f, t := e.Exist(file); t {
	//	fmt.Printf("remove %s \n", f)
	//	os.Remove(f)
	//} else {
	//	fmt.Printf("create %s \n", f)
	//	os.Create(f)
	//	ff := path.Join("aaaaa", f)
	//	os.MkdirAll(ff, 0755) //递归建立
	//	fmt.Println(filepath.Abs(f))
	//	fmt.Println(filepath.Abs(ff))
	//}
	//
	//cmd := exec.Command("ipconfig")
	//cmd.Stdin = strings.NewReader("some input")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err := cmd.Run()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("in all caps: %q\n", out.String()) //in all caps: "SOME INPUT"
	//
	//a, err := os.Open("1.txt")
	//fmt.Println(*a, err)
	//
	//input := bufio.NewScanner(os.Stdin) //需要命令行 输入
	//// fmt.Println(input)
	//for input.Scan() {
	//	fmt.Println(input.Text())
	//
	//}
	//if err := input.Err(); err != nil {
	//	log.Fatal(err)
	//
	//}
	//
	//// todo 标准输入输出  一起
	//cmd = exec.Command("/bin/bash", "-c", "grep -i 'model name' /proc/cpuino |uniq|cut -d ':' -f 2")
	//buf, _ := cmd.CombinedOutput()
	//cmd.Run()
	////fmt.Println(string(buf))
	//cm:=strings.Trim(string(buf)," ") // todo 统一返回
	//cm=strings.Trim(cm,"\n\r")
	//fmt.Println(cm)
	//
	//// todo 标准输入输出  分开
	//cmd = exec.Command("/bin/bash", "-c", "grep -i 'model name' /proc/cpuino |uniq|cut -d ':' -f 2")
	//
	//var out1 bytes.Buffer
	//cmd.Stdout = &out1
	//var e bytes.Buffer
	//cmd.Stderr =&e
	////cmd.Start()
	////buf, _ := cmd.CombinedOutput()
	//cmd.Run()
	//fmt.Println(e.String())
	//fmt.Println(out.String())

	cmd1 := exec.Command("/bin/bash", "-c", "/root/tomcat_start.sh")

	var out2 bytes.Buffer
	cmd1.Stdout = &out2
	var ee bytes.Buffer
	cmd1.Stderr =&ee
	//cmd1.Start()
	//buf, _ := cmd.CombinedOutput()
	cmd1.Run()
	fmt.Println(ee.String())
	fmt.Println(out2.String())

	time.Sleep(time.Duration(2)*time.Second)

	cmd2 := exec.Command("/bin/bash", "-c", "/root/tomcat_stop.sh")

	var out3 bytes.Buffer
	cmd2.Stdout = &out3
	var ee1 bytes.Buffer
	cmd2.Stderr =&ee1
	//cmd1.Start()
	//buf, _ := cmd.CombinedOutput()
	cmd2.Run()
	fmt.Println(ee1.String())
	fmt.Println(out3.String())
}
