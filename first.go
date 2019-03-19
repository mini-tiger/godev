package main

import (
	"strconv"
	"net"
	"fmt"
	"log"
)

func main()  {
	fmt.Printf("%+v\n",CheckTCPPortUsed(8080))
	//port:="80"
	//cmd := exec.Command("cmd.exe ", "/c", fmt.Sprintf("netstat -ano|findstr %s",port))
	//
	//stdout, err := cmd.StdoutPipe()
	//if err != nil {
	//	log.Fatalf("ips_business run port:%d, err:%v")
	//
	//}
	//cmd.Start()
	//
	//reader := bufio.NewReader(stdout)
	//pids:=make([]string,0)
	//for {
	//	line, err := reader.ReadString('\n')
	//	if err != nil || io.EOF == err {
	//		break
	//	}
	//	line = strings.TrimSpace(strings.Trim(line, "\n"))
	//
	//	//fmt.Println([]byte(line))
	//	pids = append(pids, line)
	//}
	//
	//cmd.Wait()
	//for i:=0;i<len(pids);i++{
	//	ii:=strings.Split(pids[i],"    ")
	//	iii:=strings.Split(ii[1],":")
	//	fmt.Println(iii[1],len(iii))
	//}
	}

func IsTCPPortUsed(addr string, port int64) bool {
	connString := addr + strconv.FormatInt(port, 10)
	fmt.Println(connString)
	conn, err := net.Dial("tcp", connString)
	if err != nil {
		log.Println(connString, conn, err)
		return false
	}
	conn.Close()
	return true
}

func CheckTCPPortUsed(port int64) bool {
	if IsTCPPortUsed("0.0.0.0:", port) {
		return true
	}
	if IsTCPPortUsed("127.0.0.1:", port) {
		return true
	}
	if IsTCPPortUsed("[::1]:", port) {
		return true
	}
	if IsTCPPortUsed("[::]:", port) {
		return true
	}
	return false
}