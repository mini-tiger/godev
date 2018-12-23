package g

import (
	"encoding/json"
	"github.com/toolkits/file"
	"log"
	"sync"
	"github.com/go-xorm/xorm"
)

type redisinfo struct {
	Dsn      string `json:"dsn"`
	Password string `json:"password"`
	PoolSize int    `json:"poolSize"`
}

type GlobalConfig struct {
	Debug         bool       `json:"debug"`
	OracleDsn     string     `json:"oracledsn"`
	Redis         *redisinfo `json:"redis"`
	TrainInterval uint       `json:"traininterval"`
	Logfile       string     `json:"logfile"`
	LogMaxDays    int        `json:"logMaxDays"`
	Daemon        bool       `json:"daemon"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
	lock       = new(sync.RWMutex)
	Engine *xorm.Engine //定义引擎全局变量
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	configLock.Lock()
	defer configLock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
}
