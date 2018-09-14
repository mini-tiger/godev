package cron

import (
	"time"
	"godev/mymodels/windows-agent/common/model"
	"godev/mymodels/windows-agent/g"
	extend_g "godev/mymodels/windows-agent/extend/g"
	"fmt"
)




func Updateportprocess_env_task() {
	g.Logger().Println("start update portprocess_env  ",g.Config().Heartbeat.Enabled," -> ",g.Config().Heartbeat.Addr)
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		Updateportprocess_env(-1)
		//go Updateportprocess_env(time.Duration(5)*time.Second)
		go Updateportprocess_env(time.Duration(extend_g.EnvPortConfig.DataInterval) * time.Second)
	}
}



func Updateportprocess_env(interval time.Duration)  {
	for {
		g.Logger().Println("ready start update portprocess_env ",interval)
		req := extend_g.Getportprocess_data()
		g.Logger().Println("collect start update portprocess_env ok:",req.Pr)

		var resp model.SimpleRpcResponse


		err := g.HbsClient.Call("Port.Updateportprocess_env", req, &resp)  //send portprocess -> hbs
		if err != nil || resp.Code != 0 {
			g.Logger().Println("call Port.Updateportprocess_env fail:", err, "Request:", req, "Response:", resp)
		}
		if err == nil && resp.Code == 0{
			g.Logger().Println("call Port.Updateportprocess_env Sucess:", err, "Request:", req, "Response:", resp)
		}
		if interval < 0 {
			break
		}
		time.Sleep(interval)
	}

}




func Loadportporcess_taskConfig() {
	g.Logger().Println("start load portprocess_env config", g.Config().Heartbeat.Enabled, " -> ", g.Config().Heartbeat.Addr)
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		loadportporcessConfig(-1)
		//fmt.Println(extend_g.EnvPortConfig.ConfigInterval)

		go loadportporcessConfig(time.Duration(extend_g.EnvPortConfig.ConfigInterval) * time.Second)
		//go loadportporcessConfig(time.Duration(5 * time.Second))
	}
}

func loadportporcessConfig(interval time.Duration) {
	for {
		//g.Logger().Println("ready get PortProcess Env config ", interval)

		var req model.NullRpcRequest
		var resp model.PortprocessConfResponse
		err := g.HbsClient.Call("Port.GetportConfig", req, &resp)
		if err != nil {
			g.Logger().Println("call ==============Port.GetportConfig fail:", err, "Request:", req, "Response:", resp)
		} else {
			g.Logger().Println("call ===============Port.GetportConfig Response:", resp)

			gp := extend_g.NewPortProcessConfig()
			gp.JsonConfig = resp
			gp.ConfigInterval = resp.ConfigInterval
			gp.DataInterval = resp.DataInterval


			fmt.Println(resp.ConfigInterval,resp.DataInterval)

			g.Logger().Println("===============================")
			g.Logger().Println(gp.JsonConfig.Portprocess_slice)

			//for i,v := range gp.JsonConfig.Portprocess_slice {
			//	fmt.Printf("index:%d,value:%+v,%T\n",i,v,v)
			//}
		}

		if interval < 0 {
			break
		}
		time.Sleep(interval)
	}
}


