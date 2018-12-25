package cron

import (
	"godev/works/train/g"
	"time"
	"sync"
	"fmt"
	"godev/tjtools/net/udp"
	"encoding/json"
)

const imageExt = "_image"  //redis 中存入IMAGES SET时key的名字  vehicle_serial+imageExt
const vehicleInterval = 10 // 扫描vehicle 数据库间隔

type VehicleResult struct {
	sync.RWMutex
	VehicleJson *VehicleJson // 每个要 车厢的json
}

type VehicleJson struct {
	TrainID       string   `json:"TRAIN_ID"`
	PASS_TIME     string   `json:"PASS_TIME"`
	FOLDER        string   `json:"FOLDER"`
	VehicleOrder  string   `json:"VEHICLE_ORDER"`
	VehicleType   string   `json:"VEHICLE_TYPE"`
	POS_AB_FALG   string   `json:"AB"`
	Images        []string `json:"Images"`
	SaveUnixStamp int      `json:"save_unix_stamp,omitempty"` // todo 发送时要 清空此项, 存入redis对比 时间是否超过最大    "maxvehicletime":120
	Send          bool     `json:"send,omitempty"`            // todo 是否发送过 ，存入redis对比 时间是否超过最大
}

func GetPicMission(result map[string]string) {
	g.Logger().Debug("开始获取train_serial:%s ", result["TRAIN_SERIAL"])
	go BeginMVehicle(result)
}

func BeginMVehicle(trainInfo map[string]string) {
	now := g.GetNow()
	for { //  todo 每次都获取 train_serial下面的所有 veh ，每次 判断存入redis的vehJSON是否 到时，到时则获取sys1发送
		CheckMVehcile(trainInfo)
		if int(time.Now().Unix())-now >= g.Config().Train.TrainInterval+60 { //todo 当前时间 － 循环开始时间  超过两个列车的间隔延长60秒
			// 把没发送的，所有都获取 sys 发出去
			break
		}

		time.Sleep(time.Duration(vehicleInterval) * time.Second)
	}

}
func CheckMVehcile(trainInfo map[string]string) {
	now := g.GetNow()
	sql := "SELECT TRAIN_SERIAL,VEHICLE_SERIAL,VEHICLE_ORDER,VEHICLE_TYPE,POS_AB_FLAG FROM TF_OP_VEHICLE " +
		"WHERE train_serial  ='32F6ACC0B0A243FC974BC2BCECF2AA51' and vehicle_order is not null  order by vehicle_order "
	if mveh, b := g.GetSqlData(sql); b { // todo 扫描 train_serial 下的所有vehicle
		for _, v := range mveh {
			if g.Redis.StringExists(v["VEHICLE_SERIAL"]) { // 存入过判断是否超时，是否发送过
				go SendJson(v["VEHICLE_SERIAL"])
			} else {
				var tmpVehJson VehicleResult
				tmpVehJson.SetRedis(trainInfo, v, now)
			}
		}
	}
}

func (v *VehicleResult) SetRedis(trainInfo, veh map[string]string, now int) {
	v.Lock()
	defer v.Unlock()

	v.VehicleJson = &VehicleJson{TrainID: trainInfo["TRAIN_ID"],
		PASS_TIME: trainInfo["PASS_TIME:"],
		FOLDER: trainInfo["INDEX_ID"],
		VehicleOrder: veh["VEHICLE_ORDER"],
		VehicleType: veh["VEHICLE_TYPE"],
		POS_AB_FALG: veh["POS_AB_FLAG"],
		SaveUnixStamp: now, // 写入当前时间，后面判断是超时
		Send: false} //没有发送过

	g.Redis.SetJson(veh["VEHICLE_SERIAL"], v.VehicleJson, time.Duration(1800)*time.Second) // 超时时间30分钟
}

func SendJson(vehicleKey string) {
	vehjsonStruct := g.Redis.GetJson(vehicleKey)
	if vehjsonStruct.Send || (g.GetNow()-vehjsonStruct.SaveUnixStamp) < g.Config().Train.MaxVehicleTime {
		return // todo 发送过或者没有到发送的时间
	} else {
		g.Logger().Printf("准备发送vehicleSerial:%s\n", vehicleKey)
		vehImages := GetImages(vehicleKey, &vehjsonStruct)
		vehjsonStruct.Images = vehImages
		byteVeh, err := json.Marshal(&vehjsonStruct)
		if err != nil {
			g.Logger().Error("vehicle 转换为json 失败err:%s, data:%+v", err, &vehjsonStruct)
		}
		err = udp.SendUdp(g.Config().UdpInfo.UdpAddr, byteVeh)
		if err != nil {
			g.Logger().Error("发送 vehicle json 失败 err:%s, data:%+v", err, &vehjsonStruct)
		}
	}
}

func GetImages(vehicleKey string, vehjsonStruct *VehicleJson) (tmpSlice []string) {
	sql := fmt.Sprintf("select train_serial,vehicle_serial,img_index,img_name,vehicle_order, t.rowid  from SYS1 t "+
		"here vehicle_serial is not null and vehicle_order is not null and vehicle_order = %s and vehicle_serial='%s'", vehjsonStruct.VehicleOrder, vehicleKey)
	if result, b := g.GetSqlData(sql); b {
		//tmpSlice := make([]string, 0)
		for _, v := range result {
			tmpSlice = append(tmpSlice, v["IMG_NAME"])

		}
	} else {
		g.Logger().Error("GetImage vehicleKey:%s Error\n", vehicleKey)
	}
	return
}
