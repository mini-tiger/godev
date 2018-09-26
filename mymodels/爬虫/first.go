package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"path/filepath"
)
// todo https://godoc.org/github.com/PuerkitoBio/goquery

var imagelinks []string = make([]string,0)
func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://www.zhihu.com/pub/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items // 同级元素连接写， 父子级中间有空格
	doc.Find("ul.BookList.PubIndex-recommends li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		link, _ := s.Find("a").Attr("href") //获取属性值
		sonel := s.Find("a .Image")
		imagelink, _ := sonel.Attr("src")
		imagelinks=append(imagelinks,imagelink)
		//title, _ := sonel.Attr("alt")
		title := s.Find(".BookItem-title").Text() //获取文本
		fmt.Printf("NO.%d: link:%s - title:%-16s - imagelink:%-50s\n", i+1, link, title, imagelink)
	})
	for _,url:=range imagelinks{
		downfile(url)
	}
}


func downfile(url string)  {
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

	fp := string(filepath.Join("c:\\","1"))

	err = ioutil.WriteFile(fp, body, 0777)
	if err != nil {
		fmt.Printf("%v fp:[%v]\n", err.Error(), fp)
		return
	}
	fmt.Printf("Download: %+v\n", 1)
}

func main() {
	ExampleScrape()
}
