package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {

	r := &net.Resolver{
		//PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: 10 * time.Second,
			}
			return d.DialContext(ctx, network, "192.168.1.11:53")
		},
	}

	ips, _ := r.LookupHost(context.Background(), "21vianet.com")
	fmt.Println(ips)
}
