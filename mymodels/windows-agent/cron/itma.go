package cron

import (
	"godev/mymodels/windows-agent/common/model"
	"godev/mymodels/windows-agent/g"
	"time"
)

var ConfigInterval int = 20 //todo HBS没准备好前，数据更新与配置更新 间隔的默认值
var DataInterval int = 40

func UploadEnvironmentGrid() {
	g.Logger().Printf("start environment grid Ubs Addr:%s", g.Config().Ubs.Addr)
	if g.Config().Ubs.Enabled && g.Config().Ubs.Addr != "" {
		//loadEnvironmentGridConfig(-1)
		//go uploadEnvironmentGrid(time.Duration(5)*time.Second)
		go uploadEnvironmentGrid(time.Duration(DataInterval) * time.Second)
	}
}

func uploadEnvironmentGrid(interval time.Duration) {
	for {
		interval := time.Duration(DataInterval) * time.Second
		//interval := time.Duration(10)*time.Second
		g.Logger().Printf("ready uploadEnvironmentGrid %v", interval)
		req := g.EnvGrid()
		g.Logger().Printf("Call UBS Itma.UploadEnvironmentGrid :%+v \n", req)

		var resp model.AgentUpdateResp

		err := g.UbsClient.Call("Itma.UploadEnvironmentGrid", req, &resp)
		if err != nil {
			g.Logger().Error("call Itma.UploadEnvironmentGrid fail:", err, "Request:", req, "Response:", resp)
			time.Sleep(interval)
			continue
		}

		if interval < 0 {
			break
		}
		time.Sleep(interval)
	}
}

func LoadEnvironmentGridConfig() {
	g.Logger().Printf("start load environment grid config %v -> %s", g.Config().Ubs.Enabled, g.Config().Ubs.Addr)
	if g.Config().Ubs.Enabled && g.Config().Ubs.Addr != "" {
		//loadEnvironmentGridConfig(-1)

		go loadEnvironmentGridConfig(time.Duration(ConfigInterval) * time.Second)
	}
}

func loadEnvironmentGridConfig(interval time.Duration) {
	for {
		interval := time.Duration(ConfigInterval) * time.Second
		g.Logger().Printf("ready get environment grid config %s", interval)

		var req model.NullRpcRequest
		var resp model.EnvGridConfigResponse
		err := g.UbsClient.Call("Itma.GetEnvironmentGridConfig", req, &resp)
		if err != nil {
			g.Logger().Printf("call Itma.GetEnvironmentGridConfig fail: %s  Request:%v Response:%v", err,  req,  resp)
		} else {
			g.Logger().Printf("call Itma.GetEnvironmentGridConfig Response: %v", resp)
			g.GetEnvGridConfig().JsonConfig = resp
			//g.GetEnvGridConfig().ConfigInterval = resp.ConfigInterval
			//g.GetEnvGridConfig().DataInterval = resp.DataInterval

			ConfigInterval = resp.ConfigInterval
			DataInterval = resp.DataInterval

			//var pa []g.ProcessAppsysWorker
			//for _,worker := range resp.Config {
			//	pa = append(pa, g.ProcessAppsysWorker(worker))
			//}
			//g.GetEnvGridConfig().Appsys = pa
		}

		if interval < 0 {
			break
		}
		time.Sleep(interval)
	}
}
