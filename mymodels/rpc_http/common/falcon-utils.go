package common

import (
	"errors"
	"log"
)

type Falcon struct {}

type Recv struct {
	Num1,Num2 int
	Method string
}

type ResultResp struct {
	Result int
}


func (t *Falcon) Compute(args *Recv, reply *ResultResp) error {
	log.Println("接收到的参数",args)
	if args.Num2 == 0 {
		return errors.New("除数不能为零！")
	}
	reply.Result = args.Num1 / args.Num2
	log.Println("结果是",reply.Result)
	return nil
}