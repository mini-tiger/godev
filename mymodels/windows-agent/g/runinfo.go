package g

import (
	"time"

	"tjtools/utils"
	"runtime"
)

var startRunTime int64
var timeLayout string

func init() {
	startRunTime = time.Now().Unix()
	timeLayout = "2006-01-02 15:04:05"
}

func RunStatus() {
	go TimeStatus()
	//go MemStatus()
}

func MemStatus()  {
	for {
		var m runtime.MemStats

		runtime.ReadMemStats(&m)

		logger.Printf("程序使用内存:%+v\n",m)
		time.Sleep(time.Duration(60) * time.Second)
	}
}

func TimeStatus()  {
	for {

		d, h, m, s := utils.GetTime(time.Now().Unix() - startRunTime)
		logger.Printf("开始运行时间 %s,已经运行%d天%d小时%d分钟%d秒", time.Unix(startRunTime, 0).Format(timeLayout), d, h, m, s)
		time.Sleep(time.Duration(3600) * time.Second)
	}
}

//func getTime(b, e int64) (day, hour int64) {
//	day = getDay(b, e)
//	if day <= 0 {
//		hour = getHour(b, e)
//		return
//	} else {
//		b = b + day*86400
//		hour = getHour(b, e)
//		return
//	}
//	return
//
//}
//
//func getHour(begin, end int64) (h int64) {
//	if end-begin <= 0 {
//		return
//	}
//	return (end - begin) / 3600
//}
//func getDay(begin, end int64) (h int64) {
//	if end-begin <= 0 {
//		return
//	}
//	return (end - begin) / 86400
//}

//func GetTime(sinter int64) (d, h, m, s int64) {
//	if day, interval := getBasic(sinter, 86400); day > 0 {
//		d = day
//		h, m = oneDay(interval)
//	} else {
//		h, m = oneDay(interval)
//	}
//	s = sinter - (d * 86400) - (h * 3600) - (m * 60)
//	return
//}
//
//func oneDay(interval int64) (h, m int64) { // Ò»ÌìÄÚµÄÐ¡Ê±·ÖÖÓ
//	h = interval / 3600
//	if h == 0 {
//		m = interval / 60
//		return
//	} else {
//		interval = interval - (3600 * h)
//		m = interval / 60
//		return
//	}
//	return
//}
//
//func getBasic(sinter, tt int64) (num, interval int64) { // ÊÇ·ñ²»×ã1Ìì, ÓÐ¼¸Ìì, ¼õÈ¥ÌìÊýºóµÄÊ±¼ä²î
//	if num = sinter / tt; num > 0 {
//		return num, sinter - (num * tt)
//	} else {
//		return 0, sinter
//	}
//	return
//}
