package cron

import (
	"godev/mymodels/windows-agent/common/model"
	"godev/mymodels/windows-agent/g"
	"time"
)

var ConfigInterval int = 20 //todo HBS没准备好前，数据更新与配置更新 间隔的默认值
var DataInterval int = 40

func UploadEnvironmentGrid() {
	g.Logger().Println("start environment grid ", g.Config().Ubs.Enabled, " -> ", g.Config().Ubs.Addr)
	if g.Config().Ubs.Enabled && g.Config().Ubs.Addr != "" {
		loadEnvironmentGridConfig(-1)
		//go uploadEnvironmentGrid(time.Duration(5)*time.Second)
		go uploadEnvironmentGrid(time.Duration(DataInterval) * time.Second)
	}
}

func uploadEnvironmentGrid(interval time.Duration) {
	for {
		interval := time.Duration(DataInterval) * time.Second
		g.Logger().Println("++++++++++++++++++++++++++++++++++ready Send agent version", interval)
		req := g.EnvGrid()
		g.Logger().Printf("Call UBS Itma.UploadEnvironmentGrid :%+v \n", req)

		var resp model.AgentUpdateResp

		err := g.HbsClient.Call("Itma.UploadEnvironmentGrid", req, &resp)
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
	g.Logger().Println("start load environment grid config", g.Config().Ubs.Enabled, " -> ", g.Config().Ubs.Addr)
	if g.Config().Ubs.Enabled && g.Config().Ubs.Addr != "" {
		//loadEnvironmentGridConfig(-1)

		go loadEnvironmentGridConfig(time.Duration(ConfigInterval) * time.Second)
	}
}

func loadEnvironmentGridConfig(interval time.Duration) {
	for {
		g.Logger().Println("ready get environment grid config ", interval)

		var req model.NullRpcRequest
		var resp model.EnvGridConfigResponse
		err := g.HbsClient.Call("Itma.GetEnvironmentGridConfig", req, &resp)
		if err != nil {
			g.Logger().Println("call Itma.GetEnvironmentGridConfig fail:", err, "Request:", req, "Response:", resp)
		} else {
			g.Logger().Println("call Itma.GetEnvironmentGridConfig Response:", resp)
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
