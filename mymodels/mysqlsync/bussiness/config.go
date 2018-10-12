package bussiness

import (
	"log"
	"io/ioutil"
	"strings"
	"encoding/json"
)

type SrcDb struct {
	Ip      string   `json:"ip"`
	DbUser  string   `json:"dbUser"`
	DbPass  string   `json:"dbPass"`
	DbName  string   `json:"dbName"`
	SortCol string   `json:"sortCol"`
	NeedCol []string `json:"needCol"`
}

type DstDb struct {
	Ip      string   `json:"ip"`
	DbUser  string   `json:"dbUser"`
	DbPass  string   `json:"dbPass"`
	DbName  string   `json:"dbName"`
	SortCol string   `json:"sortCol"`
	NeedCol []string `json:"needCol"`
}

type GlobalConfig struct {
	LogLevel   string `json:"logLevel"`
	Logfile    string `json:"logfile"`
	LogMaxDays int    `json:"logMaxDays"`
	Interval   int    `json:"interval"`
	SrcDb      *SrcDb `json:"srcDb"`
	DstDb      *DstDb `json:"dstDb"`

	IgnoreMetrics map[string]bool `json:"ignore"`
}

var Config *GlobalConfig
var Cfg string
func InitConfig() {
	//if cfg == "" {
	//	log.Fatalln("没有传递配置文件参数")
	//}

	b, err := ioutil.ReadFile(Cfg)
	if err != nil {
		log.Fatalf("Configfile: [%s], Load err:%s", Cfg, err)
	}
	configContent := strings.TrimSpace(string(b))

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", Cfg, "fail:", err)
	}
	Config = &c
}
