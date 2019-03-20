package cron

import (
	"fmt"

	"time"

	"godev/mymodels/windows-agent/g"
	"godev/mymodels/windows-agent/common/model"
)

func ReportAgentStatus() {
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		go reportAgentStatus(time.Duration(g.Config().Heartbeat.Interval) * time.Second)
	}
}

func reportAgentStatus(interval time.Duration) {
	for {
		hostname, err := g.Hostname()
		if err != nil {
			hostname = fmt.Sprintf("error:%s", err.Error())
		}
		manufacturer, productName, version, serialNumber := g.GetHardware()
		req := model.AgentReportRequest{
			Hostname:      hostname,
			IP:            g.IP(),
			AgentVersion:  g.VERSION,
			PluginVersion: "",
			Manufacturer:  manufacturer,
			ProductName:   productName,
			SystemVersion: version,
			SerialNumber:  serialNumber,
		}

		var resp model.SimpleRpcResponse
		err = g.HbsClient.Call("Agent.ReportStatus", req, &resp)
		if err != nil || resp.Code != 0 {
			g.Logger().Error("call Agent.ReportStatus fail: %s Request: %v Response:%v", err, req, resp)
		}

		time.Sleep(interval)
	}
}
