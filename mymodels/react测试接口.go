package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Resp struct {
	Status int `json:"status"`
	StatusText string `json:"statusText"`
	Data []string `json:"data"`
}

func react1(w http.ResponseWriter, r *http.Request) { // 上传form数据
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println("Form: ", r.Form)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println(r.Form["a"])
	fmt.Println(r.Form["b"])
	for k, v := range r.Form {
		fmt.Println(k, "=>", v, strings.Join(v, "-"))
	}
	var s = make([]string,0)
	s = append(s,[]string{"a","b"}...)
	//fmt.Fprint(w, s)
	var rs Resp
	rs.Data = s
	rs.Status = 200
	rs.StatusText = "ok"

	jb,err:=json.Marshal(rs)
	if err!=nil{
		fmt.Println("err:",err)
	}

	fmt.Fprint(w, string(jb))

}


func react2(w http.ResponseWriter, r *http.Request) { // 上传json数据
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	//w.Header().Set("content-type", "application/json")             //返回数据格式是json
	body, _ := ioutil.ReadAll(r.Body)
	//    r.Body.Close()
	body_str := string(body)
	fmt.Println(body_str)
	//    fmt.Fprint(w, body_str)
	var user User
	//    user.Name = "aaa"
	//    user.Age = 99
	//    if bs, err := json.Marshal(user); err == nil {
	//        fmt.Println(string(bs))
	//    } else {
	//        fmt.Println(err)
	//    }

	if err := json.Unmarshal(body, &user); err == nil {
		fmt.Println(user)
		user.Age += 100
		fmt.Println(user)
		ret, _ := json.Marshal(user)
		fmt.Fprint(w, string(ret))
	} else {
		fmt.Println("err:",err)
	}
}

func main() {
	//http.HandleFunc("/", index)
	http.HandleFunc("/react2/", react2)
	http.HandleFunc("/react1/", react1)

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
