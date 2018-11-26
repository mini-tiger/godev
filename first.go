package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:1080")
	}

	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport, Timeout: 60 * time.Second}
	resp, err := client.Get("http://www.google.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode == 200 {
		//fmt.Println(resp.Body)
		s, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(s))
	} else {
		fmt.Println("Fail")
	}
}
