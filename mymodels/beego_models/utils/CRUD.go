package utils

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"godev/mymodels/beego_models/models"
	"time"
)

func GetApps() {
	o1 := orm.NewOrm()
	o1.Using("default")
	//fmt.Println(o1.Driver().Name())                // 数据库别名，default  切换不同数据库用

	//result, err := BaseSelect(tmp)
	//tmp = result.(models.Mission)
	ct := time.Now().Unix()
	fmt.Println(ct)
	var tmp []orm.Params
	m := new(models.MissionDetail)
	num, _ := o1.QueryTable(m).Filter("uuid", "b38ae270-4e5a-4a2f-9ce8-9c49a6c171a1").Filter("installtime__lte", ct).Exclude("appname", "agent").Filter("status__lte", 0).Values(&tmp, "appname", "installpath", "version", "ftpinfo")
	if num == 0 {
		return
	}
	//for _,v:=range tmp{
	//	for _,vv:=range []string{"vsftpd"}{
	//		if ap,_:=v["AppName"];ap == "vsftpd" {
	//			if v["Version"] != "v1"{
	//				fmt.Println("vsftp add")
	//			}
	//		}
	//	}
	//
	//}
	//fmt.Println(num,err)
	fmt.Println(tmp)
}

func sql()  {
	o1 := orm.NewOrm()
	o1.Using("default")
	res, err := o1.Raw("UPDATE missiondetail SET status=?,success=?,lasttime=?,stdout=?,stderr=? where id=?",
		1,1,1551152139000,"","",89).Exec()
	if err == nil {
		num, err := res.RowsAffected()
		fmt.Println(err)
		fmt.Println("mysql row affected nums: ", num)
	}
	fmt.Println("111",err)
}

func CRUD(ver, appname, uuid string) { // 查询任务表 ,判断是否执行
	//GetApps()
	sql()

	//currt := time.Now().Unix()
	//	//
	//	//b, app, _ := SelectAppMission(appname, uuid)
	//	////fmt.Println(currt,b,app)
	//	//if b {
	//	//	if app.Version != ver && app.Count <= 3 && currt >= app.InstallTime { //版本不一样,重复不到3次，时间已过, 就下发任务
	//	//		fmt.Println("send agent")
	//	//		_, err := UAppSend(app, currt)
	//	//		if err != nil {
	//	//			log.Printf("uappsend err:%s\n", err)
	//	//			return // ubs 下发版本, FTP 信息给agent
	//	//		}
	//	//	} else {
	//	//		fmt.Println("no send agent")
	//	//		return
	//	//	}
	//	//
	//	//}
}

//func UAppSend(app *models.Mission, currt int64) (tmp models.Mission, err error) {
//
//	o1 := orm.NewOrm()
//	o1.Using("default")
//
//	//tmp = models.Mission{Version: ver, AppName: appname, UUID: uuid}
//	////err = o1.Read(&tmp)
//	//
//	//err = o1.Read(&tmp)
//	//if err != nil {
//	//	return
//	//}
//	//m:=new(models.Mission)
//	fmt.Println(app)
//	m := models.Mission{Id: app.Id}
//
//	m.Count = app.Count + 1
//	m.LastTime = currt
//	fmt.Println(m)
//	num, err := o1.Update(&m, "count")
//	fmt.Println("ddd", num)
//	if err != nil {
//		return
//	}
//	if num == 0 {
//		err = errors.New("update num =0 ,Fail")
//	}
//	return
//}

//func BaseSelect(m interface{}) (tmp1 interface{}, err error) {
//	o1 := orm.NewOrm()
//
//	o1.Using("default")
//	k := reflect.TypeOf(m).String()
//	//fmt.Println(a.String())
//	//fmt.Printf("%s\n",k)
//
//	switch k {
//	case "models.Mission":
//		m := new(models.Mission)
//		o1.QueryTable(m)
//
//		//fmt.Println(1)
//		//tmp := m.(models.Mission)
//		////fmt.Printf("%T\n",tmp)
//		////tmp = models.Mission{AppName: "agent", UUID: "98dd507a-e2d7-496b-8332-be61a9732a5"}
//		//fmt.Println(tmp)
//		//err = o1.Read(&tmp)
//		//tmp1 = tmp
//	}
//	fmt.Println(tmp1)
//	return tmp1, err
//
//}
//func SelectAppMission(appname, uuid string) (b bool, m *models.Mission, err error) {
//	o1 := orm.NewOrm()
//	o1.Using("default")
//	//fmt.Println(o1.Driver().Name())                // 数据库别名，default  切换不同数据库用
//	//fmt.Println(o1.Driver().Type() == orm.DRMySQL) // 数据库类型
//	//var ti interface{}
//	//tmp = models.Mission{AppName: appname, UUID: uuid}
//	//err = o1.Read(&tmp)
//	fmt.Println(appname, uuid)
//	//result, err := BaseSelect(tmp)
//	//tmp = result.(models.Mission)
//	m = new(models.Mission)
//	qs := o1.QueryTable(m).Filter("appname", appname).Filter("uuid", uuid)
//
//	//cond := orm.NewCondition()
//	//cond1 := cond.And("AppName", appname).And("UUID",uuid)
//	//qs.SetCond(cond1)
//	err = qs.One(m)
//	//fmt.Println(m)
//
//	//fmt.Println(result)
//	//fmt.Println(tmp)
//	switch err {
//	case orm.ErrNoRows:
//		fmt.Println("nowrow")
//		return false, m, err
//	case orm.ErrMissPK:
//		fmt.Println("misspk")
//		return false, m, err
//	case nil:
//		return true, m, nil
//	default:
//		log.Printf("select app name:%s,uuid:%s,err:%s\n", appname, uuid, err)
//		return false, m, err
//	}
//
//	return
//
//}
