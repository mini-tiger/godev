package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"runtime"
	"sync"
	"math/rand"
	"time"
	"os"
	"path/filepath"
)

// todo https://godoc.org/github.com/PuerkitoBio/goquery

const (
	MasterUrl   = "http://thzu.net/"
	MasterDir   = "c:\\work\\image\\"
	pages       = 3    //最多看3页的数据，3
	max_old     = 4    //最大7天前
	exist_cover = false //存在是否覆盖
)

var tmp_chan_web chan struct{} = make(chan struct{}, pages) //主页退出 通道
var tmp_chan chan struct{} = make(chan struct{}, 1) //最后下载图片种子后 退出，通道
var tnow = time.Now().Unix()
var w sync.WaitGroup
var masterchan chan *goquery.Document = make(chan *goquery.Document, pages)
var user_agent_slice = []string{
	"chrome 67",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:2.0b13pre) Gecko/20110307 Firefox/4.0b13pre"}

//var	dir_url map[string]string =make(map[string]string)// 下载图片目录  第二层链接地址
var dir_imageurls map[string][]string = make(map[string][]string)   // 下载图片目录  链接地址
var dir_torrenturls map[string][]string = make(map[string][]string) // 下载种子  链接地址
func checktime(ts, base_format string) bool {
	//base_format := "2006-01-02 15:04"
	parse_str_time, _ := time.Parse(base_format, ts) // todo 字符串转时间
	//fmt.Println(ts,base_format,parse_str_time,tnow < max_old*86400+parse_str_time.Unix())
	if tnow < max_old*86400+parse_str_time.Unix() {
		return true
	} else {
		return false
	}
}

func images_url(url string) ([]string, []string) {
	tmp_slice := make([]string, 0)
	tmp_slice1 := make([]string, 0)

	dom := UrlDomGet(fmt.Sprintf("%s%s", MasterUrl, url))
	td := dom.Find("table>tbody td.t_f") //todo 子元素选择器 不是直接上下级关系 的 中间有空格
	//td:=dom.Find("table>tbody>tr>td.t_f")//todo 子元素选择器 是直接上下级关系 的 > 号

	// 查找图片
	td.Find("img ").Each(func(i int, s *goquery.Selection) {
		img, ok := s.Attr("file")
		if ok {
			tmp_slice = append(tmp_slice, img)
		}
	})

	//查找种子链接
	td = dom.Find("p.attnm > a")
	td.Each(func(i int, s *goquery.Selection) {
		torrent_parenturl, ok := s.Attr("href")
		if ok {
			tmp_dom := UrlDomGet(fmt.Sprintf("%s%s", MasterUrl, torrent_parenturl)) //再次请求种子页面
			init_href := tmp_dom.Find("div.f_c div[style^=padding-left] a")         //在种子下载页面查找
			//fmt.Println(init_href.Html())
			torrent, ok := init_href.Attr("href")
			if ok {
				tmp_slice1 = append(tmp_slice1, torrent)
			}
		}
	})

	w.Done()
	return tmp_slice, tmp_slice1
}
func ParsMasterweb(dom *goquery.Document) { //解析第一层主页
	defer func() { //url  不能打开的 恢复机制
		if err := recover(); err != nil {
			//log.Printf("跳过err:%s \n",err)
			//panic(fmt.Sprintf("err:%s\n",url))
			w.Done()
		}
	}()
	t_tbody := dom.Find("table[summary]").Find("tbody")
	//log.Printf("request url:%s, tbody math:%d 个",dom.Url,t_tbody.Length())
	t_tbody.Each(func(i int, s *goquery.Selection) {
		sa := s.Find("tr>td.by").First().Find("em").Children()
		/*
		一种是
		<em><span><span title="2018-9-20">7&nbsp;天前</span></span></em>
		一种是
		<em><span>2018-9-13</span></em>
		*/
		if v, b := sa.Find("span").Attr("title"); b { //需要两种判断日期的DOM结构，是否span 下是否有title
			if checktime(v, "2006-1-2") { //时间是否在范围内
				dir_string := s.Find("tr>th>a").Text()
				//fmt.Println(dir_string)
				url_string, _ := s.Find("tr>th>a.s.xst").Attr("href")
				//fmt.Println(url_string)
				w.Add(1)
				dir_imageurls[dir_string], dir_torrenturls[dir_string] = images_url(url_string) //不加go 并发太大可能可能会拒绝连接
			}
		} else {
			aa, _ := sa.Html()
			if checktime(aa, "2006-1-2") {
				dir_string := s.Find("tr>th>a").Text()
				//fmt.Println(dir_string)
				url_string, _ := s.Find("tr>th>a.s.xst").Attr("href")
				//fmt.Println(url_string)
				w.Add(1)
				dir_imageurls[dir_string], dir_torrenturls[dir_string] = images_url(url_string) //不加go 否则可能会拒绝连接
			}
		}
	})

	w.Done()
}

func UnLinks() {
	for {
		select {
		case dom, ok := <-masterchan:
			if ok {
				w.Add(1) // todo 阻塞 二级以下页面解析
				tmp_chan_web<- struct{}{} //todo 阻塞一级页面解析
				go ParsMasterweb(dom)
			} else {
				break
			}
		}
	}
}
func ForumGet() {

	for i := 1; i <= pages; i++ {
		url := fmt.Sprintf("%sforum-181-%d.html", MasterUrl, i)
		log.Printf("request : %s\n", url)
		go func() {
			masterchan <- UrlDomGet(url)
		}()
	}
	for i:=1;i<=pages;i++{
		<-tmp_chan_web
	}
	//time.Sleep(10 * time.Second)
	close(masterchan)
}
func UrlDomGet(url string) *goquery.Document {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, _ := http.NewRequest("GET", url, nil)

	//request.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//request.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	//request.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	//request.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	//request.Header.Set("Cache-Control","max-age=0")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", user_agent_slice[rand.Intn(len(user_agent_slice))])


	defer func() {
		if err := recover(); err != nil {
			log.Printf("跳过url:%s,err:%s \n",err,url)
			//panic(fmt.Sprintf("err:%s\n",url))

		}
	}()


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

func downfile(url, fp string) {
	log.Printf("download %s,url:%s", fp, url)
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

	//fp := string(filepath.Join("c:\\", "1"))

	err = ioutil.WriteFile(fp, body, 0777)
	if err != nil {
		fmt.Printf("%v fp:[%v]\n", err.Error(), fp)
		return
	}
	fmt.Printf("Download: %+v\n", fp)
}

func makedir() {
	for k, _ := range dir_imageurls {
		err := os.MkdirAll(filepath.Join(MasterDir, k), 0777)
		if err != nil {
			log.Println("mkdirerr %s err:%s", k, err)
		}
	}
}
func exist(file string) bool {

	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			//fmt.Printf("文件: %s 不存在\n", file)
			return false
		}
	} else {
		//fmt.Printf("文件: %s 存在\n", file)
		return true
	}

	return false
}
func downloadall() {
	makedir()
	for k, v := range dir_imageurls {
		m_dir := filepath.Join(filepath.Join(MasterDir, k))
		for i := 0; i < len(v); i++ {
			//u:=fmt.Sprintf("%s%s", MasterUrl, v[i])
			tmp_file := fmt.Sprintf("%d.jpg", i)

			if exist(filepath.Join(m_dir, tmp_file)) {
				if exist_cover { //存在且常量定义为覆盖，覆盖
					downfile(v[i], filepath.Join(m_dir, tmp_file))
				} else {
					log.Printf("file:%s 跳过", filepath.Join(m_dir, tmp_file))
					continue
				}
			}else {
				downfile(v[i], filepath.Join(m_dir, tmp_file))
			}
		}
	}
	for k, v := range dir_torrenturls {
		m_dir := filepath.Join(filepath.Join(MasterDir, k))
		for i := 0; i < len(v); i++ {
			//u:=fmt.Sprintf("%s%s", MasterUrl, v[i])
			tmp_file := fmt.Sprintf("%d.torrent", i)
			if exist(filepath.Join(m_dir, tmp_file)) {
				if exist_cover { //存在且常量定义为覆盖，覆盖
					downfile(v[i], filepath.Join(m_dir, tmp_file))
				} else {
					log.Printf("file:%s 跳过", filepath.Join(m_dir, tmp_file))
					continue
				}
			}else{
				downfile(v[i], filepath.Join(m_dir, tmp_file))
			}
		}
	}
	tmp_chan <- struct{}{}
}
func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	go UnLinks()
	ForumGet()

	w.Wait()

	downloadall()
	<-tmp_chan
	//for k, v := range dir_imageurls {
	//	fmt.Println(k, len(v))
	//}
	//fmt.Println(len(dir_imageurls))
	//for k, v := range dir_torrenturls {
	//	fmt.Println(k, len(v))
	//}
	//fmt.Println(len(dir_torrenturls))
}
