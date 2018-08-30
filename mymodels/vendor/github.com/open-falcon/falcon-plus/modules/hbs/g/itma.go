package g

import (
	"encoding/json"
	"github.com/toolkits/file"
	"github.com/open-falcon/falcon-plus/common/model"
	"log"
	"sync"
)

type EnvGridConfig struct {
	Debug     		bool        						`json:"debug"`
	ConfigInterval  int      							`json:"configInterval"`
	DataInterval	int 								`json:"dataInterval"`
	Jeesite			string 								`json:"jeesite"`
	CmdbUrl			string 								`json:"cmdbUrl"`
	Monitor_port	[]model.PortProcessEnv				`json:"monitor_port"`
	Appsys      	[]model.DefaultProcessAppsysWorker 	`json:"appsys"`
}

var (
	EnvGridConfigFile	string
	envGridConfig     	*EnvGridConfig
	envGridConfigLock 	= new(sync.RWMutex)
)

func GetEnvGridConfig() *EnvGridConfig {
	envGridConfigLock.RLock()
	defer envGridConfigLock.RUnlock()
	return envGridConfig
}

func ParseEnvGridConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -g to specify env grid configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("env grid config file:", cfg, "is not existent")
	}

	EnvGridConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read env grid config file:", cfg, "fail:", err)
	}

	var c EnvGridConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse env grid config file:", cfg, "fail:", err)
	}

	envGridConfigLock.Lock()
	defer envGridConfigLock.Unlock()

	envGridConfig = &c

	log.Println("read env grid config file:", cfg, "successfully")
	log.Println("env grid config:", configContent)
}
