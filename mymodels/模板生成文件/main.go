package main

import (
	"github.com/satori/go.uuid"
	"log"
	"fmt"
	"text/template"
	"os"
)

type Variables struct {
	UUID string
	biz int
}

func createConf(bizid int,filename string)  {
	u,err:=uuid.NewV4()
	if err!=nil{
		log.Println(err)
	}

	t, err := template.ParseFiles("C:\\work\\go-dev\\src\\godev\\mymodels\\模板生成文件\\agent.json") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755) //读写模式
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	p := Variables{UUID:u.String(),biz:bizid}
	//if err := t.Execute(os.Stdout, p); err != nil {
	//	fmt.Println("There was an error:", err.Error())
	//}
	if err := t.Execute(f, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}



	}
func main()  {
	createConf(3,"c:\\cfg.json")

}