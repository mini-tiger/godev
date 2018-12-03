package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"strconv"
	"os"
)

//access token interface
type ati struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
}

//upload file response
type uploadresp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Media_id string `json:"media_id"`
	Type     string `json:"type"`
}

const (
	corpid     = "wx0a5e0ea42d34d2e1"
	corpsecret = "1YwqhYfsUKDoRYYJWjMRcL_T0G2ZSPCPmQ2RwhDTzcw"
	Agentid    = 1             //报警监控 应用ID
	filename   = "c:\\111.png" //上传图片路径
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
	log.Println("request access_token")
	accesstoken, err := get_accesstoken()
	if err != nil {
		log.Fatalln(accesstoken, err)
	}

	//log.Println("sendmessage , message: test abc")
	//sendmessage(accesstoken, "test abc","text")

	log.Printf("uploadfile : %s\n", filename)

	mediaid := uploadmedia(accesstoken, filename)
	sendmessage(accesstoken, mediaid, "image")
}

func uploadmedia(access, filename string) (mediaid string) {
	urladdress := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s",
		access, "image")

	client := &http.Client{}
	//log.Println(url, "post")

	//方法一
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	temp := bytes.NewBuffer(file)

	fileinfo, err := os.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}
	request, err := http.NewRequest("POST", urladdress, temp)
	if err != nil {
		log.Fatalln(err)
	}

	//方法二
	//temp, err := os.Open(filename)
	//fileinfo, err := temp.Stat()
	//b:=make([]byte,fileinfo.Size())
	//_,err =temp.Read(b)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//request, err := http.NewRequest("POST", urladdress, bytes.NewReader(b))
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//
	//request.Header.Set("Content-type", fmt.Sprintf("multipart/form-data; boundary=-------------------------%d", time.Now().UnixNano())) //不是必须
	request.Header.Set("Content-Disposition", fmt.Sprintf("form-data; name=media;filename=%s; filelength=%s",
		fileinfo.Name(), strconv.FormatInt(fileinfo.Size(), 10)))
	request.Header.Set("Content-Type", " application/octet-stream")

	response, err1 := client.Do(request)
	if err1 != nil {
		log.Fatalln(err1)
	}

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	var upr uploadresp
	err = json.Unmarshal(body, &upr)
	if err != nil || upr.Errcode != 0 {
		log.Fatalln(err, upr)
	}
	//fmt.Printf("%+v\n",upr)
	return upr.Media_id
}

type msg_template struct {
	Touser  string            `json:"touser"` //json 项要大写字母开头
	Toparty string            `json:"toparty"`
	Totag   string            `json:"totag"`
	Msgtype string            `json:"msgtype"`
	Agentid string            `json:"agentid"`
	Text    map[string]string `json:"text"`
	Image   map[string]string `json:"image"`
}

func sendmessage(args ...string) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s",
		args[0])

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
	req := make(map[string]interface{})
	req["toparty"] = "2"
	req["totag"] = ""
	req["agentid"] = "1"
	req["touser"] = ""
	switch args[2] {
	case "text":
		req["msgtype"] = "text"
		req["text"] = map[string]string{"content": args[1]}
	case "image":
		req["msgtype"] = "image"
		req["image"] = map[string]string{"media_id": args[1]}
	default:
		log.Fatalln("switch no match")
	}

	jr, err := json.Marshal(req)
	fmt.Println(string(jr))
	if err != nil {
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
