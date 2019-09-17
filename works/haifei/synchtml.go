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
	"strconv"
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
var (
	CommCell        string // 客户端名
	StartTime       string
	ApplicationSize string
	Subclient       string
)

var Log *logDiy.Log1
var c Config

type Config struct {
	HtmlfileReg string `json:"htmlfileReg"`
	HtmlBakDir  string `json:"htmlBakDir"`
	Logfile     string `json:"logfile"`
	LogMaxDays  int    `json:"logMaxDays"`
	Timeinter   int    `json:"timeinter"`
	OracleDsn   string `json:"oracledsn"`
}

func readconfig() {
	cfgstr := g.ParseConfig("/home/go/src/godev/works/haifei/synchtml.json")
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
		Log.Error("mkdir htmlbak Error DIR: %s, err:%s", c.HtmlBakDir, err)
	}
	_, fileName := filepath.Split(file)
	Log.Printf("move file old:%s,new:%s\n", file, filepath.Join(c.HtmlBakDir, string(os.PathSeparator), fileName))
	err = os.Rename(file, filepath.Join(c.HtmlBakDir, string(os.PathSeparator), fileName))
	if err != nil {
		Log.Error("osRename htmlbak Error DIR: %s, err:%s", c.HtmlBakDir, err)
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

func formatFields(index int, s string) (rstr string) {
	switch true {
	case strings.Contains(s, "Phase(Write End Time)"):
		rstr = strings.Replace(s, "(Write End Time)", "", -1)
		break
	case strings.Contains(s, " (Compression Rate)"):
		rstr = strings.Replace(s, " (Compression Rate)", "", -1)
		break
	case strings.Contains(s, "(Space Saving Percentage)"):
		rstr = strings.Replace(s, "(Space Saving Percentage)", "", -1)
		break
	case strings.Contains(s, "(Current)"):
		rstr = strings.Replace(s, "(Current)", "", -1)
		break
	case s == "Agent /Instance":
		rstr = "AgentInstance"
		break
	case s == "Backup Set /Subclient":
		rstr = "BackupSetSubclient"
		break

	default:
		rstr = strings.TrimSpace(s)
	}
	return
}

func formatValues(index int, s string) (rstr string) {
	//fmt.Println(index, s)
	rstr = s
	switch true {
	case strings.Contains(s, "N/A"):
		rstr = strings.Replace(s, "N/A", "", -1)
		break
	case index == 2: // subclient
		if strings.Contains(s, "/") {
			Subclient = strings.Split(s, "/")[1]
		} else {
			Subclient = s
		}
		break

	case index == 6: // starttime 赋值 全局变量，对应自定义列
		if strings.Contains(s, "(") {
			StartTime = strings.Split(s, "(")[0]
		} else {
			StartTime = s
		}
		StartTime = strings.TrimSpace(StartTime)
		//StartTime = fmt.Sprintf("to date(%s mm/dd/yyyy hh24:mi:ss)",StartTime)
		//StartTime = "2019-08-12 04:00:00"
		break
	case index == 8: // starttime 赋值 全局变量，对应自定义列
		as := ""
		if strings.Contains(s, "(") {
			as = strings.Split(s, "(")[0]
		} else {
			as = s
		}

		switch true {
		case strings.Contains(as, "TB"):
			as = strings.Split(as, " TB")[0]
			fl, err := strconv.ParseFloat(as, 64)
			if err != nil {
				Log.Printf("应用程序大小转换float失败 err:%s\n", err)
				break
			}

			//am:=strconv.FormatFloat(fl*1024,'E',-1,64)
			ApplicationSize = fmt.Sprintf("%.2f", fl*1024*1014)
		case strings.Contains(as, "GB"):
			as = strings.Split(as, " GB")[0]
			fl, err := strconv.ParseFloat(as, 64)
			if err != nil {
				Log.Printf("应用程序大小转换float失败 err:%s\n", err)
				break
			}

			//am:=strconv.FormatFloat(fl*1024,'E',-1,64)
			ApplicationSize = fmt.Sprintf("%.2f", fl*1024)

		case strings.Contains(as, "MB"):
			as = strings.Split(as, " MB")[0]
			fl, err := strconv.ParseFloat(as, 64)
			if err != nil {
				Log.Printf("应用程序大小转换float失败 err:%s\n", err)
				break
			}
			ApplicationSize = fmt.Sprintf("%.2f", fl)
		case strings.Contains(as, "KB"):
			as = strings.Split(as, " KB")[0]
			fl, err := strconv.ParseFloat(as, 64)
			if err != nil {
				Log.Printf("应用程序大小转换float失败 err:%s\n", err)
				break
			}
			ApplicationSize = fmt.Sprintf("%.2f", fl/1024)
		case strings.Contains(as, "Not Run because of another job running for the same subclient"):
				break
		default:

			application, e := strconv.Atoi(strings.TrimSpace(as))
			if e != nil {
				Log.Error("change type string:%s,err:%s", as, e)
			}
			if application == 0 {
				ApplicationSize = "0"
			}

		}

		break
	}
	return
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

	CommCells := dom.Find("body:contains(CommCell)")
	CommCells.Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			ss := strings.Split(strings.Split(selection.Text(), "CommCell:")[1], "--")[0]
			//fmt.Printf("CommCell :%+v\n", ss)
			CommCell = strings.TrimSpace(ss)
		}

	})

	FiledsHeader := make([]string, 0) // 列头，序号为KEY
	t_tbodyHeader := dom.Find("body > table:nth-child(13) > tbody > tr:nth-child(1) > td")

	//有两种格式的HTML
	if len(t_tbodyHeader.Nodes) == 0{
		t_tbodyHeader = dom.Find("body > table:nth-child(12) > tbody > tr:nth-child(1) > td")
	}

	t_tbodyHeader.Each(func(i int, s *goquery.Selection) {
		sa := s.Text()
		//switch sa {
		//case "Agent /Instance":
		//	sa = "AgentInstance"
		//case "Backup Set /Subclient":
		//	sa = "BackupSetSubclient"
		//default:
		//	sa = strings.TrimSpace(sa)
		//}
		sa = formatFields(i, sa)
		FiledsHeader = append(FiledsHeader, sa)
	})

	//fmt.Println(FiledsHeader)
	t_tbodyData := dom.Find("body > table:nth-child(13) > tbody > tr")
	//fmt.Println(t_tbody.Html())

	FiledsHeader = append(FiledsHeader, []string{"START TIME", "COMMCELL", "APPLICATIONSIZE", "subclient"}...)
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
			//switch true {
			//case strings.Contains(ss1,"N/A"):
			//	ss1=strings.Replace(ss1,"N/A","",-1)
			//}
			ss1 = formatValues(i, ss1)
			tmpSubData = append(tmpSubData, ss1)
			//fmt.Println(i,ss1)
		})
		tmpSubData = append(tmpSubData, []string{StartTime, CommCell, ApplicationSize, Subclient}...)
		if len(tmpSubData) == headlen { // 数据列数要与 列头一样长
			Data = append(Data, tmpSubData)
		}

	})
	//fmt.Println(Data)
	f.Close()
	// todo 列头加入三列

	genSql(Data, FiledsHeader, htmlfile) // 这里应该包括数据库插入 ，可以是后台 go 后面

}

func sliceToString(sl []string) (returnstring string) {
	sl1 := sl[0:21]
	for i := 0; i < len(sl1); i++ {
		if i == len(sl1)-1 {
			returnstring = returnstring
			returnstring = returnstring + "\"" + sl1[i] + "\"" + ")"
		} else {
			returnstring = returnstring
			returnstring = returnstring + "\"" + sl1[i] + "\"" + ","
		}
	}
	return
}

func updatesliceToString(sl []string, value []string) (returnstring string) {
	wherestring := ""
	if (len(sl) == len(value)) {
		for i := 0; i < len(sl); i++ {
			if (value[i] == "") { // 空值路过
				continue
			}
			if i == len(sl)-1 {
				returnstring = returnstring + "\"" + sl[i] + "\"" + "='" + value[i] + "'"
				wherestring = wherestring + "\"" + sl[i] + "\"" + "='" + value[i] + "'"
			} else {
				returnstring = returnstring + "\"" + sl[i] + "\"" + "='" + value[i] + "',"
				wherestring = wherestring + "\"" + sl[i] + "\"" + "='" + value[i] + "' and "
			}

		}
	}
	returnstring = returnstring + " where " + wherestring
	return
}

func genSql(Data [][]string, header []string, htmlfile string) {

	baseSql := "insert into HF_BACKUPDETAIL(" + sliceToString(header) + " values("
	updateSql := "update HF_BACKUPDETAIL set "

	//fmt.Println(updateSql)
	//fmt.Println(baseSql)
	sqlArr := make([]map[string]string, 0)
	for _, value := range Data {
		tmpmap := make(map[string]string, 0)
		//fmt.Println(index, baseSql+sliceToString(value))
		valueStr := strings.Replace(sliceToString(value), "\"", "'", -1)
		//fmt.Println(valueStr)
		//sqlArr = append(sqlArr, baseSql+valueStr)
		tmpmap["insert"] = baseSql + valueStr
		tmpmap["update"] = updateSql + updatesliceToString(header, value)
		sqlArr = append(sqlArr, tmpmap)
	}

	stmtSql(sqlArr, htmlfile)
}

func stmtSql(sqlArr []map[string]string, htmlfile string) {
	os.Setenv("NLS_LANG", "")
	//if len(os.Args) != 2 {
	//	log.Fatalln(os.Args[0] + " user/password@host:port/sid")
	//}

	db, err := sql.Open("oci8", c.OracleDsn)
	//fmt.Printf("%+v\n",db)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	updatenum := 0
	insertnum := 0
	for _, sql := range sqlArr {
		//fmt.Println(sql["update"])
		//fmt.Println(sql["insert"])
		//r,err:=db.Exec(fmt.Sprintf("insert into BCD(\"Client\",\"AgentInstance\",\"BackupSetSubclient\") values('a','b',%d)", 1))
		Result, err := db.Exec(sql["update"])
		//fmt.Println(Result.LastInsertId())
		//fmt.Println(Result.RowsAffected())
		if err != nil {
			Log.Error("HtmlFile:%s ,Sql Exec Err:%s", htmlfile, err)
			continue
		}
		updatenum = updatenum + 1
		r, e := Result.RowsAffected()
		if err != nil {
			Log.Error("HtmlFile:%s ,Sql Exec Err:%s", htmlfile, e)
			continue
		}

		if r == 0 {
			_, err := db.Exec(sql["insert"])
			if err != nil {
				Log.Error("HtmlFile:%s ,Sql Exec Err:%s", htmlfile, err)
				continue
			}
			insertnum = insertnum + 1
		}

	}
	Log.Printf("htmlfile : %s ,total sql:%d,success update sql:%d,insert sql:%d\n", htmlfile, len(sqlArr), updatenum, insertnum)

	if err == nil{
		MoveFileChan <- htmlfile
	}

	//rows, err := db.Query("select 'Client' from BCD")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//defer rows.Close()
	//
	//for rows.Next() {
	//	var data string
	//	rows.Scan(&data)
	//	fmt.Println(data)
	//}
	//if err = rows.Err(); err != nil {
	//	log.Fatalln(err)
	//}

}
