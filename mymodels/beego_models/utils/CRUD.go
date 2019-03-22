package utils

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math/rand"
)

//func GetApps() {
//	o1 := orm.NewOrm()
//	o1.Using("default")
//	//fmt.Println(o1.Driver().Name())                // 数据库别名，default  切换不同数据库用
//
//	//result, err := BaseSelect(tmp)
//	//tmp = result.(models.Mission)
//	ct := time.Now().Unix()
//	fmt.Println(ct)
//	var tmp []orm.Params
//	m := new(models.MissionDetail)
//	num, _ := o1.QueryTable(m).Filter("uuid", "b38ae270-4e5a-4a2f-9ce8-9c49a6c171a1").Filter("installtime__lte", ct).Exclude("appname", "agent").Filter("status__lte", 0).Values(&tmp, "appname", "installpath", "version", "ftpinfo")
//	if num == 0 {
//		return
//	}
//	//for _,v:=range tmp{
//	//	for _,vv:=range []string{"vsftpd"}{
//	//		if ap,_:=v["AppName"];ap == "vsftpd" {
//	//			if v["Version"] != "v1"{
//	//				fmt.Println("vsftp add")
//	//			}
//	//		}
//	//	}
//	//
//	//}
//	//fmt.Println(num,err)
//	fmt.Println(tmp)
//}

var years []int = []int{2016, 2017, 2018}
var month []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
var result []string = []string{"良好", "合格"}
var days []int
var itemName = []string{"油饰不良，号码不明", "熔丝容量不符规定", "盘面铭牌缺失，字迹不清", "插接器材不良、焊点不良或脱焊", "移位器安装不良，挤岔功能失灵"}

func return_s(s interface{}) string {
	if s == nil || s.(string) == "" {
		return ""
	}
	return s.(string)
}

//func insert_total_days(m map[string]interface{}) {
//	o := orm.NewOrm()
//	var d models.TotalDayDevice
//
//	Score := rand.Intn(15) - 10
//	d.Score = Score
//	y := years[rand.Intn(len(years))]
//	mm := month[rand.Intn(len(month))]
//	dd := days[rand.Intn(len(days))]
//	d.Year = y
//	d.Month = mm
//	d.Days = dd
//
//	tm2, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%d-%02d-%02d 00:00:01", y, mm, dd))
//	d.DateTime = tm2
//	if Score < 0 {
//		d.Result = "不合格"
//		if Score < -5 {
//			d.ItemClassName = itemName[rand.Intn(len(itemName))] + "," + itemName[rand.Intn(len(itemName))]
//		} else {
//			d.ItemClassName = itemName[rand.Intn(len(itemName))]
//		}
//
//	} else {
//		d.Result = result[rand.Intn(len(result))]
//	}
//	d.DeviceId = m["DeviceId"].(string)
//
//	//fmt.Printf("%T,%v\n",m,m)
//	d.DeviceName = return_s(m["DeviceName"])
//	d.DeviceType = return_s(m["DeviceType"])
//	d.DeviceId = m["DeviceId"].(string)
//
//	d.ORG1ID = m["ORG1_ID"].(string)
//	d.ORG1NAME = m["ORG1_NAME"].(string)
//	d.ORG2ID = m["ORG2_ID"].(string)
//	d.ORG2NAME = m["ORG2_NAME"].(string)
//	d.ORG3ID = m["ORG3_ID"].(string)
//	d.ORG3NAME = m["ORG3_NAME"].(string)
//	d.STAID = m["STA_ID"].(string)
//	d.STANAME = m["STA_NAME"].(string)
//
//	_, err := o.Insert(&d)
//	if err != nil {
//		fmt.Println(err)
//	}
//}

func insert_total_days(maps []orm.Params) {
	o := orm.NewOrm()

	p, err := o.Raw("INSERT INTO `Total_day_Device` " +
		"(`YEAR`, `MONTH`,`Day`, `DeviceType`, `DeviceName`, `DeviceID`, `STA_NAME`, `STA_ID`, `ORG1_ID`, `ORG1_NAME`, `ORG2_ID`, `ORG2_NAME`, `ORG3_ID`," +
		" `ORG3_NAME`, `Total_DEV`, `Total_Appraisal`, `Total_Good`, `Total_Qualified`, `Total_UnQualified`, `Total_WorkBase`, `Total_Work_People`, `datetime`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)").Prepare()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(maps); i++ {
		m := maps[i]
		Total_Dev := 2000 // 每个车站 设备总数固定,
		for _, y := range years { // 遍历年
			for _, mo := range month { // 遍历月
				for _, day := range days {
					//Total_Dev = Total_Dev + (y-2000)*2 + im
					total_good := rand.Intn(30)
					total_Qualified := rand.Intn(30)
					total_unQualified := rand.Intn(30)
					total_Appraisal := total_good + total_Qualified + total_unQualified

					total_workbase := total_Appraisal * 3 //工单数
					Total_Work_People := rand.Intn(100)   // 相关人员
					_, err := p.Exec(y, mo,day, return_s(m["DeviceType"]), return_s(m["DeviceName"]), m["DeviceId"].(string), m["STA_NAME"].(string),
						m["STA_ID"].(string), m["ORG1_ID"].(string), m["ORG1_NAME"].(string), m["ORG2_ID"].(string), m["ORG2_NAME"].(string), m["ORG3_ID"].(string), m["ORG3_NAME"].(string),
						Total_Dev, total_Appraisal, total_good, total_Qualified, total_unQualified, total_workbase, Total_Work_People, fmt.Sprintf("%d-%02d-%02d 00:00:01", y, mo,day))
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
	p.Close() // 别忘记关闭 statement

}

func insert_total_month(maps []orm.Params) {
	o := orm.NewOrm()

	p, err := o.Raw("INSERT INTO `Total_Month_Device` " +
		"(`YEAR`, `MONTH`, `DeviceType`, `DeviceName`, `DeviceID`, `STA_NAME`, `STA_ID`, `ORG1_ID`, `ORG1_NAME`, `ORG2_ID`, `ORG2_NAME`, `ORG3_ID`," +
		" `ORG3_NAME`, `Total_DEV`, `Total_Appraisal`, `Total_Good`, `Total_Qualified`, `Total_UnQualified`, `Total_WorkBase`, `Total_Work_People`, `datetime`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)").Prepare()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(maps); i++ {
		m := maps[i]
		Total_Dev := 2000 // 每个车站 设备总数固定,
		for _, y := range years { // 遍历年
			for _, mo := range month { // 遍历月
				//Total_Dev = Total_Dev + (y-2000)*2 + im
				total_good := rand.Intn(88)
				total_Qualified := rand.Intn(88)
				total_unQualified := rand.Intn(88)
				total_Appraisal := total_good + total_Qualified + total_unQualified


				total_workbase := total_Appraisal * 3 //工单数
				Total_Work_People := rand.Intn(100)   // 相关人员
				_, err := p.Exec(y, mo, return_s(m["DeviceType"]), return_s(m["DeviceName"]), m["DeviceId"].(string), m["STA_NAME"].(string),
					m["STA_ID"].(string), m["ORG1_ID"].(string), m["ORG1_NAME"].(string), m["ORG2_ID"].(string), m["ORG2_NAME"].(string), m["ORG3_ID"].(string), m["ORG3_NAME"].(string),
					Total_Dev, total_Appraisal, total_good, total_Qualified, total_unQualified, total_workbase, Total_Work_People,fmt.Sprintf("%d-%02d-01 00:00:01", y, mo))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	p.Close() // 别忘记关闭 statement

}

func insert_total_season(maps []orm.Params) {
	o := orm.NewOrm()

	p, err := o.Raw("INSERT INTO `Total_Season_Device` " +
		"(`YEAR`, `season`, `DeviceType`, `DeviceName`, `DeviceID`, `STA_NAME`, `STA_ID`, `ORG1_ID`, `ORG1_NAME`, `ORG2_ID`, `ORG2_NAME`, `ORG3_ID`," +
		" `ORG3_NAME`, `Total_DEV`, `Total_Appraisal`, `Total_Good`, `Total_Qualified`, `Total_UnQualified`, `Total_WorkBase`, `Total_Work_People`, `datetime`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)").Prepare()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(maps); i++ {
		m := maps[i]
		Total_Dev := 2000 // 每个车站 设备总数固定,
		for _, y := range years { // 遍历年
			for _, season := range []int{1, 2, 3, 4} { // 遍历季度
				//Total_Dev = Total_Dev + (y-2000)*2 + im //每过一段时间 加一些设备
				total_good := rand.Intn(300)
				total_Qualified := rand.Intn(300)
				total_unQualified := rand.Intn(300)
				total_Appraisal := total_good + total_Qualified + total_unQualified

				total_workbase := total_Appraisal * 3 //工单数
				Total_Work_People := rand.Intn(100)   // 相关人员
				_, err := p.Exec(y, season, return_s(m["DeviceType"]), return_s(m["DeviceName"]), m["DeviceId"].(string), m["STA_NAME"].(string),
					m["STA_ID"].(string), m["ORG1_ID"].(string), m["ORG1_NAME"].(string), m["ORG2_ID"].(string), m["ORG2_NAME"].(string), m["ORG3_ID"].(string), m["ORG3_NAME"].(string),
					Total_Dev, total_Appraisal, total_good, total_Qualified, total_unQualified, total_workbase, Total_Work_People,fmt.Sprintf("%d-%02d-30 23:59:59", y,season*3 ))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	p.Close() // 别忘记关闭 statement

}

func insert_total_year(maps []orm.Params) {
	o := orm.NewOrm()

	p, err := o.Raw("INSERT INTO `Total_Year_Device` " +
		"(`YEAR`,`DeviceType`, `DeviceName`, `DeviceID`, `STA_NAME`, `STA_ID`, `ORG1_ID`, `ORG1_NAME`, `ORG2_ID`, `ORG2_NAME`, `ORG3_ID`," +
		" `ORG3_NAME`, `Total_DEV`, `Total_Appraisal`, `Total_Good`, `Total_Qualified`, `Total_UnQualified`, `Total_WorkBase`, `Total_Work_People`,`datetime`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)").Prepare()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(maps); i++ {
		m := maps[i]
		Total_Dev := 2000 // 每个车站 设备总数固定,
		for _, y := range years { // 遍历年

			//Total_Dev = Total_Dev + (y-2000)*3 //每过一段时间 加一些设备
			total_good := rand.Intn(500)
			total_Qualified := rand.Intn(250)
			total_unQualified := rand.Intn(250)
			total_Appraisal := total_good + total_Qualified + total_unQualified

			total_workbase := total_Appraisal * 3 //工单数
			Total_Work_People := rand.Intn(100)   // 相关人员
			_, err := p.Exec(y, return_s(m["DeviceType"]), return_s(m["DeviceName"]), m["DeviceId"].(string), m["STA_NAME"].(string),
				m["STA_ID"].(string), m["ORG1_ID"].(string), m["ORG1_NAME"].(string), m["ORG2_ID"].(string), m["ORG2_NAME"].(string), m["ORG3_ID"].(string), m["ORG3_NAME"].(string),
				Total_Dev, total_Appraisal, total_good, total_Qualified, total_unQualified, total_workbase, Total_Work_People,fmt.Sprintf("%d-12-31 23:59:59", y ))
			if err != nil {
				fmt.Println(err)
			}

		}
	}
	p.Close() // 别忘记关闭 statement

}
func sql() {
	o1 := orm.NewOrm()
	o1.Using("default")
	var maps []orm.Params
	//num, err := o1.Raw("select * from B_DeviceInfor LIMIT 0,?", 1013).Values(&maps) // 创建天数统计表
	num, err := o1.Raw("select * from B_DeviceInfor GROUP BY STA_ID").Values(&maps) // 创建月，季度，年统计表
	if err == nil && num > 0 {
		//insert_total_days(maps) // 日统计
		insert_total_month(maps) // 月统计
		insert_total_season(maps) // 季度统计
		insert_total_year(maps) // 年统计



		//for i := 0; i < len(maps); i++ {
		//	data := maps[i]
		//	insert_total_days(data) //todo 添加 days表
		//}
	} else {
		fmt.Println("sql err:", err)
	}
}

func CRUD(ver, appname, uuid string) { // 查询任务表 ,判断是否执行
	//GetApps()

	for i := 1; i < 30; i++ {
		days = append(days, i)
	}

	//for i := 0; i < 100; i++ {
	//	fmt.Println("随机总分数", rand.Intn(15)-10)
	//	fmt.Println("随机年", years[rand.Intn(len(years))])
	//	fmt.Println("随机月", month[rand.Intn(len(years))])
	//	fmt.Println("随机日", days[rand.Intn(len(years))])
	//	fmt.Println("随机结果", result[rand.Intn(len(result))])
	//}
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
