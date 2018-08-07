package main

import (
	e "./ex"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec" //执行命令行
	"path"
	"path/filepath"
	"strings"
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
	fmt.Println(os.Args[0]) //执行文件
	fmt.Println(os.Executable())
	var file string = "1.txt"
	// fmt.Println(file)
	if f, t := e.Exist(file); t {
		fmt.Printf("remove %s \n", f)
		os.Remove(f)
	} else {
		fmt.Printf("create %s \n", f)
		os.Create(f)
		ff := path.Join("aaaaa", f)
		os.MkdirAll(ff, 0755) //递归建立
		fmt.Println(filepath.Abs(f))
		fmt.Println(filepath.Abs(ff))
	}

	cmd := exec.Command("ipconfig")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String()) //in all caps: "SOME INPUT"

	a, err := os.Open("1.txt")
	fmt.Println(*a, err)

	input := bufio.NewScanner(os.Stdin) //需要命令行 输入
	// fmt.Println(input)
	for input.Scan() {
		fmt.Println(input.Text())

	}
	if err := input.Err(); err != nil {
		log.Fatal(err)

	}
}
