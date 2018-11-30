package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	//    "os"
	"net/url"
	"time"
)

func main() {
	//Init jar
	j, _ := cookiejar.New(nil)
	// Create client
	client := &http.Client{Jar: j}
	// Create request
	req, err := http.NewRequest("GET", "http://i.baidu.com/", nil)
	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failure : ", err)
	}
	//开始修改缓存jar里面的值
	var clist []*http.Cookie
	clist = append(clist, &http.Cookie{
		Name:    "BDUSS",
		Domain:  ".baidu.com",
		Path:    "/",
		Value:   "VLLVVIYVR-TUhqenRtSDk3bm1aU1lXZzRQNE04T3BONGp5UFhDUFlVMTRpQ2RjQVFBQUFBJCQAAAAAAAAAAAEAAACl6JECd3MwMDMxMDk3NgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHj7~1t4-~9bM1",
		Expires: time.Now().AddDate(1, 0, 0),
	})
	urlX, _ := url.Parse("http://i.baidu.com/")
	j.SetCookies(urlX, clist)

	fmt.Printf("Jar cookie : %v", j.Cookies(urlX))
	// Fetch Request
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
	fmt.Printf("response Cookies :%v", resp.Cookies())

}
