package main

import (
	"github.com/shenshouer/ftp4go"
	"path"
	"fmt"
	"encoding/json"
	"errors"
	"io/ioutil"
)

type DownLoadSub struct {
	DirPath string `json:"dir_path"`
	FilePath string 	`json:"file_path"`
}

type DownLoad struct {
	Type string
	Host string
	Port int
	User, passwd string
	DownloadFile []*DownLoadSub
}


func main() {
	var pp DownLoad
	// pp := &Jtest{}                       // 关联 struct
	f,_:=ioutil.ReadFile("C:\\work\\go-dev\\src\\github.com\\open-falcon\\falcon-plus\\modules\\agent\\download.json")
	ee := json.Unmarshal(f, &pp) //解析
	fmt.Println(ee)
	fmt.Println(pp)
	fmt.Println(pp.DownloadFile[0])
	fmt.Println(pp.DownloadFile[1])
	err := ftpClient()
	fmt.Println("1",err)
}

func ftpClient() (err error) {
	ftpClient := ftp4go.NewFTP(0) // 1 for debugging
	//connect
	_, err = ftpClient.Connect("192.168.43.12", ftp4go.DefaultFtpPort, "")
	if err != nil {
		return
	}
	defer ftpClient.Quit()
	_, err = ftpClient.Login("ftp", "ftp", "")
	if err != nil {
		//fmt.Println("The login failed")
		//fmt.Println(err)
		return
	}
	//Print the current working directory
	//var cwd string
	//cwd, err = ftpClient.Pwd()
	//if err != nil {
	//	//fmt.Println("The Pwd command failed")
	//	return
	//}
	//fmt.Println("The current folder is", cwd)
	Basedir := "/pub/1"
	fileName := "putty.zip"
	file := path.Join(Basedir, fileName)

	_, err = ftpClient.Size(file)
	if err != nil {
		return errors.New(fmt.Sprintf("ftpfile: %s No Exist", file))
	}

	err=ftpClient.DownloadResumeFile(file,"c:\\1.zip",false)
	if err != nil {
		return
	}


	//dirFiles, err := ftpClient.Nlst(s) //  文件夹是否存在
	//
	//if err != nil {
	//	return errors.New(fmt.Sprintf("Basedir No exist %s\n", s))
	//}
	//for _,filepath := range dirFiles{
	//	if filepath == ss{
	//		fmt.Println("file exist")
	//	}
	//}

	return nil
}
