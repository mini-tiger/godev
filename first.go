package main

import "fmt"

type DownLoad struct {
	//Type         string
	Host string
	Port int
	User string
	Pass string
	//DownloadFile []*DownLoadSub
	LocalPath string
}

func main() {
	var ftp *DownLoad

	//s:=DownLoad{}
	//s.Port=21
	ftp = &DownLoad{Port: 21}
	fmt.Println(ftp)
}
