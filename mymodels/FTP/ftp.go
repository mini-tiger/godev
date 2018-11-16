package main


import (
	"github.com/shenshouer/ftp4go"
	"path"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"os"
	"errors"
)

type DownLoadSub struct {
	DirPath  string `json:"dir_path"`
	FilePath string `json:"file_path"`
}

type DownLoad struct {
	Type         string
	Host         string
	Port         int
	User         string
	Pass         string
	DownloadFile []*DownLoadSub
	LocalPath    string
}

func createDir(dir string) (err error) {
	dirInfo, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err=os.MkdirAll(dir,os.ModePerm)
			return
		}
	}
	if dirInfo.IsDir() { //是目录
		return nil
	} else { //是文件
		return errors.New(fmt.Sprintf("dir:%s not dir", dir))
	}

	return nil

}


func main() {
	f, _ := ioutil.ReadFile("/home/go/src/godev/mymodels/FTP/download.json")
	err, dl := UnJson(f) // 解析Json
	//fmt.Println(dl.LocalPath)
	if err != nil {
		fmt.Println(1, err)
	}
	//fmt.Println(dl)
	createDir(dl.LocalPath)
	err = dl.FtpDownLoads() // 下载文件
	if err != nil {
		fmt.Println(2, err)
	}
}

func UnJson(jb []byte) (err error, dl *DownLoad) {
	err = json.Unmarshal(jb, &dl) //解析
	if err != nil {
		log.Println("解析json失败")
		return err, dl
	}
	return nil, dl
}

func (d *DownLoad) FtpDownLoads() (err error) {
	ftpClient := ftp4go.NewFTP(0) // 1 for debugging
	//connect
	_, err = ftpClient.Connect(d.Host, d.Port, "")
	if err != nil {
		return
	}
	defer ftpClient.Quit()
	_, err = ftpClient.Login(d.User, d.Pass, "")
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
	for _, v := range d.DownloadFile {
		Basedir := v.DirPath
		fileName := v.FilePath
		file := path.Join(Basedir, fileName)

		_, err = ftpClient.Size(file)
		if err != nil {
			log.Printf("ftpfile: %s No Exist", file)
			continue
		}

		err = ftpClient.DownloadResumeFile(file, filepath.Join(d.LocalPath, fileName), false)
		if err != nil {
			log.Printf("ftpfile: %s download fail err:%s", file, err)
		}
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
