package g

import (
	"strconv"
	"strings"
	"time"
)

func FormatTime(toBeCharge string) string {
	//toBeCharge := "09/04/2019 11:03:16"
	// 如果不是时间格式
	if _, err := strconv.Atoi(strings.Split(toBeCharge, "/")[0]); err != nil {
		return toBeCharge
	}

	// 待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006/01/02 15:04:05" // 中英文 时间格式不一样
	if ii, _ := strconv.Atoi(strings.Split(toBeCharge, "/")[0]); ii <= 12 {
		timeLayout = "01/02/2006 15:04:05" //转化所需模板
	}

	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	//fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	//fmt.Println(sr)
	return strconv.FormatInt(sr, 10) + "000"
}
