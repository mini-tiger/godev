package model

import (
	"fmt"
)

type PortProcessEnv struct {
	Port    int    `json:"port"`
	Type    int      `json:"type"`
	Process []string `json:"process"`
}

type PortprocessConfResponse struct {
	ConfigInterval    int              `json:"configInterval"`
	DataInterval      int              `json:"dataInterval"`
	Portprocess_slice []PortProcessEnv `json:"portprocess_slice"`
	Timestamp         int64            `json:"timestamp"`
}

func (self *PortprocessConfResponse) String() string {
	return fmt.Sprintf(
		"<PortProcess_Config:%v, Timestamp:%v>",
		self.Portprocess_slice,
		self.Timestamp,
	)
}

//result
type Portprocess_sub_result struct {
	Port int
	Pids []int
	Type int
	Ip_route map[string]Route_link
}


type Route_link struct{
	L map[int]string   //use int sort
}


type Portprocess_result struct {
	Pr []*Portprocess_sub_result
}


func Newroute_list() *Route_link {
	return &Route_link{L: make(map[int]string)}
}
