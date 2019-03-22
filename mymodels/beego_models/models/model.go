package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func RegitDB(dbstr string) {
	orm.RegisterModel(new(TotalDayDevice))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)

	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql",
		dbstr,
		maxIdle, maxConn)
}

type TotalDayDevice struct {
	Id             int    `orm:"pk;auto"`
	Year           int    `orm:"column(year);size(4)"`
	Month          int    `orm:"column(month);size(2)"`
	Days           int    `orm:"column(days);size(2)"`
	DeviceId       string `orm:"column(DeviceId);size(36)"`       //IP
	DeviceTypeName string `orm:"column(DeviceTypeName);size(36)"` //IP
	DeviceName     string `orm:"column(DeviceName);size(36)"`     //IP
	DeviceType     string `orm:"column(DeviceType);size(36)"`     //IP
	OUCode         string `orm:"column(OUCode);size(30)"`         //IP
	ORG1ID         string `orm:"column(ORG1_ID);size(30)"`        //IP
	ORG1NAME       string `orm:"column(ORG1_NAME);size(30)"`      //IP
	ORG2ID         string `orm:"column(ORG2_ID);size(30)"`        //IP
	ORG2NAME       string `orm:"column(ORG2_NAME);size(30)"`      //IP
	ORG3ID         string `orm:"column(ORG3_ID);size(30)"`        //IP
	ORG3NAME       string `orm:"column(ORG3_NAME);size(30)"`
	STAID          string `orm:"column(STA_ID);size(30)"`
	STANAME        string `orm:"column(STA_NAME);size(30)"`
	Score          int    `orm:"column(Score);size(10)"`
	Result         string `orm:"column(Result);size(12)"`
	ItemClassName  string `orm:"column(ItemClassName);size(255)"`
	DateTime 	time.Time  `orm:"column(datetime)"`
}

//type BDeviceInfor struct {
//	Id          int       `orm:"pk;auto"`
//	BusinessId  string    `orm:"column(BusinessId);size(36)"`
//	DeviceId    string    `orm:"column(DeviceId);size(36)"`
//	DeviceName  string    `orm:"column(DeviceName);size(36)"` //IP
//	DeviceType  string    `orm:"column(DeviceType);size(36)"` //IP\
//	CreatedBy   string    `orm:"column(CreatedBy);size(36)"`  //IP
//	CreatedDate time.Time `orm:"column(CreatedDate);type(datetime)"`
//	Status      int       `orm:"column(Status);size(11)"`
//	TaskId      string    `orm:"column(TaskId);size(36)"` //IP
//	Remark      string    `orm:"column(Remark);size(36)"` //IP
//
//	OUCode        string `orm:"column(OUCode);size(30)"`    //IP
//	ORG1ID        string `orm:"column(ORG1_ID);size(30)"`   //IP
//	ORG1NAME      string `orm:"column(ORG1_NAME);size(30)"` //IP
//	ORG2ID        string `orm:"column(ORG2_ID);size(30)"`   //IP
//	ORG2NAME      string `orm:"column(ORG2_NAME);size(30)"` //IP
//	ORG3ID        string `orm:"column(ORG3_ID);size(30)"`   //IP
//	ORG3NAME      string `orm:"column(ORG3_NAME);size(30)"`
//	STAID         string `orm:"column(STA_ID);size(30)"`
//	STANAME       string `orm:"column(STA_NAME);size(30)"`
//	Score         int    `orm:"column(Score);size(10)"`
//	Result        string `orm:"column(Result);size(12)"`
//	ItemClassName string `orm:"column(ItemClassName);size(255)"`
//}

//自定义表名
func (m *TotalDayDevice) TableName() string {
	return "Total_day_Device11"
}

//func (b *BDeviceInfor) TableName() string {
//	return "B_DeviceInfor"
//}

// 设置引擎为 INNODB
func (m *TotalDayDevice) TableEngine() string {
	return "INNODB"
}
