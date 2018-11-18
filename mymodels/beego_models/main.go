package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"godev/mymodels/beego_models/models"
	"log"
	"time"
)

func main() {
	// todo 注册DB
	models.RegitDB("falcon:123456@tcp(192.168.1.104:3306)/test?charset=utf8&loc=Asia%2FShanghai")
	// 数据库别名
	name := "default"
	// drop table 后再建表
	force := false
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	orm.Debug = true
	CRUD("v2", "agent", "98dd507a-e2d7-496b-8332-be61a9732a5")

}
func CRUD(ver, appname, uuid string) { // 查询任务表 ,判断是否执行
	currt := time.Now().Unix()

	b, app, _ := SelectAppMission(appname, uuid)
	if b {
		if app.Version != ver && app.Count <= 3 && currt >= app.InstallTime { //版本不一样,重复不到3次，时间已过, 就下发任务
			fmt.Println("send agent")
			_, err := UAppSend(ver, appname, uuid, currt)
			if err != nil {
				log.Printf("uappsend err:%s\n", err)
				return // ubs 下发版本, FTP 信息给agent
			}
		} else {
			fmt.Println("no send agent")
			return
		}

	}
}

func UAppSend(ver, appname, uuid string, currt int64) (tmp models.Mission, err error) {

	o1 := orm.NewOrm()
	o1.Using("default")

	tmp = models.Mission{Version: ver, AppName: appname, UUID: uuid}
	//err = o1.Read(&tmp)

	err = o1.Read(&tmp)
	if err != nil {
		return
	}
	tmp.Count = tmp.Count + 1
	tmp.LastTime = currt
	num, err := o1.Update(&tmp)
	if err != nil || num == 0 {
		return
	}
	return
}

func BaseSelect(m interface{}) (tmp1 interface{}, err error) {
	o1 := orm.NewOrm()
	o1.Using("default")

	switch m {
	case m.(models.Mission):
		tmp := m.(models.Mission)
		//fmt.Printf("%T\n",tmp)
		//tmp = models.Mission{AppName: "agent", UUID: "98dd507a-e2d7-496b-8332-be61a9732a5"}
		err = o1.Read(&tmp)
		tmp1 = tmp
	}
	return tmp1, err

}
func SelectAppMission(appname, uuid string) (b bool, tmp models.Mission, err error) {
	//o1 := orm.NewOrm()
	//o1.Using("default")
	//fmt.Println(o1.Driver().Name())                // 数据库别名，default  切换不同数据库用
	//fmt.Println(o1.Driver().Type() == orm.DRMySQL) // 数据库类型
	//var ti interface{}
	tmp = models.Mission{AppName: appname, UUID: uuid}
	//err = o1.Read(&tmp)
	result, err := BaseSelect(tmp)
	tmp = result.(models.Mission)

	switch err {
	case orm.ErrNoRows:
		return false, tmp, err
	case orm.ErrMissPK:
		return false, tmp, err
	case nil:
		return true, tmp, nil
	default:
		log.Printf("select app name:%s,uuid:%s,err:%s\n", appname, uuid, err)
		return false, tmp, err
	}

	return

}
