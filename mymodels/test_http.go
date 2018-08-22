package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//httpPostForm()
	httpPostjson()
}

func httpPostjson() {
	//登陆用户名
	usrId := "aaaaa"

	//登陆密码
	pwd := "pwd1234"

	//json序列化
	post := "{\"username\":\"" + usrId + "\",\"passwd\":\"" + pwd + "\"}"
	url := "http://127.0.0.1:8001/api/v1/user/login"
	fmt.Println(url, "post", post)

	var jsonStr = []byte(post)
	// fmt.Println("jsonStr", jsonStr)
	// fmt.Println("new_str", bytes.NewBuffer(jsonStr))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func httpPostForm() {

	client := &http.Client{}

	params := strings.NewReader("username=ddd&passwd=abc&age=11")
	req, err := http.NewRequest("POST", "http://127.0.0.1:8001/login", params)

	if err != nil {

		// handle error

	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		// handle error

	}
	fmt.Println(string(body))

}
