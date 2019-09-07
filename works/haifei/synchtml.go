package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"path/filepath"
	"time"
	"tjtools/utils"
)

//https://github.com/360EntSecGroup-Skylar/excelize
// 中文 https://xuri.me/excelize/zh-hans/
// https://godoc.org/github.com/360EntSecGroup-Skylar/excelize#File.GetSheetName

const (
	//ReadFile    = "10.155.2.4_yuebao_English.html"
	htmlfileReg = "D:\\work\\project-dev\\src\\godev\\htmlData\\*.html"
	htmlBakDir  = "D:\\work\\project-dev\\src\\godev\\htmlData\\bak\\"
	timeinter   = 10
	timeLayout  = "2006-01-02 15:04:05"
)

var MoveFileChan chan string = make(chan string, 0)
var startRunTime = time.Now().Unix()

func main() {
	go waitMovefile()
	for {
		files, _ := filepath.Glob(htmlfileReg)
		//fmt.Println(files) // contains a list of all files in the current directory
		d, h, m, s := utils.GetTime(time.Now().Unix() - startRunTime)
		log.Printf("本次读取到 %d个文件,%v\n", len(files), files)
		log.Printf("开始运行时间 %s,已经运行%d天%d小时%d分钟%d秒", time.Unix(startRunTime, 0).Format(timeLayout), d, h, m, s)
		if len(files) >= 1 {
			waitHtmlFile(files)
		}

		time.Sleep(time.Duration(timeinter) * time.Second)
	}
}

func MoveFile(file string) {
	err := os.MkdirAll(htmlBakDir, 0777)
	if err != nil {
		log.Printf("mkdir htmlbak Error DIR: %s, err:%s", htmlBakDir, err)
	}
	_, fileName := filepath.Split(file)
	log.Printf("move file old:%s,new:%s\n", file, filepath.Join(htmlBakDir, "\\", fileName))
	err = os.Rename(file, filepath.Join(htmlBakDir, "\\", fileName))
	if err != nil {
		log.Printf("osRename htmlbak Error DIR: %s, err:%s", htmlBakDir, err)
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
		log.Printf("covert file:%s\n", file)
		ReadHtml(file)
		log.Printf("finish covert file:%s\n", file)
	}

}

func ReadHtml(htmlfile string) {

	f, err := os.OpenFile(htmlfile, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Print(err)
	}

	dom, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		//log.Fatal(err)
		log.Printf("[Error]:%s", err)
	}
	//fmt.Println(dom)

	FiledsHeader := make([]string, 0) // 列头，序号为KEY
	t_tbodyHeader := dom.Find("body > table:nth-child(13) > tbody > tr:nth-child(1) > td")
	//fmt.Println(t_tbody.Html())

	t_tbodyHeader.Each(func(i int, s *goquery.Selection) {
		sa := s.Text()
		FiledsHeader = append(FiledsHeader, sa)
	})

	//fmt.Println(FiledsHeader)
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
	for i := 0; i < len(sl); i++ {
		if i == len(sl)-1 {
			returnstring = returnstring + sl[i] + ")"
		} else {
			returnstring = returnstring + sl[i] + ","
		}
	}
	return
}

func genSql(Data [][]string, header []string) {

	baseSql := "insert into BCD(" + sliceToString(header) + ") value("

	//fmt.Println(baseSql)
	for index, value := range Data {
		fmt.Println(index, baseSql+sliceToString(value))
	}

}
