package main

import (
	"fmt"
	"godev/mymodels/appinstall/utils"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"time"
	"github.com/satori/go.uuid"
)

//todo  下载agent.zip 到本地 /TmpDir/TmpFile 解压缩 到/TmpDir/TmpUnZipDir

type InstallStruct struct {
	TmpDir               string // /tmp
	TmpFile, TmpUnZipDir string // 都是以当前时间作为
	CurrTime             int64
	ZipFile              string // zip ftp上的文件， agent.zip
	version              string //agent version
	Uuid                 string
	DataJson             *utils.Install
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
	//fmt.Println(dataJson.InstallAll.RunDir)
	//fmt.Println(filepath.Join(t.TmpDir, t.TmpUnZipDir))

	t.DataJson = dataJson
	utils.CreateDir(dataJson.InstallAll.RunDir) //create rundir
	stderr, _ := utils.RunCommand(fmt.Sprintf("/usr/bin/cp -rf %s/* %s", filepath.Join(t.TmpDir, t.TmpUnZipDir), dataJson.InstallAll.RunDir))
	if stderr != "" {
		log.Printf("cp -rf fail")
	}
}

func (t *InstallStruct) CreateCfg() (err error) {
	u, err := uuid.NewV4()
	if err != nil {
		log.Printf("create uuid fail err:%s\n", err)

	}
	t.Uuid = u.String()
	homedir := t.DataJson.InstallAll.RunDir
	err = utils.CreateConf(1, filepath.Join(homedir, "cfg.json"), filepath.Join(homedir, "agent.json.tmp"), t.Uuid)
	if err != nil {
		log.Printf("create agent cfg.json fail err:%s\n", err)
	}
	return nil
}

func (t *InstallStruct) RunInstall() (utils.InstallResp) {
	//for _,step:=range []string{"installStep"}{
	//	resp,stop:=utils.StepRun(step, t.DataJson)
	//	if stop{
	//		fmt.Println(resp)
	//		break
	//	}
	//}
	var resp utils.InstallResp
	resp, _ = utils.StepRun("installStep", t.DataJson)

	return resp
}

func main() {
	var agentStruct InstallStruct
	agentStruct.TmpDir = "/tmp"
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

	agentStruct.CreateCfg()

	//run install.json  installstep

	resp := agentStruct.RunInstall()
	if !resp.Success {
		log.Printf("runinstall faile info:%+v\n", resp)
	}
	//agent  add  proclist monitor queue

	// sqlite version uuid

	//return uuid version appname missionid
}
