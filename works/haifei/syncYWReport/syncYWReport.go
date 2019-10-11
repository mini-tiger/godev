package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-oci8"
	"godev/works/haifei/syncYWReport/g"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
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
	timeLayout   = "2006-01-02 15:04:05"
	FieldsLen    = 21 // 列一共有几列,使用时 要 减一
	DetailTable  = "HF_YWREPORTDETAIL"
	SummaryTable = "HF_YWREPORTSUMMARY"
)

var MoveFileChan chan string = make(chan string, 0)
var startRunTime = time.Now().Unix()
var (
	CommCell  string // 客户端名
	GenTime   string
	StartTime string
	//ApplicationSize string
	Subclient string
	//StartTimeArr  []string = make([]string, 0)
	HtmlFile      string
	DetailSqlArr  []map[string]string
	SummarySqlArr []map[string]string
	Version       int
)

var Log *logDiy.Log1
var c Config

// 摘要表
//var FieldsMap map[string]string=map[string]string{"客户端":"ReportClient","主机名":"HOSTNAME","总作业数":"TOTALJOB",
//													"已完成":"COMPLETED","完成但有错误":"COMPLETIONERROR","完成但有警告":"COMPLETIONWARN",
//													"已终止":"TERMINATION","不成功":}

var SummaryFieldsMap map[int]string = g.SummaryFieldsMap
var SummaryFieldsMapCv10 map[int]string = g.SummaryFieldsMapCv10
var SummaryFieldsMapCv8 map[int]string = g.SummaryFieldsMapCv8
var SummaryFieldsMapPlus map[int]string = g.SummaryFieldsMapPlus
var DetailFieldsMap map[int]string = g.DetailFieldsMap
var DetailFieldsMapCv8 map[int]string = g.DetailFieldsMapCv8
var DetailFieldsMapPlus map[int]string = g.DetailFieldsMapPlus

var StatusColors map[string]g.StatusColor = g.StatusColors
var StatusColorsCv8 map[string]g.StatusColor = g.StatusColorsCv8

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
		HtmlFile = file
		resultNum := ReadHtml()
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
			MoveFileChan <- file
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

func formatStr(s string) string {
	return strings.Replace(strings.TrimSpace(s), "\n", "", -1)
}

func formatDetailValues(index int, s string) (rstr string) {
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
		StartTimeValue = g.FormatTime(StartTimeValue)
		//StartTimeArr = append(StartTimeArr, StartTimeValue)
		StartTime = StartTimeValue

		//StartTime = fmt.Sprintf("to date(%s mm/dd/yyyy hh24:mi:ss)",StartTime)
		//StartTime = "2019-08-12 04:00:00"
		break
		//case index == 8: // starttime 赋值 全局变量，对应自定义列
		//	as := ""
		//	if strings.Contains(s, "(") {
		//		as = strings.Split(s, "(")[0]
		//	} else {
		//		as = s
		//	}
		//
		//	switch true {
		//	case strings.Contains(as, "TB"):
		//		as = strings.Split(as, " TB")[0]
		//		fl, err := strconv.ParseFloat(as, 64)
		//		if err != nil {
		//			Log.Printf("应用程序大小转换float失败 err:%s\n", err)
		//			break
		//		}
		//
		//		//am:=strconv.FormatFloat(fl*1024,'E',-1,64)
		//		ApplicationSize = fmt.Sprintf("%.2f", fl*1024*1014)
		//	case strings.Contains(as, "GB"):
		//		as = strings.Split(as, " GB")[0]
		//		fl, err := strconv.ParseFloat(as, 64)
		//		if err != nil {
		//			Log.Printf("应用程序大小转换float失败 err:%s\n", err)
		//			break
		//		}
		//
		//		//am:=strconv.FormatFloat(fl*1024,'E',-1,64)
		//		ApplicationSize = fmt.Sprintf("%.2f", fl*1024)
		//
		//	case strings.Contains(as, "MB"):
		//		as = strings.Split(as, " MB")[0]
		//		fl, err := strconv.ParseFloat(as, 64)
		//		if err != nil {
		//			Log.Printf("应用程序大小转换float失败 err:%s\n", err)
		//			break
		//		}
		//		ApplicationSize = fmt.Sprintf("%.2f", fl)
		//	case strings.Contains(as, "KB"):
		//		as = strings.Split(as, " KB")[0]
		//		fl, err := strconv.ParseFloat(as, 64)
		//		if err != nil {
		//			Log.Printf("应用程序大小转换float失败 err:%s\n", err)
		//			break
		//		}
		//		ApplicationSize = fmt.Sprintf("%.2f", fl/1024)
		//	case strings.Contains(as, "Not Run because of another job running for the same subclient"):
		//		break
		//	default:
		//
		//		application, e := strconv.Atoi(strings.TrimSpace(as))
		//		if e != nil {
		//			Log.Error("change type string:%s,err:%s", as, e)
		//		}
		//		if application == 0 {
		//			ApplicationSize = "0"
		//		}
		//
		//	}

		//break
	}
	return
}

func ReadHtml() (resultNum int) {

	f, err := os.OpenFile(HtmlFile, os.O_RDONLY, 0755)
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
	cvSelection := dom.Find("body:contains(备份作业摘要报告)")
	if cvSelection.Size() < 1 {
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
	//fmt.Println(CommCell)

	// 生成的时间
	//GenTimeSource := dom.Find("body:contains(备份作业摘要报告)")
	cvSelection.Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			//fmt.Println(HtmlFile,selection.Text())
			//fmt.Println(strings.Contains(selection.Text(),"generated"))

			ss := strings.Split(strings.Split(selection.Text(), "备份作业摘要报告")[1], "上生成的报表版本")[0]
			//fmt.Println(strings.Split(selection.Text(), "备份作业摘要报告")[1])
			//fmt.Println(strings.Split(selection.Text(), "备份作业摘要报告"))
			//fmt.Println(ss)
			if strings.Contains(ss, "\n") {
				ss = strings.Split(ss, "\n")[1]
			}
			GenTime = formatStr(ss)
			//strings.TrimSpace(ss)
			//GenTime = strings.Replace(ss, "\n", "", -1)
			//fmt.Printf("GenTime :%+v\n", ss)
		}
	})

	//fmt.Println(GenTime)

	//FiledsHeader := new([]string) // 列头，序号为KEY

	SummaryDomFindStr := "body > table:nth-child(10) > tbody > tr"
	Summary_tbodyHeader := dom.Find(SummaryDomFindStr + ":nth-child(2) > td")

	detailDomFindStr := "body > table:nth-child(13) > tbody > tr"
	detail_tbodyHeader := dom.Find(detailDomFindStr + ":nth-child(1) > td")

	version := 11
	//有三种格式的HTML
	switch true {
	case len(detail_tbodyHeader.Nodes) == 0:
		version = 8
		GenTime = strings.Split(GenTime, "CommCell")[0]
		GenTime = strings.Split(GenTime, "生成于")[1]
		GenTime = strings.TrimSpace(GenTime)

		SummaryDomFindStr = "body > table:nth-child(9) > tbody > tr"
		Summary_tbodyHeader = dom.Find(SummaryDomFindStr + ":nth-child(2) > td")
		detailDomFindStr = "body > table:nth-child(12) > tbody > tr"
		detail_tbodyHeader = dom.Find(detailDomFindStr + ":nth-child(1) > td")
		break
	case len(Summary_tbodyHeader.Nodes) == 18:
		version = 10
		break
	default:
		version = 11
	}
	Version = version

	GenTime = g.FormatTime(GenTime)
	/*
		version 10 : summery:18cols,detail: 17cols
				11 :20 17
				8 :16 16
	*/

	// 循环找出列头
	//detail_tbodyHeader.Each(func(i int, s *goquery.Selection) {
	//	sa := s.Text()
	//	sa = formatFields(i, sa) // todo 格式化 列头
	//	*FiledsHeader = append(*FiledsHeader, sa)
	//})

	// 如果是CV8版本，手动修改列名, 删除 size of Backup cols 这列
	//switch version {
	//case 8:
	//
	//	break
	//case 10:
	//	break
	//case 11:
	//	break
	//
	//}

	//fmt.Println("222222222",*FiledsHeader)
	//for i := 0; i < len(*FiledsHeader); i++ {
	//	fmt.Printf("%d, value:%s\n", i, (*FiledsHeader)[i])
	//}

	//switch true {
	//case len(*FiledsHeader) < FieldsLen-6: // col not enough,运维报告中不用判断
	//	//fmt.Println(FiledsHeader)
	//	//fmt.Println(len(*FiledsHeader))
	//	return 1
	//case (*FiledsHeader)[0] != "DATACLIENT": // col not English
	//	return 2
	//}
	fmt.Printf("htmfile:%s, commcell:%s, GenTime:%s, version:%d,len(sum):%d, len(detail):%d", path.Base(HtmlFile), CommCell, GenTime, version, len(Summary_tbodyHeader.Nodes), len(detail_tbodyHeader.Nodes))
	fmt.Println()
	//fmt.Println(CommCell)
	//fmt.Println(GenTime)
	//fmt.Println(cv8)

	//添加自定义列
	//*FiledsHeader = append(*FiledsHeader, []string{"COMMCELL", "APPLICATIONSIZE", "DATASUBCLIENT", "START TIME"}...)
	var rn int
	switch version {

	case 11:

		GenDetailData(detailDomFindStr, dom, DetailFieldsMap, StatusColors) // 生成 详细数据,version 11 使用默认DetailFieldsMap
		GenSummaryData(SummaryDomFindStr, dom, SummaryFieldsMap)            // 生成 详细数据,version 11 使用默认DetailFieldsMap

		rn = GenSqls(DetailSqlArr, DetailTable)
		//if rn > 0 {
		//	return rn
		//}
		rn = GenSqls(SummarySqlArr, SummaryTable)
		//fmt.Println(SummarySqlArr)
		//if rn > 0 {
		//	return rn
		//}

	case 10:
		GenDetailData(detailDomFindStr, dom, DetailFieldsMap, StatusColors) // 生成 详细数据,version 10 使用默认DetailFieldsMap
		GenSummaryData(SummaryDomFindStr, dom, SummaryFieldsMapCv10)        // 生成 详细数据,version 10 使用默认DetailFieldsMap

		rn = GenSqls(DetailSqlArr, DetailTable)
		//if rn > 0 {
		//	return rn
		//}
		rn = GenSqls(SummarySqlArr, SummaryTable)
		//fmt.Println(SummarySqlArr)
		//if rn > 0 {
		//	return rn
		//}
	case 8:

		GenDetailData(detailDomFindStr, dom, DetailFieldsMapCv8, StatusColorsCv8) // 生成 详细数据,version 8 使用默认DetailFieldsMap
		GenSummaryData(SummaryDomFindStr, dom, SummaryFieldsMapCv8)               // 生成 详细数据,version 8 使用默认DetailFieldsMap

		rn = GenSqls(DetailSqlArr, DetailTable)
		//if rn > 0 {
		//	return rn
		//}
		rn = GenSqls(SummarySqlArr, SummaryTable)
		//fmt.Println(SummarySqlArr)
		//if rn > 0 {
		//	return rn
		//}
	}

	f.Close()
	DetailSqlArr = make([]map[string]string, 0)
	SummarySqlArr = make([]map[string]string, 0)
	//fmt.Println(FiledsHeader)

	//return GenSqls(Data, FiledsHeader, htmlfile) // 这里应该包括数据库插入 ，可以是后台 go 后面
	return rn
}

func GenSummaryData(domstr string, dom *goquery.Document, FiledsHeader map[int]string) (Data [][]string) {
	var t_tbodyData *goquery.Selection

	t_tbodyData = dom.Find(domstr)
	//Data = make([][]string, 0)
	//Log.Printf("HtmlFile:%s,Summary rows:%d \n", HtmlFile, len(t_tbodyData.Nodes))
	t_tbodyData.Each(func(i int, rowsele *goquery.Selection) {
		if i < 2 { // 0行是 摘要 1行是列名
			return
		}

		tmpSubData := make(map[string]string, 0)
		cols := rowsele.Find("td")
		if len(FiledsHeader) != len(cols.Nodes) { //
			//Log.Error("htmfile:%s ,summary no eq rows:%d", HtmlFile, i)
			return
		}
		cols.Each(func(colnum int, colsele *goquery.Selection) {
			ss1 := colsele.Text()
			//fmt.Println(selection.Attr("bgcolor"))
			//switch true {
			//case strings.Contains(ss1,"N/A"):
			//	ss1=strings.Replace(ss1,"N/A","",-1)
			//}

			//ss1 = formatDetailValues(colnum, ss1) // todo 格式化 数据
			ss1 = strings.TrimSpace(ss1)
			if strings.Contains(FiledsHeader[colnum], "STARTTIME") {
				ss1 = g.FormatTime(ss1)
				if _, err := strconv.Atoi(ss1); err != nil {
					Log.Error("HTMLFILE:%s,headr:%s, 转换数字失败: 实际字符：%s", HtmlFile, "START TIME", ss1)
					ss1 = "NULL"
					return
				}
			}

			if strings.Contains(FiledsHeader[colnum], "FAILEOBJECT") || strings.Contains(FiledsHeader[colnum], "FAILEDFOLDERS") ||
				strings.Contains(FiledsHeader[colnum], "TOTALJOB") || strings.Contains(FiledsHeader[colnum], "COMPLETED") ||
				strings.Contains(FiledsHeader[colnum], "COMPLETEDWITHERRORS") ||
				strings.Contains(FiledsHeader[colnum], "COMPLETEDWITHWARNINGS") ||
				strings.Contains(FiledsHeader[colnum], "KILLED") ||
				strings.Contains(FiledsHeader[colnum], "UNSUCCESSFUL") ||
				strings.Contains(FiledsHeader[colnum], "RUNNING") ||
				strings.Contains(FiledsHeader[colnum], "DELAYED") ||
				strings.Contains(FiledsHeader[colnum], "NORUN") ||
				strings.Contains(FiledsHeader[colnum], "NOSCHEDULE") ||
				strings.Contains(FiledsHeader[colnum], "COMMITTED") ||
				strings.Contains(FiledsHeader[colnum], "COMMITTED") ||
				strings.Contains(FiledsHeader[colnum], "PROTECTEDOBJECTS") {
				ss1 = strings.Join(strings.Split(ss1, ","), "")
				if _, err := strconv.Atoi(ss1); err != nil {
					Log.Error("HTMLFILE:%s,headr:%s, 转换数字失败", HtmlFile, FiledsHeader[colnum])
					ss1 = "NULL"
					return
				}

			}

			//fmt.Println(i,colnum,FiledsHeader[colnum],strings.Contains(FiledsHeader[colnum],"Failed Folder"))
			//
			//fmt.Println(strings.Split())
			tmpSubData[FiledsHeader[colnum]] = ss1
			//tmpSubData = append(tmpSubData, ss1)
			//fmt.Println(i,ss1)
		})
		tmpSubData[SummaryFieldsMapPlus[20]] = CommCell
		tmpSubData[SummaryFieldsMapPlus[21]] = GenTime
		//fmt.Println(tmpSubData)
		SummarySqlArr = append(SummarySqlArr, tmpSubData)
	})

	return

}

//func VersionDetailFields(mapFields map[int]string ,max int) (DetailFields []string) {
//
//}

func GenDetailData(domstr string, dom *goquery.Document, FiledsHeader map[int]string, StatusColors map[string]g.StatusColor) (Data [][]string) {
	var t_tbodyData *goquery.Selection

	t_tbodyData = dom.Find(domstr)
	//MaxRows:=len(t_tbodyData.Nodes)
	//body > table:nth-child(13) > tbody > tr
	//fmt.Println(t_tbodyData.Html())

	//headlen := len(FiledsHeader)
	Data = make([][]string, 0)
	//循环表格数据
	t_tbodyData.Each(func(i int, rowsele *goquery.Selection) {
		if i == 0 { // 0行是列名
			return
		}
		var rowjobtype string
		//var rowstarttimeformat string //  行格式化后的 unixstamp
		//var rowsubclient string
		var rowReasonforfailure string //行 失败原因
		var rowsolvetype string        // 行解决状态

		rowColor, exists := rowsele.Attr("bgcolor")
		if !exists {
			//找不到颜色的，可能是失败原因的行
			return
		}

		Statuscolor, e := StatusColors[rowColor]
		//fmt.Println(Statuscolor)

		if !e {
			Log.Error("没找到颜色 Htmfile:%s,detail_rows:%d", HtmlFile, i)
		}
		//需要判断是不是 超过最大行数

		switch Statuscolor.Status {
		case 2:
			return
		case 1:

			if len(rowsele.Next().Find(" td").Nodes) == 1 {
				//fmt.Println(i, rowsele.Next().Find(" td").Text())
				rowjobtype = Statuscolor.ColorStatus
				rowsolvetype = "未解决"
				rowReasonforfailure = rowsele.Next().Find(" td").Text()
			} else {
				rowjobtype = Statuscolor.ColorStatus
				rowsolvetype = "未解决"
				rowReasonforfailure = "NULL"
			}
		case 0:
			rowjobtype = Statuscolor.ColorStatus
			rowsolvetype = "NULL"
			rowReasonforfailure = "NULL"

		}

		//tmpSubData := make([]string, 0)
		tmpSubData := make(map[string]string, 0)
		cols := rowsele.Find("td")

		if Version != 8 {
			if len(FiledsHeader) != len(cols.Nodes) {
				//Log.Error("htmfile:%s ,no eq rows:%d", HtmlFile, i)
				return
			}
		} else {
			if len(FiledsHeader) != len(cols.Nodes)-1 { // 备份大小列 跳过
				//Log.Error("htmfile:%s ,no eq rows:%d", HtmlFile, i)
				return
			}

		}

		cols.Each(func(colnum int, colsele *goquery.Selection) {
			if Version == 8 && colnum == 9 {
				return
			}
			ss1 := colsele.Text()
			//fmt.Println(selection.Attr("bgcolor"))
			//switch true {
			//case strings.Contains(ss1,"N/A"):
			//	ss1=strings.Replace(ss1,"N/A","",-1)
			//}

			ss1 = formatDetailValues(colnum, ss1) // todo 格式化 数据

			//fmt.Println(i,colnum,FiledsHeader[colnum],strings.Contains(FiledsHeader[colnum],"Failed Folder"))
			if strings.Contains(FiledsHeader[colnum], "Failed Object") || strings.Contains(FiledsHeader[colnum], "Failed Folder") {
				ss1 = strings.Join(strings.Split(ss1, ","), "")
				if sslvalue, err := strconv.Atoi(ss1); err == nil && sslvalue > 0 {
					rowjobtype = "失败"
					rowsolvetype = "未解决"
				}

			}

			if strings.Contains(FiledsHeader[colnum], "Data Size Change") {
				if cv, ce := colsele.Attr("bgcolor"); ce {
					tmpStatcolor, ee := StatusColors[cv]
					//fmt.Println(i,colnum,ee,cv,tmpStatcolor,strings.Contains(tmpStatcolor.ColorStatus,"数据大小按 10% 或更多增加/减少"))
					if ee {
						if strings.Contains(tmpStatcolor.ColorStatus, "数据大小按 10% 或更多增加/减少") {
							rowjobtype = tmpStatcolor.ColorStatus
						}
					}
				}

			}
			//
			//fmt.Println(strings.Split())
			tmpSubData[FiledsHeader[colnum]] = ss1
			//tmpSubData = append(tmpSubData, ss1)
			//fmt.Println(i,ss1)
		})
		tmpSubData[DetailFieldsMapPlus[17]] = CommCell
		tmpSubData[DetailFieldsMapPlus[18]] = GenTime
		tmpSubData[DetailFieldsMapPlus[19]] = rowjobtype
		tmpSubData[DetailFieldsMapPlus[20]] = StartTime
		tmpSubData[DetailFieldsMapPlus[21]] = Subclient
		tmpSubData[DetailFieldsMapPlus[22]] = rowReasonforfailure
		tmpSubData[DetailFieldsMapPlus[23]] = rowsolvetype

		//tmpSubData = append(tmpSubData, []string{CommCell, GenTime, rowjobtype, StartTime, Subclient, rowReasonforfailure, rowsolvetype}...)
		//fmt.Printf("%+v\n", tmpSubData)
		//fmt.Println()
		DetailSqlArr = append(DetailSqlArr, tmpSubData)
		//if version == 8 {
		//	//fmt.Printf("%d,%d\n", len(tmpSubData), len(*FiledsHeader))
		//	if len(tmpSubData) == headlen { // 因为列头删除了一列
		//		tmp := tmpSubData[:]
		//		tmpSubData = tmp[0:9]
		//		//tmpSubData = append(tmpSubData, []string{"", ""}...)
		//		tmpSubData = append(tmpSubData, tmp[10:]...)
		//		//fmt.Printf("%d,%v\n",len(tmpSubData),tmpSubData)
		//		Data = append(Data, tmpSubData)
		//	}
		//
		//} else {
		//
		//	if len(tmpSubData) == headlen-1 { // 数据列数要与 列头一样长,少了start time
		//		Data = append(Data, tmpSubData)
		//	}
		//}

	})
	//fmt.Println(Data)
	return
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

func sliceToStringValue(sl map[string]string) (returnstring string) {
	//sl1 := sl[0:len(sl)]
	var keys []string = make([]string, 0)
	for k, _ := range sl {
		keys = append(keys, k)

	}
	sort.Strings(keys)
	valueStr := ""
	colsStr := ""
	for i, v := range keys {
		if sl[v] == "NULL" || sl[v] == "" {
			continue
		}
		switch true {
		case i == len(sl)-1:
			//returnstring = returnstring
			//returnstring = returnstring + "\"" + sl1[i] + "\"" + ")"
			colsStr = colsStr + "\"" + v + "\")"
			valueStr = valueStr + "'" + sl[v] + "')"

			break
		case i < len(sl)-1:
			//returnstring = returnstring
			colsStr = colsStr + "\"" + v + "\"" + ","
			valueStr = valueStr + "'" + sl[v] + "'" + ","
			break
		}

	}
	//fmt.Println(len(keys),len(sl))
	returnstring = colsStr + " values(" + valueStr

	return
}

func updatesliceToString(sl map[string]string) (returnstring string) {

	//fmt.Println(len(sl),len(value))
	//if len(sl)-1 == len(value) { // header  多start time
	//	for i := 0; i < len(sl); i++ {
	//		if len(value) > i {
	//			if value[i] == "" { // 空值路过
	//				continue
	//			}
	//		}
	//
	//		if i == len(sl)-1 {
	//			//returnstring = returnstring + "\"" + sl[i] + "\"" + "='" + value[i] + "'"
	//			//wherestring = wherestring + "\"" + sl[i] + "\"" + "='" + value[i] + "'"
	//			returnstring = returnstring + "\"" + sl[i] + "\"" + "=%s"
	//			wherestring = wherestring + "\"" + sl[i] + "\"" + "=%s"
	//		} else {
	//			returnstring = returnstring + "\"" + sl[i] + "\"" + "='" + value[i] + "',"
	//			wherestring = wherestring + "\"" + sl[i] + "\"" + "='" + value[i] + "' and "
	//		}
	//
	//	}
	//}
	//returnstring = returnstring + " where " + wherestring

	var keys []string = make([]string, 0)
	for k, _ := range sl {
		keys = append(keys, k)

	}
	sort.Strings(keys)
	wherestring := ""
	colsStr := ""
	for i, v := range keys {
		if sl[v] == "NULL" || sl[v] == "" {
			continue
		}
		switch true {
		case i == len(sl)-1:
			//returnstring = returnstring
			//returnstring = returnstring + "\"" + sl1[i] + "\"" + ")"
			colsStr = colsStr + "\"" + v + "\"='" + sl[v] + "'"
			wherestring = wherestring + "\"" + v + "\"='" + sl[v] + "'"

			break
		case i < len(sl)-1:
			//returnstring = returnstring
			colsStr = colsStr + "\"" + v + "\"='" + sl[v] + "',"
			wherestring = wherestring + "\"" + v + "\"='" + sl[v] + "' and "
			break
		}

	}
	returnstring = colsStr + " where " + wherestring
	return
}

func GenSqls(DetailSqlArrCopy []map[string]string, tablename string) (resultNum int) {

	baseSql := "insert into " + tablename + "("
	updateSql := "update " + tablename + " set "
	//fmt.Println(Data)
	//fmt.Println(updateSql)
	//fmt.Println(baseSql)
	sqlArr := make([]map[string]string, 0)
	for _, value := range DetailSqlArrCopy {
		tmpmap := make(map[string]string, 0)
		//fmt.Println(index, baseSql+sliceToString(value))
		valueStr := sliceToStringValue(value)
		//fmt.Println(valueStr)
		//sqlArr = append(sqlArr, baseSql+valueStr)
		tmpmap["insert"] = baseSql + valueStr
		tmpmap["update"] = updateSql + updatesliceToString(value)
		sqlArr = append(sqlArr, tmpmap)
	}
	//for _, v := range sqlArr {
	//fmt.Println(v["insert"])
	//fmt.Println(v["update"])
	//}
	return stmtSql(sqlArr, tablename)
	//return 1
}

func stmtSql(sqlArr []map[string]string, tablename string) (resultNum int) {
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

	var updatenum, insertnum int
	//insertnum := 0
	if len(sqlArr) == 0 {
		return 3
	}
	//ss := []string{"to_date('2019-08-25 22:11:11','yyyy-mm-dd hh24:mi:ss')"}
	for i := 0; i < len(sqlArr); i++ {
		//fmt.Println(sqlArr[i]["update"])
		//fmt.Println(fmt.Sprintf(sqlArr[i]["update"], StartTimeArr[i], StartTimeArr[i]))
		//fmt.Println(fmt.Sprintf(sqlArr[i]["insert"]+",%s)", StartTimeArr[i]))
		//fmt.Println(sqlArr[i]["insert"]+",:1)",StartTimeArr[i])
		//r, err := db.Exec(fmt.Sprintf("insert into HF_BACKUPDETAIL(\"START TIME\") values(%s)", "to_date('2019-08-25 22:11:11','yyyy-mm-dd hh24:mi:ss')"))
		//fmt.Println(r,err)
		//ctx,cancel:=context.WithTimeout(context.Background(),20*time.Second)
		//r,err:=db.ExecContext(ctx,sqlArr[i]["insert"]+",:1)",StartTimeArr[i])
		//cancel()
		Result, err := db.Exec(sqlArr[i]["update"])
		//fmt.Println(Result.LastInsertId())
		//fmt.Println(Result.RowsAffected())
		//fmt.Println("===========", r, err)
		if err != nil {
			fmt.Println(sqlArr[i]["update"])
			Log.Error("HtmlFile:%s ,Sql Exec Err:%s", HtmlFile, err)
			resultNum = 4
			continue
		}
		updatenum = updatenum + 1
		r, e := Result.RowsAffected()
		if err != nil {
			Log.Error("HtmlFile:%s ,Sql Exec Err:%s", HtmlFile, e)
			resultNum = 4
			continue
		}

		if r == 0 {
			_, err := db.Exec(sqlArr[i]["insert"])

			if err != nil {
				fmt.Println(sqlArr[i]["insert"])
				Log.Error("HtmlFile:%s ,Sql Exec Err:%s", HtmlFile, err)
				resultNum = 4
				continue
			}
			insertnum = insertnum + 1
		}

	}
	Log.Printf("htmlfile : %s ,tablname:%s,total sql:%d,success update sql:%d,insert sql:%d\n", HtmlFile, tablename, len(sqlArr), updatenum, insertnum)
	return

}
