package funcs

import (
	"godev/mymodels/windows-agent/g"
	"godev/mymodels/windows-agent/common/model"
)

type FuncsAndInterval struct {
	Fs       []func() []*model.MetricValue
	Interval int
}

var Mappers []FuncsAndInterval

func BuildMappers() {
	interval := g.Config().Transfer.Interval
	Mappers = []FuncsAndInterval{
		FuncsAndInterval{
			Fs: []func() []*model.MetricValue{
				AgentMetrics,
				CpuMetrics,
				NetMetrics,
				MemMetrics,
				DeviceMetrics,
				DiskIOMetrics,
				TcpipMetrics,
			},
			Interval: interval,
		},
		FuncsAndInterval{
			Fs: []func() []*model.MetricValue{
				PortMetrics,
				ProcMetrics,
			},
			Interval: interval,
		},
		//FuncsAndInterval{
		//	Fs: []func() []*model.MetricValue{
		//		iisMetrics,
		//		mssqlMetrics,
		//	},
		//	Interval: interval,
		//},
	}
}
