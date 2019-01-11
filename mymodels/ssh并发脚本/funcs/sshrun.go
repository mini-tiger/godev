package funcs

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"godev/mymodels/ssh并发运行脚本/g"
	"tjtools/nmap"
)

var hostSet mapset.Set
var passwdSet mapset.Set
var HostPass *nmap.SafeMap = nmap.NewSafeMap()
var FailHosts []string = make([]string, 0)

var hostchan chan struct{}

func SSHRun() {

	hosts := g.Config().Hosts
	hostSet = mapset.NewSetFromSlice(hosts) // 去重
	hostchan = make(chan struct{}, hostSet.Cardinality())
	passwds := g.Config().PasswdMap

	passwdSet = mapset.NewSetFromSlice(passwds) // 去重
	for _, host := range hostSet.ToSlice() {
		fmt.Println(host)

		go SSHSingle(host)
	}

	for i := 0; i < hostSet.Cardinality(); i++ {
		<-hostchan
	}

}
func SSHSingle(host interface{}) {
	defer func() {
		hostchan <- struct{}{}
	}()
	h := host.(string)
	for _, pass := range passwdSet.ToSlice() {
		ssh1 := g.New_ssh(22, []string{h, "root", pass.(string)}...)
		//fmt.Println(ssh1)
		err := ssh1.Connect()
		if err == nil {
			HostPass.Put(h, pass)
			pass = pass.(string)
			return
		}
	}
	FailHosts = append(FailHosts, h)
}
