package main

import (
	"fmt"
	"godev/mymodels/appinstall/utils"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"time"
)

//todo  下载agent.zip 到本地 /TmpDir/TmpFile 解压缩 到/TmpDir/TmpUnZipDir

type InstallStruct struct {
	TmpDir               string // /tmp
	TmpFile, TmpUnZipDir string // 都是以当前时间作为
	CurrTime             int64
	ZipFile              string // zip ftp上的文件， agent.zip
	version              string //agent version
	uuid                 string
}

func (t *InstallStruct) ftpDown() {
	//ftp download
	var ftpdl utils.DownLoad
	ftpdl.Port = 21
	ftpdl.Host = "192.168.1.104"
	ftpdl.Pass = "ftp"
	ftpdl.User = "ftp"
	ftpdl.LocalPath = t.TmpDir // /tmp
	var ftpfile utils.DownLoadSub
	ftpfile.DirPath = "/pub"
	ftpfile.FilePath = "agent.zip"
	ftpdl.DownloadFile = []*utils.DownLoadSub{&ftpfile}

	utils.CreateDir(ftpdl.LocalPath)
	//fmt.Println(t.TmpFile)
	err := ftpdl.FtpDownLoads(t.TmpFile) // 下载文件 本地文件改为 时间戳的名字
	if err != nil {
		log.Printf("ftpdownload fail err:%s\n", err)
	}
	t.ZipFile = filepath.Join(ftpdl.LocalPath, ftpfile.FilePath) // 下载完成后 zip文件绝对路径
}

func (t *InstallStruct) unZip() {
	//unzip /tmp/unix/unix.zip
	//fmt.Println(t.ZipFile)
	err := utils.UnCompress(filepath.Join(t.TmpDir, t.TmpFile), filepath.Join(t.TmpDir, t.TmpUnZipDir)) // 1.压缩文件，2.解压缩路径(/tmp + 临时目录)
	if err != nil {
		log.Printf("unzip file %s,err:%s\n", t.ZipFile, err)
	}
}

func (t *InstallStruct) unInsJson() {
	f, err := ioutil.ReadFile(filepath.Join(t.TmpDir, t.TmpUnZipDir, "install.json"))
	if err != nil {
		log.Printf("读取install json文件:%s 失败 err:%s\n", filepath.Join(t.TmpDir, t.TmpUnZipDir, "install.json"), err)
	}
	err, dataJson := utils.UnInstallJson(f) // 解析Json

	if err != nil {
		log.Printf("解析install json文件:%s失败 err:%s\n", filepath.Join(t.TmpDir, t.TmpUnZipDir, "install.json"), err)
	}
	fmt.Println(dataJson.InstallAll.RunDir)
	fmt.Println(filepath.Join(t.TmpDir, t.TmpUnZipDir))
	utils.RunCommand(fmt.Sprintf("/usr/bin/cp -rf %s/* %s", filepath.Join(t.TmpDir, t.TmpUnZipDir), dataJson.InstallAll.RunDir))
}

func main() {
	var agentStruct InstallStruct
	agentStruct.TmpDir = "d:\\ssss"
	agentStruct.CurrTime = time.Now().Unix()
	agentStruct.TmpUnZipDir = strconv.FormatInt(agentStruct.CurrTime, 10) //时间戳的名字
	agentStruct.TmpFile = agentStruct.TmpUnZipDir + ".zip"                // 临时目录/unixtime/unixtime.zip

	//
	agentStruct.ftpDown()

	//
	agentStruct.unZip()

	//jiexi install.json,  rundir   cp -r /tmp/unixnano/ rundir

	agentStruct.unInsJson()
	// create cfg.json

	//run install.json  installstep

	//agent  add  proclist monitor queue

	// sqlite version uuid

	//return uuid version appname missionid
}
