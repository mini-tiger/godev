package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

	"path/filepath"

	"time"
)

const (
	UA = "Golang Downloader from tao.com"
)

type PostJson struct {
	RunType   map[string]string `json:"RUNTYPE"`
	SolveType map[string]string `json:"SOLVETYPE"`
	StartTime map[string]string `json:"START TIME"`
	Page      int
	PageSize  int
}

func main() {

	url := "http://192.168.43.26:4444/crud/HF_YWMONITORTOTAL/exportData"

	Runtype := map[string]string{"key": "!=", "value": "已完成"}
	Solvetype := map[string]string{"key": "in", "value": "(0,1)"}
	//Starttime := map[string]string{"key": "gt", "value": "1597639852133"}
	sj := PostJson{Page: 1, PageSize: 6000, RunType: Runtype, SolveType: Solvetype}
	jsonBytes, err := json.Marshal(sj)
	if err != nil {
		log.Fatal(err)
	}
	DownFile(url, "C:\\work\\go-dev\\src\\godev\\mymodels\\1.xlsx", jsonBytes)

}
func DownFile(url, fp string, jsonStr []byte) {

	log.Printf("开始 download %s,url:%s", fp, url)

	client := &http.Client{}

	client = &http.Client{
		Timeout: 60 * time.Second,
	}

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	//request.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//request.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	//request.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	//request.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	//request.Header.Set("Cache-Control","max-age=0")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("content-type", "application/json;charset=UTF-8")
	//request.Header.Set("User-Agent", userAgentSlice[rand.Intn(len(userAgentSlice))])

	//defer func() {
	//	if err := recover(); err != nil {
	//		log.Printf("跳过url:%s,err:%s \n", err, url)
	//		c <- struct{}{}
	//		//panic(fmt.Sprintf("err:%s\n",url))
	//
	//	}
	//}()

	response, err := client.Do(request)
	if err != nil {

		return
	}

	if response.StatusCode != 200 {
		//log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		log.Printf("status code error: %d %s", response.StatusCode, response.Status)
		//c <- struct{}{}
		//wdownload.Done()
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		f := filepath.Dir(fp)
		fmt.Printf("body err: %s,dir: %s, url:%s\n", err.Error(), f, url)
		//c <- struct{}{}
		//wdownload.Done()
		return
	}

	//fp := string(filepath.Join("c:\\", "1"))

	err = ioutil.WriteFile(fp, body, 0777)
	if err != nil {
		fmt.Printf("%v fp:[%v]\n", err.Error(), fp)
		//c <- struct{}{}
		//wdownload.Done()
		return
	}
	fmt.Printf("Download 成功: %+v\n", fp)
	//c <- struct{}{}
	//wdownload.Done()
}
