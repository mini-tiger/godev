package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-oci8"
	"godev/works/haifei/syncHtml/g"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"tjtools/logDiy"
	"tjtools/utils"
)

//https://github.com/360EntSecGroup-Skylar/excelize
// 中文 https://xuri.me/excelize/zh-hans/
// https://godoc.org/github.com/360EntSecGroup-Skylar/excelize#File.GetSheetName


/*
1.查看 是否包含中文
2.确认CV8 CV*
3.找到摘要和报告中下面的颜色规律
4.生成摘要表 sql
5.生成详细信息表sql

*/
const (
	//ReadFile    = "10.155.2.4_yuebao_English.html"
	//htmlfileReg = "C:\\work\\go-dev\\src\\godev\\works\\haifei\\*.html"
	//htmlBakDir  = "C:\\work\\go-dev\\src\\godev\\works\\haifei\\bak\\"
	//timeinter   = 10
	timeLayout = "2006-01-02 15:04:05"
	FieldsLen  = 21 // 列一共有几列,使用时 要 减一
)

var MoveFileChan chan string = make(chan string, 0)
var startRunTime = time.Now().Unix()
var (
	CommCell        string // 客户端名
	GenTime         string
	StartTime       string
	ApplicationSize string
	Subclient       string
	StartTimeArr    []string = make([]string, 0)
	StatusColors     map[string]string
)

var Log *logDiy.Log1
var c Config
// 摘要表
//var FieldsMap map[string]string=map[string]string{"客户端":"ReportClient","主机名":"HOSTNAME","总作业数":"TOTALJOB",
//													"已完成":"COMPLETED","完成但有错误":"COMPLETIONERROR","完成但有警告":"COMPLETIONWARN",
//													"已终止":"TERMINATION","不成功":}

// 摘要表
var SummaryFieldsMap map[int]string=map[int]string{0:"ReportClient",1:"HOSTNAME",2:"TOTALJOB",
	3:"COMPLETED",4:"COMPLETEDWITHERRORS",5:"COMPLETEDWITHWARNINGS", 6:"KILLED",7:"UNSUCCESSFUL",8:"RUNNING",9:"DELAYED",
	10:"NORUN",11:"NOSCHEDULE",12:"COMMITTED",13:"SIZEOFAPPLICATION",14:"DATAWRITTEN",15:"STARTTIME",16:"ENDTIME",17:"PROTECTEDOBJECTS",
	18:"FAILEDOBJECTS",19:"FAILEDFOLDERS",
	20:"COMMCELL",21:"REPORTTIME"} // 这两个字段从页面上开头获取,唯一联合字段，也是和详细表 关系的字段

// 详细数据表
var DetailFieldsMap map[int]string=map[int]string{0: "dataclient" ,1:"AgentInstance", 2:"BackupSetSubclient" , 3:"Job ID (CommCell)(Status)",
	4:"Type", 5:"Scan Type", 6:"Start Time(Write Start Time)", 7:"End Time or Current Phase" ,8:"Size of Application", 9:"Data Transferred",
	10:"Data Written", 11:"Data Size Change",12:"Transfer Time",13:"Throughput (GB/Hour)", 14:"Protected Objects", 15:"Failed Objects",
	16:"Failed Folders",
	17:"COMMCELL",18:"REPORTTIME", // 与摘要表一样
	19:"START TIME", 20:"DATASUBCLIENT", // 通过开始时间 和 子客户端，格式化出来的字段
	21:"REASONFORFAILURE",22:"SOLVETIME",23:"ENGINEER",24:"SOLVETYPE"} // 失败原因，解决时间，工程师，解决状态 这几个字段不用插入数据，
	//todo 有问题的行 解决状态默认是未解决

type Config struct {
	HtmlfileReg string `json:"htmlfileReg"`
	HtmlBakDir  string `json:"htmlBakDir"`
	Logfile     string `json:"logfile"`
	LogMaxDays  int    `json:"logMaxDays"`
	Timeinter   int    `json:"timeinter"`
	OracleDsn   string `json:"oracledsn"`
}

func readconfig() {
	cfgstr := g.ParseConfig("/home/go/src/godev/works/haifei/syncYWReport/syncYWReport.json")
	err := json.Unmarshal([]byte(cfgstr), &c)
	if err != nil {
		log.Fatalln("parse config file fail:", err)
	}
}

func main() {

	readconfig()

	// 初始化 日志
	logDiy.InitLog1(c.Logfile, c.LogMaxDays)
	Log = logDiy.Logger()

	// 打印配置
	Log.Printf("读取配置: %+v\n", c)

	go waitMovefile()
	for {
		loopStartTime := time.Now()
		files, _ := filepath.Glob(c.HtmlfileReg)
		//fmt.Println(files) // contains a list of all files in the current directory
		d, h, m, s := utils.GetTime(time.Now().Unix() - startRunTime)
		Log.Printf("本次读取到 %d个文件,%v\n", len(files), files)
		Log.Printf("开始运行时间 %s,已经运行%d天%d小时%d分钟%d秒", time.Unix(startRunTime, 0).Format(timeLayout), d, h, m, s)
		if len(files) >= 1 {
			waitHtmlFile(files)
		}
		Log.Printf("本次处理 %d个文件,共用时 %.4f 秒\n", len(files), time.Now().Sub(loopStartTime).Seconds())
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
		Log.Printf("begin covert file:%s\n", file)
		resultNum := ReadHtml(file)
		switch resultNum {
		case 1:
			Log.Error("htmlFile %s, ERR:%s", file, "col not enough")
		case 2:
			Log.Error("htmlFile %s, ERR:%s", file, "col not Chinese")
		case 3:
			Log.Error("htmlFile %s, ERR:%s", file, "sql Gen err,sqlarr len is 0")
		case 4:
			Log.Error("htmlFile %s, ERR:%s", file, "insert or update have err")
		default:
			//MoveFileChan <- file
			Log.Printf("finish covert file:%s\n", file)
		}

	}

}

func formatFields(index int, s string) (rstr string) {
	switch true {
	case index == 0 && strings.Contains(s, "Client"):
		rstr = "DATACLIENT"
		break
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

func formatTime(toBeCharge string) string {
	//toBeCharge := "09/04/2019 11:03:16"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "01/02/2006 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	//fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	//fmt.Println(sr)
	return strconv.FormatInt(sr, 10)
}
func formatValues(index int, s string) (rstr string) {
	//fmt.Println(index, s)
	if strings.Contains(s, "N/A") {
		rstr = strings.Replace(s, "N/A", "", -1)
		s = rstr
	} else {
		rstr = s
	}
	if strings.Contains(s, "%") {
		rstr = strings.Replace(s, "%", "%%", -1)
		s = rstr
	} else {
		rstr = s
	}

	switch true {
	//case strings.Contains(s, "N/A"):
	//	rstr = strings.Replace(s, "N/A", "", -1)
	//	s = rstr
	case index == 2: // subclient

		tmpstr := s
		switch true {
		case strings.Contains(tmpstr, "/"):
			tmpstr = strings.Split(s, "/")[1]
			//fallthrough
			if strings.Contains(tmpstr, "Logcommand line") {
				Subclient = "Logcommand line"
			} else {
				Subclient = tmpstr
			}
			break
		default:
			Subclient = s

		}
		//fmt.Println("333333333333",index, s, Subclient)
		break

	case index == 6: // starttime 赋值 全局变量，对应自定义列
		StartTimeValue := ""
		if strings.Contains(s, "(") {
			StartTimeValue = strings.Split(s, "(")[0]
		} else {
			StartTimeValue = s
		}

		StartTimeValue = strings.TrimSpace(StartTimeValue)
		//StartTimeValue = fmt.Sprint("to_date('" + StartTimeValue + "','mm/dd/yyyy hh24:mi:ss')")
		StartTimeValue = formatTime(StartTimeValue)
		StartTimeArr = append(StartTimeArr, StartTimeValue)
		StartTime = "%s"

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

func ReadHtml(htmlfile string) (resultNum int) {

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

	// check chinese
	cvSelection:=dom.Find("body:contains(备份作业摘要报告)")
	if (cvSelection.Size() < 1){
		return 2
	}




	CommCells := dom.Find("body:contains(CommCell)")
	CommCells.Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			//fmt.Printf("CommCell :%+v\n", selection.Text())
			ss := strings.Split(strings.Split(selection.Text(), "CommCell:")[1], "--")[0]
			//fmt.Printf("CommCell :%+v\n", ss)
			CommCell = strings.TrimSpace(ss)
		}

	})
	fmt.Println(CommCell)

	// 生成的时间
	//GenTimeSource := dom.Find("body:contains(备份作业摘要报告)")
	cvSelection.Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			//fmt.Println(htmlfile,selection.Text())
			//fmt.Println(strings.Contains(selection.Text(),"generated"))
			ss := strings.Split(strings.Split(selection.Text(), "备份作业摘要报告")[1], "上生成的报表版本")[0]
			//fmt.Println(strings.Split(selection.Text(), "备份作业摘要报告"))
			//fmt.Println(ss)
			//fmt.Printf("GenTime :%+v\n", ss)
			GenTime = strings.TrimSpace(ss)
		}
	})

	fmt.Println(GenTime)



	FiledsHeader := new([]string) // 列头，序号为KEY
	t_tbodyHeader := dom.Find("body > table:nth-child(13) > tbody > tr:nth-child(1) > td")

	cv8 := false
	//有两种格式的HTML ,代表是CV8
	if len(t_tbodyHeader.Nodes) == 0 {
		cv8 = true
		t_tbodyHeader = dom.Find("body > table:nth-child(12) > tbody > tr:nth-child(1) > td")
	}


	// 循环找出列头
	t_tbodyHeader.Each(func(i int, s *goquery.Selection) {
		sa := s.Text()
		sa = formatFields(i, sa) // todo 格式化 列头
		*FiledsHeader = append(*FiledsHeader, sa)
	})

	// 如果是CV8版本，手动修改列名, 删除 size of Backup cols 这列
	if cv8 {
		tmp := make([]string, 16)
		copy(tmp[:], (*FiledsHeader)[:])

		*FiledsHeader = tmp[0:9]
		//fmt.Println(*FiledsHeader)
		//*FiledsHeader = append(*FiledsHeader, []string{"Data Transferred", "Data Written"}...)
		//fmt.Println(*FiledsHeader)
		*FiledsHeader = append(*FiledsHeader, tmp[10:]...)
		//fmt.Println(*FiledsHeader)
		(*FiledsHeader)[3] = "Job ID (CommCell)(Status)"

		GenTime = strings.TrimSpace(strings.Split(GenTime,"CommCell")[0])
	}
	//fmt.Println("222222222",*FiledsHeader)
	//for i := 0; i < len(*FiledsHeader); i++ {
	//	fmt.Printf("%d, value:%s\n", i, (*FiledsHeader)[i])
	//}

	switch true {
	case len(*FiledsHeader) < FieldsLen-6: // col not enough,运维报告中不用判断
		//fmt.Println(FiledsHeader)
		//fmt.Println(len(*FiledsHeader))
		return 1
	case (*FiledsHeader)[0] != "DATACLIENT": // col not English
		return 2
	}

	fmt.Println(CommCell)
	fmt.Println(GenTime)

	//添加自定义列
	*FiledsHeader = append(*FiledsHeader, []string{"COMMCELL", "APPLICATIONSIZE", "DATASUBCLIENT", "START TIME"}...)

	var t_tbodyData *goquery.Selection

	if cv8 {
		t_tbodyData = dom.Find("body > table:nth-child(12) > tbody > tr")
	} else {
		t_tbodyData = dom.Find("body > table:nth-child(13) > tbody > tr")
	}

	//fmt.Println(t_tbodyData.Html())

	headlen := len(*FiledsHeader)
	Data := make([][]string, 0)
	//循环表格数据
	t_tbodyData.Each(func(i int, s *goquery.Selection) {
		if i == 0 { // 0行是列名
			return
		}
		tmpSubData := make([]string, 0)
		sa := s.Find("td")
		sa.Each(func(i int, selection *goquery.Selection) {
			ss1 := selection.Text()
			//fmt.Println(selection.Attr("bgcolor"))
			//switch true {
			//case strings.Contains(ss1,"N/A"):
			//	ss1=strings.Replace(ss1,"N/A","",-1)
			//}
			ss1 = formatValues(i, ss1) // todo 格式化 数据
			tmpSubData = append(tmpSubData, ss1)
			//fmt.Println(i,ss1)
		})
		tmpSubData = append(tmpSubData, []string{CommCell, ApplicationSize, Subclient}...)

		if cv8 {
			//fmt.Printf("%d,%d\n", len(tmpSubData), len(*FiledsHeader))
			if len(tmpSubData) == headlen { // 因为列头删除了一列
				tmp := tmpSubData[:]
				tmpSubData = tmp[0:9]
				//tmpSubData = append(tmpSubData, []string{"", ""}...)
				tmpSubData = append(tmpSubData, tmp[10:]...)
				//fmt.Printf("%d,%v\n",len(tmpSubData),tmpSubData)
				Data = append(Data, tmpSubData)
			}

		} else {

			if len(tmpSubData) == headlen-1 { // 数据列数要与 列头一样长,少了start time
				Data = append(Data, tmpSubData)
			}
		}

	})

	f.Close()
	//fmt.Println(FiledsHeader)

	return GenSqls(Data, FiledsHeader, htmlfile) // 这里应该包括数据库插入 ，可以是后台 go 后面

}

func sliceToString(sl []string) (returnstring string) {
	sl1 := sl[0:len(sl)]
	for i := 0; i < len(sl1); i++ {
		switch true {
		case i == len(sl1)-1:
			returnstring = returnstring
			returnstring = returnstring + "\"" + sl1[i] + "\"" + ")"
			//returnstring = returnstring + "\"" + sl1[i] + "\""
			break
		case i < len(sl)-1:
			returnstring = returnstring
			returnstring = returnstring + "\"" + sl1[i] + "\"" + ","
			break
		}
	}
	return
}

func sliceToStringValue(sl []string) (returnstring string) {
	sl1 := sl[0:len(sl)]
	for i := 0; i < len(sl1); i++ {
		switch true {
		case i == len(sl1)-1:
			returnstring = returnstring
			//returnstring = returnstring + "\"" + sl1[i] + "\"" + ")"
			returnstring = returnstring + "\"" + sl1[i] + "\""
			break
		case i < len(sl)-1:
			returnstring = returnstring
			returnstring = returnstring + "\"" + sl1[i] + "\"" + ","
			break
		}
	}
	return
}

func updatesliceToString(sl []string, value []string) (returnstring string) {
	wherestring := ""
	//fmt.Println(len(sl),len(value))
	if len(sl)-1 == len(value) { // header  多start time
		for i := 0; i < len(sl); i++ {
			if len(value) > i {
				if value[i] == "" { // 空值路过
					continue
				}
			}

			if i == len(sl)-1 {
				//returnstring = returnstring + "\"" + sl[i] + "\"" + "='" + value[i] + "'"
				//wherestring = wherestring + "\"" + sl[i] + "\"" + "='" + value[i] + "'"
				returnstring = returnstring + "\"" + sl[i] + "\"" + "=%s"
				wherestring = wherestring + "\"" + sl[i] + "\"" + "=%s"
			} else {
				returnstring = returnstring + "\"" + sl[i] + "\"" + "='" + value[i] + "',"
				wherestring = wherestring + "\"" + sl[i] + "\"" + "='" + value[i] + "' and "
			}

		}
	}
	returnstring = returnstring + " where " + wherestring
	return
}

func GenSqls(Data [][]string, header *[]string, htmlfile string) (resultNum int) {

	baseSql := "insert into HF_BACKUPDETAIL(" + sliceToString(*header) + " values("
	updateSql := "update HF_BACKUPDETAIL set "
	//fmt.Println(Data)
	//fmt.Println(updateSql)
	//fmt.Println(baseSql)
	sqlArr := make([]map[string]string, 0)
	for _, value := range Data {
		tmpmap := make(map[string]string, 0)
		//fmt.Println(index, baseSql+sliceToString(value))
		valueStr := strings.Replace(sliceToStringValue(value), "\"", "'", -1)
		//fmt.Println(valueStr)
		//sqlArr = append(sqlArr, baseSql+valueStr)
		tmpmap["insert"] = baseSql + valueStr
		tmpmap["update"] = updateSql + updatesliceToString(*header, value)
		sqlArr = append(sqlArr, tmpmap)
	}

	return stmtSql(sqlArr, htmlfile)
}

func stmtSql(sqlArr []map[string]string, htmlfile string) (resultNum int) {
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
	if len(sqlArr) == 0 {
		return 3
	}
	//ss := []string{"to_date('2019-08-25 22:11:11','yyyy-mm-dd hh24:mi:ss')"}
	for i := 0; i < len(sqlArr); i++ {
		//fmt.Println(fmt.Sprintf(sqlArr[i]["update"], StartTimeArr[i], StartTimeArr[i]))
		//fmt.Println(fmt.Sprintf(sqlArr[i]["insert"]+",%s)", StartTimeArr[i]))
		//fmt.Println(sqlArr[i]["insert"]+",:1)",StartTimeArr[i])
		//r, err := db.Exec(fmt.Sprintf("insert into HF_BACKUPDETAIL(\"START TIME\") values(%s)", "to_date('2019-08-25 22:11:11','yyyy-mm-dd hh24:mi:ss')"))
		//fmt.Println(r,err)
		//ctx,cancel:=context.WithTimeout(context.Background(),20*time.Second)
		//r,err:=db.ExecContext(ctx,sqlArr[i]["insert"]+",:1)",StartTimeArr[i])
		//cancel()
		Result, err := db.Exec(fmt.Sprintf(sqlArr[i]["update"], StartTimeArr[i], StartTimeArr[i]))
		//fmt.Println(Result.LastInsertId())
		//fmt.Println(Result.RowsAffected())
		//fmt.Println("===========", r, err)
		if err != nil {
			Log.Error("HtmlFile:%s ,Sql Exec Err:%s", htmlfile, err)
			resultNum = 4
			continue
		}
		updatenum = updatenum + 1
		r, e := Result.RowsAffected()
		if err != nil {
			Log.Error("HtmlFile:%s ,Sql Exec Err:%s", htmlfile, e)
			resultNum = 4
			continue
		}

		if r == 0 {
			_, err := db.Exec(fmt.Sprintf(sqlArr[i]["insert"]+",%s)", StartTimeArr[i]))
			if err != nil {
				Log.Error("HtmlFile:%s ,Sql Exec Err:%s", htmlfile, err)
				resultNum = 4
				continue
			}
			insertnum = insertnum + 1
		}

	}
	Log.Printf("htmlfile : %s ,total sql:%d,success update sql:%d,insert sql:%d\n", htmlfile, len(sqlArr), updatenum, insertnum)
	return

}
