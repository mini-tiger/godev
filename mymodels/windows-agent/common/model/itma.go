

package model

import (
	"fmt"
	"strings"
)

type Link struct {
	LocalPort	string		`json:"localPort"`
	RemoteIp	string		`json:"remoteIp"`
	RemotePort	string		`json:"remotePort"`
}
func (this *Link) String() string {
	return fmt.Sprintf("%s -- %s:%s", this.LocalPort, this.RemoteIp, this.RemotePort)
}

type Appsys struct {
	Name		string		`json:"name"`
	Type		string		`json:"type"`
	Links		[]Link		`json:"links"`
}
func (this *Appsys) String() string {
	return fmt.Sprintf("Appsys:%s(%s)\n%s", this.Name, this.Type, this.Links)
}

type Route struct {
	Target		string		`json:"target"`
	Traces		[]string	`json:"traces"`
}
func (this *Route) String() string {
	return fmt.Sprintf("Route:%s\n%s",this.Target,this.Traces)
}

//type EnvGrid struct {
//	Hostname    string		`json:"hostname"`
//	Address     []string	`json:"address"`
//	Manufacturer string		`json:"manufacturer"`
//	ProductName  string		`json:"productName"`
//	Version      string		`json:"version"`
//	SerialNumber string		`json:"serialNumber"`
//	Appsyss  	[]Appsys	`json:"appsyss"`
//	Routes 		[]Route		`json:"routes"`
//}
//func (this *EnvGrid) String() string {
//	return fmt.Sprintf(
//		"Hostname:%s, Address:%s, Manufacturer:%s, ProductName:%s, Version:%s, SerialNumber:%s, Appsys:%s, Route:%s",
//		this.Hostname,
//		this.Address,
//		this.Manufacturer,
//		this.ProductName,
//		this.Version,
//		this.SerialNumber,
//		this.Appsyss,
//		this.Routes,
//	)
//}


type DefaultProcessAppsysWorker struct {
	Appsys 	string 		`json:"appsys"`
	Type 	string 		`json:"type"`
	Relation string 		`json:"relation"`
	Process 	[]string 	`json:"process"`
}






func (this *DefaultProcessAppsysWorker) String() string {
	return fmt.Sprintf(
		"%s(%s) %s :%s",
		this.Appsys,
		this.Type,
		this.Relation,
		this.Process,
	)
}
func (worker DefaultProcessAppsysWorker) GetAppsys() string {
	return worker.Appsys;
}
func (worker DefaultProcessAppsysWorker) GetType() string {
	return worker.Type;
}
func (worker DefaultProcessAppsysWorker) Find(line string) bool {
	if worker.Relation == "or" {
		for _,p := range worker.Process {
			if strings.Contains(line, p) {

				return true
			}
		}
		return false
	} else {
		for _,p := range worker.Process {
			if !strings.Contains(line, p) {

				return false
			}
		}
		return true
	}
}

type AgentUpdateResp struct {
	Update bool `json:"update"`
	NewVersion string `json:"new_version"`
	FtpInfo Ftpinfo `json:"ftpinfo"`
	RunDir string `json:"run_dir"`
	RowId int64 `json:"row_id"`
}
type Ftpinfo struct {
	Ip       string
	Port     int
	User     string
	Passwd   string
	Dirpath  string
	Filename string
}

type EnvGridConfigResponse struct {
	ConfigInterval  int		`json:"configInterval"`
	DataInterval	int		`json:"dataInterval"`
	//Config   		[]DefaultProcessAppsysWorker	`json:"config"`
	Timestamp 		int64	`json:"timestamp"`
}

//func (this *EnvGridConfigResponse) String() string {
//	return fmt.Sprintf(
//		"<Config:%v, Timestamp:%v>",
//		this.Config,
//		this.Timestamp,
//	)
//}



type EnvGridExt struct {
	Address     string	`json:"address"`
	Cpucores	int			`json:"cpucores,string"`
	Cpumhz		int			`json:"cpumhz"`
	Cpumodel	string		`json:"cpumodel"`
	Disktotal	int64		`json:"disktotal,int"`
	Osbit		string		`json:"osbit"`
	MAC			string		`json:"mac"`
	Osname		string		`json:"osname"`
	Memtotal	int64		`json:"memtotal,int"`
	UUID 		string			`json:"uuid"`
	Biz			int				`json:"biz"`

}

type EnvGrid struct {
	Hostname    string		`json:"hostname"`
	Manufacturer string		`json:"manufacturer"`
	ProductName  string		`json:"productName"`
	Version      string		`json:"version"`  // 这个不是agent version
	SerialNumber string		`json:"serialNumber"`
	Appsyss  	[]Appsys	`json:"appsyss"`
	Routes 		[]Route		`json:"routes"`
	Extinfo		EnvGridExt   `json:"extinfo"`
	AgentInfo   AgentInfo		`json:"agent_info"`
	//Address     []string	`json:"address"`

}

type AgentInfo struct {
	Version string `json:"version"`
}