package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mini-tiger/tjtools/control"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"time"
)

var userAgentSlice = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:2.0b13pre) Gecko/20110307 Firefox/4.0b13pre"}

var sema = control.NewSemaphore(2)
var url = "https://www.domp4.cc/html/Wjg8GJ88888J.html"

func main() {
	dom := UrlDomGet()
	parse(dom)

}
func parse(dom *goquery.Document) {
	t_tbody_rows := dom.Find("#download1").Find("ul").Find("div")
	fmt.Println(t_tbody_rows.Length())

	t_tbody_rows.Each(func(i int, s *goquery.Selection) {
		child, exist := s.Find("a").Attr("href")
		//v := sa.Find(" td:last-child>span").Text() //需要两种判断日期的DOM结构，是否span 下是否有title
		if exist && strings.Contains(child, "magnet") {
			fmt.Println(child)
		}

		//childHref, exist := sa.Attr("href")
		//if exist {
		//
		//}

	})
}

func UrlDomGet() *goquery.Document {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, _ := http.NewRequest("GET", url, nil)

	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", userAgentSlice[rand.Intn(len(userAgentSlice))])

	response, err := client.Do(request)

	if err != nil {
		log.Printf("[Error]:request:%s, url:%s", err, url)
		return nil
	}

	//if response.StatusCode == 200 {
	//	body, _ := ioutil.ReadAll(response.Body)
	//	fmt.Println(string(body))
	//}

	if response.StatusCode != 200 {
		//log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		log.Printf("url :%s,status code error: %d %s", url, response.StatusCode, response.Status)
		return nil
	}

	// Load the HTML document
	dom, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		//log.Fatal(err)
		log.Printf("[Error]:response:%s, url:%s", err, url)
	}
	//fmt.Println(dom)
	return dom
}
