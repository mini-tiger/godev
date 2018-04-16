package main

import (
	"os"
	"strings"
	"text/template" //模板类似 jinja  if range for 等查看徇私文档
	//template.ParseFiles 文件模板方法
)

type Px struct {
	Name  string //首字母大写
	Age   int
	Title string
	// day   time.time
}

func main() {
	sample()
	struct_ex()
	func_ex()
}

func sample() {
	name := "world"
	tmpl, err := template.New("t").Parse("hello, {{ .}} \n") //建立一个模板，内容是"hello, {{.}}"
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, name) //将string与模板合成，变量name的内容会替换掉{{.}}
	//合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}

	//Must  两句可以 用一句
	template.Must(template.New("t").Parse("hello, {{ .}} \n")).Execute(os.Stdout, name)

}

func struct_ex() {

	sweaters := Px{"wool", 17, "abc1"}

	const mb = `------------    
title : {{ .Title | printf "%.64s"}} 
age : {{ .Age }} 
name : {{ .Name}}

`
	/* ` 号 不用\n换行 */
	tmpl, err := template.New("ttt").Parse(mb) //建立一个模板
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters) //将struct与模板合成，合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}

	//方法二
	mb1 := "------------\ntitle : {{ .Title | printf `%.64s`}} \nage : {{ .Age }} \nname : {{ .Name}}\n"
	template.Must(template.New("t").Parse(mb1)).Execute(os.Stdout, sweaters)
}

func upper(s string) string {
	return strings.ToTitle(s)
}

func func_ex() {
	// type Px struct {
	// 	Name  string //首字母大写
	// 	Age   int
	// 	Title string
	// 	// day   time.time
	// }

	sp := Px{"func_ex", 17, "bbb"}

	mb2 := "------------\ntitle : {{ .Title | printf `%.64s`}} \nage : {{ .Age }} \nname : {{ .Name |totitle }}\n"

	tmpl, err := template.New("tt").Funcs(template.FuncMap{"totitle": upper}).Parse(mb2)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, sp)

}
