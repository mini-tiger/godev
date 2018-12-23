package cron

import (
	"time"
	"godev/works/train/g"
	"fmt"
	"strconv"
	"sort"
)

func TrainCrond()  {
	for{
		nowUnix := time.Now().Unix()
		go trainBussiness(nowUnix)
		time.Sleep(time.Duration(g.Config().TrainInterval)*time.Second)
	}

}


func trainBussiness(nowUnix int64)  {
	resultTrain,err:=getTrainNew(nowUnix )
	if err!=nil{
		g.Logger().Error("getTrainNew err:%s",err)
	}
	fmt.Println(resultTrain)
	trainLen := len(resultTrain)
	switch  {
	case trainLen == 0:
		g.Logger().Printf("")
	case trainLen == 1:
		fmt.Println(1)
	case trainLen >= 2:
		tmpSlice := make([]int,trainLen)
		for i,v:=range resultTrain{
			if i64,b:=getint64(v["UNIXSTAMP"]);b{
				tmpSlice[i] = i64
			}
		}
		sort.Ints(tmpSlice)
		fmt.Println(tmpSlice)
	}

}
func getint64(s string)  (i64 int,b bool){

	i64, err := strconv.Atoi(s)
	if err!=nil{
		return i64,false
	}else{
		return i64,true
	}

}


func getTrainNew(newtime int64) (results []map[string]string,err error)  {
	sql := fmt.Sprintf("select train_serial,train_id,station_id,pass_time,vehicle_number,index_id," +
		"(pass_time - TO_DATE('1970-01-01 08:00:00', 'YYYY-MM-DD HH24:mi:ss')) * 86400 as unixstamp " +
		"from tf_op_train where (pass_time - TO_DATE('1970-01-01 08:00:00', 'YYYY-MM-DD HH24:mi:ss')) * 86400 > %d",1538016120)
	//sql = "select * from tf_op_train"

	g.Logger().Debug("sql :%s",sql)
	results, err = g.Engine.QueryString(sql)

	if err!=nil{
		g.Logger().Error("Exec sql:%s,err:%s\n",sql,err)
		return
	}
	return
	//fmt.Printf("%T,%+v\n",results,results)
}