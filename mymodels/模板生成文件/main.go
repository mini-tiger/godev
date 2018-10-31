package main

import (
	"github.com/satori/go.uuid"
	"text/template"
	"os"
)

type Variables struct {
	UUID string
	Biz  int
}

func createConf(bizid int, filename string) error {
	t, err := template.ParseFiles("/home/go/src/godev/mymodels/模板生成文件/agent.json.tmp") //parse the html file homepage.html
	if err != nil { // if there is an error
		//log.Print("template parsing error: ", err) // log it
		return err
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755) //读写模式
	if err != nil {
		//log.Println(err)
		return err
		//os.Exit(1)
	}

	u, err := uuid.NewV4()
	if err != nil {
		//log.Println(err)
		return err
	}
	p := Variables{UUID: u.String(), Biz: bizid}
	//if err := t.Execute(os.Stdout, p); err != nil {
	//	fmt.Println("There was an error:", err.Error())
	//}
	if err := t.Execute(f, p); err != nil {
		//fmt.Println("There was an error:", err.Error())
		return err
	}
	return nil

}
func main() {
	createConf(3, "/home/go/src/godev/mymodels/模板生成文件/cfg.json")

}
