package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

//access token interface
type ati struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
}

const (
	corpid     = "wx0a5e0ea42d34d2e1"
	corpsecret = "1YwqhYfsUKDoRYYJWjMRcCRPwk-oaEFNNOl7MhEX120"
	Agentid    = 1 //报警监控
)

func get_accesstoken() (accesstoken string, err error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s",
		corpid, corpsecret)

	log.Println(url, "get", )

	req, err := http.Get(url)
	// req.Header.Set("X-Custom-Header", "myvalue")
	//req.Header.Set("Content-Type", "application/json")
	if req.StatusCode != 200 {
		return "", err
	}
	s, _ := ioutil.ReadAll(req.Body)
	//ss:=bytes.NewReader(s)
	//for{
	//	tmp,err:=ss.ReadByte()
	//	if err!=nil || err==io.EOF{
	//		fmt.Println()
	//		break
	//	}
	//	fmt.Print(string(tmp))
	//}
	//fmt.Println("response Body:", req.Body)
	//results:=gjson.Get(string(s),"errcode").Exists() //是否存在
	//fmt.Println(results)
	var pp ati

	error := json.Unmarshal(s, &pp) //解析
	if error != nil {
		return "", error
	}
	//fmt.Printf("%+v\n", pp)
	return pp.Access_token, nil
}

func main() {
	accesstoken, err := get_accesstoken()
	if err != nil {
		log.Fatalln(accesstoken, err)
	}
	sendmessage(accesstoken, "test abc")
}

type msg_template struct {
	Touser  string `json:"touser"` //json 项要大写字母开头
	Toparty string `json:"toparty"`
	Totag   string`json:"totag"`
	Msgtype string`json:"msgtype"`
	Agentid string`json:"agentid"`
	Text    map[string]string`json:"text"`
}

func sendmessage(access, msg string) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s",
		access)

	log.Println(url, "post")
	//方法一
	//req:=`{"touser": "",
	//	"toparty": "2",
	//	"totag": "",
	//	"msgtype": "text",
	//	"agentid": "1",
	//	"text": {
	//		"content":"123456"}
	//}`

	//方法二
	//req :=&msg_template{}
	//req.Touser=""
	//req.Totag=""
	//req.Agentid="1"
	//req.Toparty="2"
	//req.Msgtype="text"
	//req.Text=map[string]string{"content":msg}
	//jr,err:=json.Marshal(req)
	//fmt.Println(string(jr))
	//if err!=nil{
	//	log.Fatalln(err)
	//}

	//方法三
	req:=make(map[string]interface{})
	req["toparty"]="2"
	req["totag"]=""
	req["agentid"]="1"
	req["msgtype"]="text"
	req["touser"]=""
	req["text"]=map[string]string{"content":msg}
	jr,err:=json.Marshal(req)
	fmt.Println(string(jr))
	if err!=nil{
		log.Fatalln(err)
	}

	client := &http.Client{}
	req_new := bytes.NewBuffer(jr) //方法二，三
	//方法一  req_new := bytes.NewBuffer([]byte(req))

	request, _ := http.NewRequest("POST", url, req_new)
	request.Header.Set("Content-type", "application/json")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}

}
