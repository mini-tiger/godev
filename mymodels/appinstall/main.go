package main

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"godev/mymodels/appinstall/utils"
	"io/ioutil"
	"log"
	"os"
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
	Uuid                 string
	DataJson             *utils.Install
	Version              string // agent recv_version
}

var agentStruct InstallStruct

func (t *InstallStruct) ftpDown() (err error) {
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
	err = ftpdl.FtpDownLoads(t.TmpFile) // 下载文件 本地文件改为 时间戳的名字
	if err != nil {
		log.Printf("ftpdownload fail err:%s\n", err)
		return
	}
	t.ZipFile = filepath.Join(ftpdl.LocalPath, ftpfile.FilePath) // 下载完成后 zip文件绝对路径
	return
}

func (t *InstallStruct) unZip() error {
	//unzip /tmp/unix/unix.zip
	//fmt.Println(t.ZipFile)
	err := utils.UnCompress(filepath.Join(t.TmpDir, t.TmpFile), filepath.Join(t.TmpDir, t.TmpUnZipDir)) // 1.压缩文件，2.解压缩路径(/tmp + 临时目录)
	if err != nil {
		log.Printf("unzip file %s,err:%s\n", t.ZipFile, err)
	}
	return err
}

func (t *InstallStruct) unInsJson() (err error) {
	f, err := ioutil.ReadFile(filepath.Join(t.TmpDir, t.TmpUnZipDir, "install.json"))
	if err != nil {
		log.Printf("读取install json文件:%s 失败 err:%s\n", filepath.Join(t.TmpDir, t.TmpUnZipDir, "install.json"), err)
	}
	err, dataJson := utils.UnInstallJson(f) // 解析Json

	if err != nil {
		log.Printf("解析install json文件:%s失败 err:%s\n", filepath.Join(t.TmpDir, t.TmpUnZipDir, "install.json"), err)
		return
	}
	//fmt.Println(dataJson.InstallAll.RunDir)
	//fmt.Println(filepath.Join(t.TmpDir, t.TmpUnZipDir))

	t.DataJson = dataJson
	utils.CreateDir(dataJson.InstallAll.RunDir) //create rundir

	if utils.IsFile(filepath.Join(dataJson.InstallAll.RunDir, "info.sqlite")) {
		os.Remove(filepath.Join(t.TmpDir, t.TmpUnZipDir, "info.sqlite"))
	}

	stderr, _ := utils.RunCommand(fmt.Sprintf("/usr/bin/cp -rf %s/* %s", filepath.Join(t.TmpDir, t.TmpUnZipDir), dataJson.InstallAll.RunDir))
	if stderr != "" {
		log.Printf("cp -rf fail")
		err = errors.New("cp -rf fail")
		return
	}
	return
}

func (t *InstallStruct) CreateCfg() (err error) {

	homedir := t.DataJson.InstallAll.RunDir
	err = utils.CreateConf(1, filepath.Join(homedir, "cfg.json"), filepath.Join(homedir, "agent.json.tmp"), t.Uuid)
	if err != nil {
		log.Printf("create agent cfg.json fail err:%s\n", err)
	}
	return nil
}

func (t *InstallStruct) RunInstall() utils.InstallResp {
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

func (t *InstallStruct) SqliteCheck() (err error) {

	sql1 := utils.CreateSelectSql("agent", "uuid,version", "1=1")

	sqlconn := utils.SqlConn // 其它模块时

	err, b := sqlconn.GetExist(sql1)
	if err != nil {
		log.Printf("run sql :%s faile err:%s\n", sql1, err)
	}
	var uuidstr, currver, timestr string
	//uuid = "11111111"
	//ver = "22222222222"

	timestr = strconv.FormatInt(time.Now().Unix(), 10)
	//fmt.Println(uuidstr,ver,timestr)

	if b {
		err = sqlconn.GetData(sql1, &uuidstr, &currver)
		if err != nil {
			log.Println("getdata err", err)
		}
		fmt.Println(uuidstr, currver)
		if currver == t.Version {
			log.Println("version same skip")
			return
		}

		t.Uuid = uuidstr

		err = InstallDetail()
		if err != nil {
			log.Printf("InstallDetail err:%s\n", err)
		}

		err = sqlconn.UpdateAgentVer(uuidstr, t.Version, timestr)
		if err != nil {
			log.Printf("update err:%s\n", err)
		}

	} else {

		u, err := uuid.NewV4()
		if err != nil {
			log.Printf("create uuid fail err:%s\n", err)

		}
		t.Uuid = u.String()

		err = InstallDetail()

		if err != nil {
			log.Printf("InstallDetail err:%s\n", err)
		}

		err = sqlconn.InsertAgent(t.Version, uuidstr, timestr)

		if err != nil {
			log.Printf("insert agent err:%s\n", err)
		}

	}
	return
}

func InstallDetail() (err error) {
	agentStruct.TmpDir = "/tmp"
	agentStruct.CurrTime = time.Now().Unix()
	agentStruct.TmpUnZipDir = strconv.FormatInt(agentStruct.CurrTime, 10) //时间戳的名字
	agentStruct.TmpFile = agentStruct.TmpUnZipDir + ".zip"                // 临时目录/unixtime/unixtime.zip

	//
	err = agentStruct.ftpDown()
	if err != nil {
		log.Println("ftp err")
	}

	//
	err = agentStruct.unZip()
	if err != nil {
		log.Println("unzip err")
	}
	//jiexi install.json,  rundir   cp -r /tmp/unixnano/ rundir

	err = agentStruct.unInsJson()
	if err != nil {
		log.Println("unInsJson err")
	}
	// create cfg.json

	err = agentStruct.CreateCfg()
	if err != nil {
		log.Println("createCfg err")
	}
	//run install.json  installstep

	resp := agentStruct.RunInstall()
	if !resp.Success {
		log.Printf("runinstall faile info:%+v\n", resp)
	}
	return
}

func main() {

	// 没记录 直接插入，  有记录 判断是否 版本一样，一样则跳过，不一样更新
	// ubs 发送过来 必须带有版本号

	// sqlite version uuid
	// initDb
	err := utils.NewConn("/home/go/src/godev/mymodels/appinstall/old.sqlite")
	if err != nil {
		log.Printf("sqlite conn fail err:%s\n", err)
	}

	agentStruct.Version = "v3"

	agentStruct.SqliteCheck()

	//agent  add  proclist monitor queue
	//
	//

	//return uuid version appname missionid
}
