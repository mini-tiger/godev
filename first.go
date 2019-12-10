package main

import (
	"fmt"
	"strconv"
	"time"
)

func formatTime(toBeCharge string) string {
	//toBeCharge := "09/04/2019 11:03:16"
	timeLayout := "01/02/2006 15:04:05" // 	月/日/年

	//转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	//fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)
	if sr<0{
		timeLayout = "02/01/2006 15:04:05" // 	月/日/年
		theTime, _ = time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
		sr = theTime.Unix()
	}
	return strconv.FormatInt(sr, 10)
}


func main() {

	//a:=make([]string,3)
	b:=make([]string,3)
	b[1]="1"
	b[0]="0"
	b[2]="2"
	index:=1
	b=append(b[:index],b[index+1:]...)
	//fmt.Println(a)
	fmt.Println(b,len(b))


}