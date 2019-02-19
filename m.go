package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	//sendjson() //发送并接收JSON
	//sendForm()
	get()
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Hostgroup struct {
	Id          int    `json:"id"`
	Grp_name    string `json:"grp_name"`
	Create_user string `json:"create_user"`
}

func get() {
	req := httplib.NewBeegoRequest("http://10.240.81.85:8080/api/v1/hostgroup", "GET")
	req.SetTimeout(time.Duration(10)*time.Second, time.Duration(10)*time.Second)

	resp, err := req.Response()
	//fmt.Println(resp,err)

	if err != nil {
		log.Println(err)
	}
	if resp.StatusCode == 200 {
		fmt.Println(200)
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var respdata []Hostgroup
		err = json.Unmarshal(respBytes, &respdata)
		if err != nil {
			log.Printf("json err:%s\n", err)
		}
		fmt.Println(respdata)
	} else {
		log.Println(resp.Body)
	}

}

func sendForm() {
	req := httplib.NewBeegoRequest("http://127.0.0.1:8080/", "GET")
	//r,err:=httplib.Get("http://www.163.com/").Debug(true).Response()
	//fmt.Println(req.Response())
	//fmt.Println(r,err)
	req.Header("Accept-Encoding", "gzip,deflate,sdch")
	req.Header("Connection", "keep-alive")
	req.Header("Content-type", "application/json")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")

	req.SetTimeout(time.Duration(10)*time.Second, time.Duration(10)*time.Second)
	req.Param("username", "astaxie")
	req.Param("password", "123456")
	//resp,err:=req.Response()
	//if err!=nil{
	//	log.Println(err)
	//}
	//if resp.StatusCode ==200{
	//	resp.Body
	//}
	s, err := req.String()
	if err == nil {
		fmt.Println(s)
	}
}

func sendjson() {
	req := httplib.NewBeegoRequest("http://127.0.0.1:8080/test/", "get")
	//r,err:=httplib.Get("http://www.163.com/").Debug(true).Response()
	//fmt.Println(req.Response())
	//fmt.Println(r,err)
	req.Header("Accept-Encoding", "gzip,deflate,sdch")
	req.Header("Connection", "keep-alive")
	req.Header("Content-type", "application/json")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")

	req.SetTimeout(time.Duration(10)*time.Second, time.Duration(10)*time.Second)
	u := User{"abc", 11}

	req.JSONBody(u) //todo 直接发送json

	//方法一 todo 返回解析body内容
	//resp,err:=req.Response()
	////fmt.Println(resp,err)
	//
	//if err!=nil{
	//	log.Println(err)
	//}
	//if resp.StatusCode == 200{
	//	fmt.Println(200)
	//	respBytes, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	var respdata User
	//	err=json.Unmarshal(respBytes,&respdata)
	//	if err!=nil{
	//		log.Printf("json err:%s\n",err)
	//	}
	//	fmt.Println(respdata)
	//}else{
	//	log.Println(resp.Body)
	//}

	//方法二 todo 返回的结果直接 JSON解析,不能查看 statucode
	var respdata User
	err := req.ToJSON(&respdata)
	if err != nil {
		log.Println("resp err:%s\n", err)
	}
	fmt.Println(respdata)
}
