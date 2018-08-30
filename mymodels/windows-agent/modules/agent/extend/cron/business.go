package cron

import (
	"log"
	"time"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/open-falcon/falcon-plus/modules/agent/g"
	extend_g "github.com/open-falcon/falcon-plus/modules/agent/extend/g"
	"fmt"
)




func Updateportprocess_env_task() {
	log.Println("start update portprocess_env  ",g.Config().Heartbeat.Enabled," -> ",g.Config().Heartbeat.Addr)
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		Updateportprocess_env(-1)
		go Updateportprocess_env(time.Duration(5)*time.Second)
		//go Updateportprocess_env(time.Duration(extend_g.EnvPortConfig.DataInterval) * time.Second)
	}
}



func Updateportprocess_env(interval time.Duration)  {
	for {
		log.Println("ready start update portprocess_env ",interval)
		req := extend_g.Getportprocess_data()
		log.Println("collect start update portprocess_env ok:",req.Pr)

		var resp model.SimpleRpcResponse


		err := g.HbsClient.Call("Port.Updateportprocess_env", req, &resp)  //send portprocess -> hbs
		if err != nil || resp.Code != 0 {
			log.Println("call Port.Updateportprocess_env fail:", err, "Request:", req, "Response:", resp)
		}

		if interval < 0 {
			break
		}
		time.Sleep(interval)
	}

}




func Loadportporcess_taskConfig() {
	log.Println("start load portprocess_env config", g.Config().Heartbeat.Enabled, " -> ", g.Config().Heartbeat.Addr)
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		loadportporcessConfig(-1)
		fmt.Println(extend_g.EnvPortConfig.ConfigInterval)

		go loadportporcessConfig(time.Duration(extend_g.EnvPortConfig.ConfigInterval) * time.Second)
		//go loadportporcessConfig(time.Duration(5 * time.Second))
	}
}

func loadportporcessConfig(interval time.Duration) {
	for {
		log.Println("ready get PortProcess Env config ", interval)

		var req model.NullRpcRequest
		var resp model.PortprocessConfResponse
		err := g.HbsClient.Call("Port.GetportConfig", req, &resp)
		if err != nil {
			log.Println("call ==============Port.GetportConfig fail:", err, "Request:", req, "Response:", resp)
		} else {
			log.Println("call ===============Port.GetportConfig Response:", resp)

			gp := extend_g.NewPortProcessConfig()
			gp.JsonConfig = resp
			gp.ConfigInterval = resp.ConfigInterval
			gp.DataInterval = resp.DataInterval


			fmt.Println(resp.ConfigInterval,resp.DataInterval)

			log.Println("===============================")
			log.Println(gp.JsonConfig.Portprocess_slice)

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


