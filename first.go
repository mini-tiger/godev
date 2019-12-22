package main

import (
	"gitee.com/taojun319/tjtools/utils"
	"haifei/syncHtml/g"
	"log"

	"time"
)

func ClearMap() {
	tmpLock := make(chan struct{}, 0)
	_=cap(tmpLock)
}

func main() {
	startRunTime:=time.Now().Unix()
	for {
		sinter := time.Now().Unix() - startRunTime
		utils.GetTime(&sinter, &d, &h, &m, &s)
		log.Info("开始运行时间 %s,已经运行%d天%d小时%d分钟%d秒\n", time.Unix(startRunTime, 0).Format(g.TimeLayout), d, h, m, s)
		time.Sleep(time.Duration(900) * time.Second)
	}

}