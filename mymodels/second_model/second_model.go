package main

import (
	"encoding/json"
	"fmt"
	"godev/g"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

//定义配置文件解析后的结构
type MongoConfig struct {
	MongoAddr      string
	MongoPoolLimit int
	MongoDb        string
	MongoCol       string
}

type Config struct {
	Addr  string
	Port  int `json:",string"`
	Mongo MongoConfig
}

var version string
var tt string

func main() {
	Test_var := g.Test_var
	Test_var = 1

	fmt.Println(version, Test_var, tt)
	fmt.Printf("%T\n", Test_var)
	dir, err := filepath.Abs(filepath.Dir(filepath.Dir(os.Args[0])))
	if err != nil {
		return
	}

	JsonParse := NewJsonStruct()
	p := path.Join(dir, "config", "config.json")
	fmt.Println(p)
	v := Config{}
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	JsonParse.Load(p, &v)
	fmt.Println(v.Addr)
	fmt.Printf("%T,%+v\n", v.Port, v.Port)
	fmt.Println(v.Mongo.MongoDb)
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
