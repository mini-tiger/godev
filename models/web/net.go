package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

//!+main
//  反射  解析json :ch12 search
func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)

	s := &http.Server{
		Addr: ":8080",
		// Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))
	log.Fatal(s.ListenAndServe())
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) } //绑定Strings方法

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price) // %s 调用Strings方法,等价price.String()
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	// log.Println(req.URL.Query())
	item := req.URL.Query().Get("item") //获取GET POST 参数
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price) //写入Response
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Query()) //第一种方法获取GET POS方法 参数
	req.ParseForm()
	log.Println(req.Form, req.FormValue("item"), req.Form.Get("price")) //第二种方法获取GET POS方法 参数
	item := req.URL.Query().Get("item")                                 //获取GET POS方法 参数
	cp := req.URL.Query().Get("price")

	cp1, _ := strconv.ParseFloat(cp, 32) //
	cp2 := dollars(cp1)

	if _, ok := db[item]; ok {
		db[item] = cp2
		fmt.Fprintf(w, "%s: %s\n", item, db[item]) //写入Response
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
