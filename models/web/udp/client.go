package main
import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)
//var host = flag.String("host", "localhost", "host")
//var port = flag.String("port", "37", "port")
//go run timeclient.go -host time.nist.gov
func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("send to server ")) // todo 发送到server
	if err != nil {
		fmt.Println("failed:", err)
		os.Exit(1)
	}
	data := make([]byte, 4)
	_, err = conn.Read(data) //todo 读取server发回来的
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err)
		os.Exit(1)
	}
	t := binary.BigEndian.Uint32(data)
	fmt.Println(time.Unix(int64(t), 0).String())
	os.Exit(0)
}