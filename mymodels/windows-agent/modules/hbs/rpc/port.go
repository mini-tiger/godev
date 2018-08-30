package rpc

import (
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/modules/hbs/g"
	"time"
	"fmt"
)

func (p *Port) GetportConfig(args *model.NullRpcRequest, reply *model.PortprocessConfResponse) error {
	//	log.Println("GetEnvironmentGridConfig")

	reply.ConfigInterval = g.GetEnvGridConfig().ConfigInterval
	reply.DataInterval = g.GetEnvGridConfig().DataInterval
	reply.Portprocess_slice = g.GetEnvGridConfig().Monitor_port
	reply.Timestamp = time.Now().Unix()

	return nil
}

func (p *Port) Updateportprocess_env(args *model.Portprocess_result, reply *model.SimpleRpcResponse) error {
	//fmt.Println(args)
	for _, v := range args.Pr {
		//fmt.Printf("Port:%d Pids:%v,type:%d \n", v.Port, v.Pids, v.Type)
		for k, vv := range v.Ip_route {
			fmt.Printf("port:%d,ip routes ip:%v,routes:%+v\n ",v.Port ,k,vv.L)
		}
	}
	return nil
}
