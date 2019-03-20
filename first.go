package main

import (
	"os"
	"time"
	"fmt"
	"path/filepath"
)

func main()  {
	//openfile1()
	f,_:=os.Getwd()
	a:=filepath.Join(f,"cfg.json")
	fmt.Println(a)
}

func openfile1() {

	f, err := os.OpenFile("c:\\abc.log",os.O_APPEND| os.O_RDWR|os.O_CREATE, 0755) //读写模式
	if err != nil {
		os.Exit(1)
	}


	begintime:=time.Now().String()
	f.WriteString(begintime)
	f.Write([]byte("\n"))
	for{
		str1 := "aaaaa,bbbbbbbb,cccccccccc\n"
		ws := []byte(str1)
		f.Write(ws) //追加写
		time.Sleep(time.Duration(1)*time.Second)
	}


	f.Close()

}