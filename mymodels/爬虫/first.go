package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"sync"
)

// todo https://godoc.org/github.com/PuerkitoBio/goquery

const (
	MasterUrl = "http://thzu.net/"
	MasterDir = "D:\\work\\image\\"
	pages     = 4
)

var masterchan chan *goquery.Document = make(chan *goquery.Document, pages)

type Links struct {
	L *sync.Mutex
	M map[string]string // 下载图片目录  链接地址
}

func ParsMasterweb(dom *goquery.Document) { //解析第一层主页
	a, err := dom.Find("table[summary]").Html()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

func UnLinks() {
	for {
		select {
		case dom, ok := <-masterchan:
			if ok {
				go ParsMasterweb(dom)
			} else {
				break
			}
		}
	}
}
func ForumGet() {
	for i := 1; i < pages; i++ {
		url := fmt.Sprintf("%sforum-181-%d.html", MasterUrl, i)
		log.Printf("request : %s\n", url)
		masterchan <- UrlDomGet(url)
	}
	close(masterchan)
}
func UrlDomGet(url string) *goquery.Document {

	client := &http.Client{}

	request, _ := http.NewRequest("GET", url, nil)

	//request.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//request.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	//request.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	//request.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	//request.Header.Set("Cache-Control","max-age=0")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", "chrome 67")

	response, err := client.Do(request)
	if err != nil {
		log.Printf("[Error]:%s, url:%s", err, url)
	}

	//if response.StatusCode == 200 {
	//	body, _ := ioutil.ReadAll(response.Body)
	//	fmt.Println(string(body))
	//}

	if response.StatusCode != 200 {
		//log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		log.Printf("status code error: %d %s", response.StatusCode, response.Status)
	}

	// Load the HTML document
	dom, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		//log.Fatal(err)
		log.Printf("[Error]:%s, url:%s", err, url)
	}
	//fmt.Println(dom)
	return dom

	//// Find the review items // 同级元素连接写， 父子级中间有空格
	//doc.Find("ul.BookList.PubIndex-recommends li").Each(func(i int, s *goquery.Selection) {
	//	// For each item found, get the band and title
	//	link, _ := s.Find("a").Attr("href") //获取属性值
	//	sonel := s.Find("a .Image")
	//	imagelink, _ := sonel.Attr("src")
	//	imagelinks = append(imagelinks, imagelink)
	//	//title, _ := sonel.Attr("alt")
	//	title := s.Find(".BookItem-title").Text() //获取文本
	//	fmt.Printf("NO.%d: link:%s - title:%-16s - imagelink:%-50s\n", i+1, link, title, imagelink)
	//})
	//for _, url := range imagelinks {
	//	downfile(url)
	//}
}

func downfile(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	fp := string(filepath.Join("c:\\", "1"))

	err = ioutil.WriteFile(fp, body, 0777)
	if err != nil {
		fmt.Printf("%v fp:[%v]\n", err.Error(), fp)
		return
	}
	fmt.Printf("Download: %+v\n", 1)
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	ForumGet()
	UnLinks()
	select {}
}
