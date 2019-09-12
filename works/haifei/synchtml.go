package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"path/filepath"
	"time"
	"tjtools/utils"
	"tjtools/logDiy"
	"godev/works/haifei/g"
	"encoding/json"
	"log"
	"strings"
	"database/sql"
	_ "github.com/mattn/go-oci8"

)

//https://github.com/360EntSecGroup-Skylar/excelize
// 中文 https://xuri.me/excelize/zh-hans/
// https://godoc.org/github.com/360EntSecGroup-Skylar/excelize#File.GetSheetName

const (
	//ReadFile    = "10.155.2.4_yuebao_English.html"
	//htmlfileReg = "C:\\work\\go-dev\\src\\godev\\works\\haifei\\*.html"
	//htmlBakDir  = "C:\\work\\go-dev\\src\\godev\\works\\haifei\\bak\\"
	//timeinter   = 10
	timeLayout = "2006-01-02 15:04:05"
)

var MoveFileChan chan string = make(chan string, 0)
var startRunTime = time.Now().Unix()

var Log *logDiy.Log1
var c Config

type Config struct {
	HtmlfileReg string `json:"htmlfileReg"`
	HtmlBakDir  string `json:"htmlBakDir"`
	Logfile     string `json:"logfile"`
	LogMaxDays  int    `json:"logMaxDays"`
	Timeinter   int    `json:"timeinter"`
}

func readconfig() {
	cfgstr := g.ParseConfig("C:\\work\\go-dev\\src\\godev\\works\\haifei\\synchtml.json")
	err := json.Unmarshal([]byte(cfgstr), &c)
	if err != nil {
		log.Fatalln("parse config file fail:", err)
	}
}

func main() {

	readconfig()

	fmt.Printf("%+v", c)

	logDiy.InitLog1(c.Logfile, c.LogMaxDays)
	Log = logDiy.Logger()

	go waitMovefile()
	for {
		files, _ := filepath.Glob(c.HtmlfileReg)
		//fmt.Println(files) // contains a list of all files in the current directory
		d, h, m, s := utils.GetTime(time.Now().Unix() - startRunTime)
		Log.Printf("本次读取到 %d个文件,%v\n", len(files), files)
		Log.Printf("开始运行时间 %s,已经运行%d天%d小时%d分钟%d秒", time.Unix(startRunTime, 0).Format(timeLayout), d, h, m, s)
		if len(files) >= 1 {
			waitHtmlFile(files)
		}
		time.Sleep(time.Duration(c.Timeinter) * time.Second)
	}
}

func MoveFile(file string) {
	err := os.MkdirAll(c.HtmlBakDir, 0777)
	if err != nil {
		Log.Printf("mkdir htmlbak Error DIR: %s, err:%s", c.HtmlBakDir, err)
	}
	_, fileName := filepath.Split(file)
	Log.Printf("move file old:%s,new:%s\n", file, filepath.Join(c.HtmlBakDir, "\\", fileName))
	err = os.Rename(file, filepath.Join(c.HtmlBakDir, "\\", fileName))
	if err != nil {
		Log.Printf("osRename htmlbak Error DIR: %s, err:%s", c.HtmlBakDir, err)
	}
}

func waitMovefile() {
	for {
		select {
		case file, ok := <-MoveFileChan:
			if ok {
				go MoveFile(file)
			} else {
				break
			}
		}
	}
}

func waitHtmlFile(files []string) {
	for _, file := range files {
		Log.Printf("covert file:%s\n", file)
		ReadHtml(file)
		Log.Printf("finish covert file:%s\n", file)
	}

}

func ReadHtml(htmlfile string) {

	f, err := os.OpenFile(htmlfile, os.O_RDONLY, 0755)
	if err != nil {
		Log.Println(err)
	}

	dom, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		//log.Fatal(err)
		Log.Printf("[Error]:%s", err)
	}
	//fmt.Println(dom)

	FiledsHeader := make([]string, 0) // 列头，序号为KEY
	t_tbodyHeader := dom.Find("body > table:nth-child(13) > tbody > tr:nth-child(1) > td")
	//fmt.Println(t_tbody.Html())

	t_tbodyHeader.Each(func(i int, s *goquery.Selection) {
		sa := s.Text()
		switch sa {
		case "Agent /Instance":
			sa = "AgentInstance"
		case "Backup Set /Subclient":
			sa = "BackupSetSubclient"
		default:
			sa = strings.TrimSpace(sa)
		}

		FiledsHeader = append(FiledsHeader, sa)
	})

	fmt.Println(FiledsHeader)
	t_tbodyData := dom.Find("body > table:nth-child(13) > tbody > tr")
	//fmt.Println(t_tbody.Html())

	headlen := len(FiledsHeader)
	Data := make([][]string, 0)

	t_tbodyData.Each(func(i int, s *goquery.Selection) {
		if i == 0 { // 0行是列名
			return
		}
		tmpSubData := make([]string, 0)
		sa := s.Find("td")
		sa.Each(func(i int, selection *goquery.Selection) {
			ss1 := selection.Text()
			tmpSubData = append(tmpSubData, ss1)
			//fmt.Println(i,ss1)
		})
		if len(tmpSubData) == headlen { // 数据列数要与 列头一样长
			Data = append(Data, tmpSubData)
		}

	})
	//fmt.Println(Data)
	f.Close()
	genSql(Data, FiledsHeader) // 这里应该包括数据库插入 ，可以是后台 go 后面
	MoveFileChan <- htmlfile
}

func sliceToString(sl []string) (returnstring string) {
	sl1:=sl[0:3]
	for i := 0; i < len(sl1); i++ {
		if i == len(sl1)-1 {
			returnstring = returnstring + sl1[i] + ")"
		} else {
			returnstring = returnstring + sl1[i] + ","
		}
	}
	return
}

func genSql(Data [][]string, header []string) {

	baseSql := "insert into BCD(" + sliceToString(header) + " values("

	//fmt.Println(baseSql)
	for index, value := range Data {
		fmt.Println(index, baseSql+sliceToString(value))
	}

	stmtSql()
}

func stmtSql()  {
		os.Setenv("NLS_LANG", "")
		//if len(os.Args) != 2 {
		//	log.Fatalln(os.Args[0] + " user/password@host:port/sid")
		//}

		db, err := sql.Open("oci8", "test/test@192.168.43.22:1521/orcl")
		//fmt.Printf("%+v\n",db)
		if err != nil {
			log.Fatalln(err)
		}

		defer db.Close()

		rows, err := db.Query("select 'Client' from BCD")
		if err != nil {
			log.Fatalln(err)
		}

		defer rows.Close()

		for rows.Next() {
			var data string
			rows.Scan(&data)
			fmt.Println(data)
		}
		if err = rows.Err(); err != nil {
			log.Fatalln(err)
		}

}
