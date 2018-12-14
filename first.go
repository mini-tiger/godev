package main

import (
	"time"
	"fmt"
	"compress/zlib"
)

type Test struct {
	S string
}

func main() {
	timeLayout := "2006-01-02 15:04:05"
	s := time.Now().Unix()
	s = 1544600603
	fmt.Println(time.Unix(s, 0).Format(timeLayout))
	//time.Sleep(time.Duration(5)*time.Second)
	ss := time.Now().Unix()
	d,h:=getday(s, ss)
	fmt.Printf("%d天%d小时",d,h)
}
func getday(sinter int64) (bb bool,day,interval int64) { // 是否不足1天, 有几天, 减去天数后的时间差
	day =
	if day < 0 {
		b
	}
	return

}

func getBasic(sinter, tt int64) (zhengshu int64,b bool) { //除完的整数,是否整数大于0
	zhengshu = sinter / tt
	if zhengshu >0{
		b=true
		return
	}
	return
}
