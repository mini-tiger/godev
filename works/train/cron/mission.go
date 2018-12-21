package cron

import (
	"time"
	"godev/works/train/g"
	"godev/works/train/db"
)

func TrainCrond()  {
	for{
		go trainBussins()
		time.Sleep(time.Duration(g.Config().TrainInterval)*time.Second)
	}

}

func trainBussins()  {
	sql:="select train_serial from tf_op_train where rownum < 2"
	db.GetRows(sql)


}