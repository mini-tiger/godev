package cron

import (
	"godev/mymodels/windows-agent/common/model"
	"godev/mymodels/windows-agent/g"
	"time"
)

func UploadEnvironmentGrid() {
	g.Logger().Println("start environment grid ", g.Config().Heartbeat.Enabled, " -> ", g.Config().Heartbeat.Addr)
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		loadEnvironmentGridConfig(-1)
		//go uploadEnvironmentGrid(time.Duration(5)*time.Second)
		go uploadEnvironmentGrid(time.Duration(g.GetEnvGridConfig().DataInterval) * time.Second)
	}
}

func uploadEnvironmentGrid(interval time.Duration) {
	for {
		g.Logger().Println("ready collect environment grid", interval)
		req := g.EnvGrid()
		g.Logger().Println("collect environment grid ok:", req)

		var resp model.SimpleRpcResponse
		err := g.HbsClient.Call("Itma.UploadEnvironmentGrid", req, &resp)
		if err != nil || resp.Code != 0 {
			g.Logger().Println("call Itma.UploadEnvironmentGrid fail:", err, "Request:", req, "Response:", resp)
		}
		if err == nil && resp.Code == 0 {
			g.Logger().Println("call Itma.UploadEnvironmentGrid Success:", err, "Request:", req, "Response:", resp)
		}
		if interval < 0 {
			break
		}
		time.Sleep(interval)
	}
}

func LoadEnvironmentGridConfig() {
	g.Logger().Println("start load environment grid config", g.Config().Heartbeat.Enabled, " -> ", g.Config().Heartbeat.Addr)
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		loadEnvironmentGridConfig(-1)

		go loadEnvironmentGridConfig(time.Duration(g.GetEnvGridConfig().ConfigInterval) * time.Second)
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
			g.GetEnvGridConfig().ConfigInterval = resp.ConfigInterval
			g.GetEnvGridConfig().DataInterval = resp.DataInterval
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
