package cron

import (
	"time"

	"godev/mymodels/windows-agent/funcs"
	"godev/mymodels/windows-agent/g"
	"godev/mymodels/windows-agent/common/model"
)

func InitDataHistory() {
	for {
		funcs.UpdateCpuStat()
		time.Sleep(g.COLLECT_INTERVAL)
	}
}

func Collect() {

	if !g.Config().Transfer.Enabled {
		return
	}

	if len(g.Config().Transfer.Addrs) == 0 {
		return
	}

	for _, v := range funcs.Mappers {
		go collect(int64(v.Interval), v.Fs)
	}
}

func collect(sec int64, fns []func() []*model.MetricValue) {
	t := time.NewTicker(time.Second * time.Duration(sec)).C
	for {
		<-t

		hostname, err := g.Hostname()
		if err != nil {
			continue
		}

		mvs := []*model.MetricValue{}
		ignoreMetrics := g.Config().IgnoreMetrics

		for _, fn := range fns {
			items := fn()
			if items == nil {
				continue
			}

			if len(items) == 0 {
				continue
			}

			for _, mv := range items {
				if b, ok := ignoreMetrics[mv.Metric]; ok && b {
					continue
				} else {
					mvs = append(mvs, mv)
				}
			}
		}

		now := time.Now().Unix()
		for j := 0; j < len(mvs); j++ {
			mvs[j].Step = sec
			mvs[j].Endpoint = hostname
			mvs[j].Timestamp = now
		}
		if len(mvs) > 1{
			g.Logger().Println("采集系统指标完成")
		}
		g.SendToTransfer(mvs)

	}
}
