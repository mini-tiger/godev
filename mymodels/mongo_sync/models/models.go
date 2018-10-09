package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

//type Person struct {
//	Id_   bson.ObjectId `bson:"_id"`
//	Name  string        `bson:"name"`
//	Phone string        `bson:"phone"`
//}

type CC_ApplicationBase struct {
	Id_                 bson.ObjectId `bson:"_id"`
	Bk_biz_tester       string
	Time_zone           string
	Create_time         time.Time
	Last_time           time.Time
	Life_cycle          string
	Bk_biz_developer    string
	Bk_biz_maintainer   string
	Bk_supplier_id      int64
	Defaultt            int64         `bson:"default"`
	Language            string
	Operator            string
	Bk_biz_id           int64
	Bk_biz_name         string
	Bk_biz_productor    string
	Bk_supplier_account string
}

type CC_History struct {
	Id_         bson.ObjectId `bson:"_id"`
	Content     string
	User        string
	Create_time time.Time
	Id          string        `bson:"id"`
}
type CC_Host struct {
	Id_              bson.ObjectId `bson:"_id"`
	Bk_host_outerip  string
	Bk_mac           string
	Bk_state_name    string
	Bk_cpu_mhz       int64
	Bk_mem           int64
	Bk_os_bit        string
	Import_from      string
	Bk_cloud_id      int64
	Bk_host_innerip  string
	Bk_os_name       string
	Bk_os_type       string
	Bk_comment       string
	Bk_sla           string
	Bk_sn            string
	Bk_host_name     string
	Bk_province_name string
	Bk_service_term  string
	Operator         string
	Bk_cpu           int64
	Bk_isp_name      string
	Bk_outer_mac     string
	Create_time      time.Time
	Bk_asset_id      string
	Bk_bak_operator  string
	Bk_cpu_module    string
	Bk_disk          int64
	Bk_os_version    string
	Bk_host_id       int64
	Last_time        time.Time
}
