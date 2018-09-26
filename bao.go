package main

import (
	"github.com/shirou/gopsutil/host"
	"log"
	"fmt"
	"reflect"
)

func main() {
	h, err := host.Info()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(h)
	fmt.Println(host.Uptime())
	fmt.Println(host.BootTime())
	fmt.Println(host.PlatformInformation())
	fmt.Println(host.Users())
	fmt.Println(host.SensorsTemperatures())
	fmt.Println(host.Virtualization())
	fmt.Println(host.KernelVersion())
	i := new(host.InfoStat)
	i,err = host.Info()
	if err!=nil{
		log.Fatalln(err)
	}
	fmt.Println(i)
	tt(i)
	ttr(i)
}
func ttr(i interface{})  {
	s:=reflect.ValueOf(i)
	switch s.Kind() {
	case reflect.Ptr:
		fmt.Println("ptr")
	}
}

func tt(i interface{}){
	switch i.(type) {
	case int:
		fmt.Println("int")
	case host.InfoStat:
		fmt.Println("infostat")
	case *host.InfoStat:
		fmt.Println("ptr infostat")

	}
}