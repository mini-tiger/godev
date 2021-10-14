package main

import (
	"fmt"
	netdata "github.com/netdata/go.d.plugin"
	"net"
)

var (
	// specify the range of trusted IP addresses
	start = net.ParseIP("198.162.1.100")
	end   = net.ParseIP("198.162.1.199")
)

func main() {
	fmt.Println(netdata.ParseRange("192.168.0.0/16"))
}
