package cron

import (
	"time"
	"godev/works/train/g"
	"fmt"
	"strconv"
)

const Lasttime = "lasttime"

var lastTime *int = new(int) // todo 最近火车进入时间戳,可以从redis读取后 写入
//var tt *time.Ticker
//var tr *time.Timer // 一次性
var cTrain chan map[string]string = make(chan map[string]string, 5)

func TrainCrond() {
	//for{
	LoadLasttime()
	if *lastTime == 0 {
		diffTrainTime()
	} else {
		for {
			results, err := getTrainNew()
			if err != nil {
				time.Sleep(time.Duration(5) * time.Second)
				continue
			}
			firstRun(results[0])
			break
		}
	}

	//go WaitMission() // loop wait mission
	go GetTrainPic()

}
func LoadLasttime() {
	if g.Redis.StringExists(Lasttime) {
		val, _ := g.Redis.StringGet(Lasttime)
		i, err := strconv.Atoi(val)
		if err != nil {
			i = 0
			lastTime = &i
			return
		}
		lastTime = &i
	} else {
		i := 0
		lastTime = &i
	}
}

func GetTrainPic() {
	for {
		select {
		case result := <-cTrain:
			fmt.Printf("go getpic %+v\n", result)
			go g.GetPicMission(result)
		}
	}
}

func firstRun(result map[string]string) {
	nowUnix := int(time.Now().Unix())
	//nowUnix = 111
	nextTime := *lastTime + g.Config().TrainInterval
	switch {
	case nowUnix <= nextTime: //todo 当前时间小于 库中最新时间戳 加 时间间隔，可能第一次启动在 列车进入期间
		g.Logger().Debug("第一次启动，创建图片获取任务")
		cTrain <- result // todo 创建协程 获取图片任务， 不影响下次 获取最新时间的任务
		go CreateNextTrainMission(nextTime - nowUnix,nowUnix) // 距离下次的间隔秒数
		g.Logger().Debug("NextTime %s", g.GetDateStr(nextTime))

	case nowUnix > nextTime:
		g.Logger().Debug("第一次启动，创建循环任务")
		go CreateLoopMission() // todo 创建循环10s 判断一次是否有新的
	}
}

func CreateLoopMission() { // 循环获取时间任务，由于一次性任务没有取到
	//tt = time.NewTicker(time.Duration(g.Config().LoopInterval) * time.Second)
	go LoopMission(g.Config().LoopInterval)
}

func CreateNextTrainMission(interval ,now int) { // 一次性，下次获取时间任务
	//tr = time.NewTimer(time.Duration(interval) * time.Second)
	time.Sleep(time.Duration(interval) * time.Second)
	go NextTrainMission(now)
}

func LoopMission(interval int) {
	for {
		tn := int(time.Now().Unix())
		if b, resultTrain := diffTrainTime(); b {

			g.Logger().Printf("循环结束\n")
			//fmt.Println("go 启动获取图片")
			cTrain <- resultTrain
			nt := *lastTime + g.Config().TrainInterval - tn
			g.Logger().Printf("启动train任务开始时间%s\n", g.GetDateStr(nt+tn))
			go CreateNextTrainMission(nt,g.GetNow())
			return
		} else {
			g.Logger().Printf("下次循环任务开始时间%s\n", time.Now().Add(time.Duration(interval)*time.Second))
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func NextTrainMission(now int) {

	if b, resultTrain := diffTrainTime(); b { //一次性 获取到新的时间数据
		cTrain <- resultTrain
		nt := *lastTime + g.Config().TrainInterval - now
		g.Logger().Printf("下次train任务开始时间%s\n", g.GetDateStr(nt+now))
		go CreateNextTrainMission(nt,int(time.Now().Unix())) // 下次一次性
	} else { // 没有获取到新的数据，进入循环获取
		g.Logger().Printf("创建循环任务开始时间%s\n", time.Now().Add(time.Duration(g.Config().LoopInterval)))
		go CreateLoopMission()
	}
}

//func WaitMission() {
//
//	for {
//		select {
//		case tc:=<-tt.C:
//			if b, resultTrain := diffTrainTime(); b {
//				tt.Stop()
//				g.Logger().Printf("循环结束\n")
//				//fmt.Println("go 启动获取图片")
//				cTrain <- resultTrain
//				nt := *lastTime + g.Config().TrainInterval - int(tc.Unix())
//				g.Logger().Printf("下次train任务开始时间%s\n", g.GetDateStr(nt+int(tc.Unix())))
//				go CreateNextTrainMission(nt)
//
//			} else {
//				g.Logger().Printf("下次循环任务开始时间%s\n", tc.Add(time.Duration(g.Config().LoopInterval)))
//			}
//		case trc11 := <-tr.C:
//			if b, resultTrain := diffTrainTime(); b { //一次性 获取到新的时间数据
//				cTrain <- resultTrain
//				nt := *lastTime + g.Config().TrainInterval - int(trc11.Unix())
//				g.Logger().Printf("下次train任务开始时间%s\n", g.GetDateStr(nt+int(trc11.Unix())))
//				go CreateNextTrainMission(nt) // 下次一次性
//			} else { // 没有获取到新的数据，进入循环获取
//				g.Logger().Printf("创建循环任务开始时间%s\n", trc11.Add(time.Duration(g.Config().LoopInterval)))
//				go CreateLoopMission()
//			}
//
//		}
//
//	}
//}

func diffTrainTime() (b bool, resultTrain map[string]string) { // 是否应该进行获取图片，库中时间
	resultTrainSlice, err := getTrainNew()
	if err != nil {
		g.Logger().Error("getTrainNew err:%s", err)
		return
	}

	if ti, b := g.Getint(resultTrainSlice[0]["UNIXSTAMP"]); b {
		g.Logger().Debug("获取到库中最新时间%d ,%s", ti, g.GetDateStr(ti))
		switch {
		case *lastTime == 0:
			lastTime = &ti
			g.Redis.StringSet(Lasttime, *lastTime)
			firstRun(resultTrainSlice[0])
		case *lastTime == ti:
			return false, resultTrain
		case *lastTime < ti:
			lastTime = &ti
			g.Redis.StringSet(Lasttime, *lastTime)
			return true, resultTrainSlice[0]
		default: // 库中数据除了第一次以外，不应该比 lasttime 小
			lastTime = &ti
			g.Redis.StringSet(Lasttime, *lastTime)
			return false, resultTrain
		}
	} else { // 获取到的时间格式如果不对
		g.Logger().Error("获取到的时间格式错误 %d 秒后再次运行", g.Config().TrainInterval)
		CreateNextTrainMission(g.Config().TrainInterval,g.GetNow())
		return false, resultTrain
	}
	return
}

func getTrainNew() (results []map[string]string, err error) {

	sql := "select train_id,station_id,pass_time,vehicle_number,index_id,(pass_time - TO_DATE('1970-01-01 08:00:00', 'YYYY-MM-DD HH24:mi:ss')) * 86400 as unixstamp " +
		"from tf_op_train  where rownum = 1 order by pass_time desc "

	g.Logger().Debug("sql :%s", sql)

	results, err = g.Engine.QueryString(sql)

	if len(results) == 0 {
		results[0]["UNIXSTMAP"] = "0"
		return
	}
	//
	//if err != nil {
	//	g.Logger().Error("Exec sql:%s,err:%s\n", sql, err)
	//	return
	//}
	return
	//fmt.Printf("%T,%+v\n",results,results)
}
