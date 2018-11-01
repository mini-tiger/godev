package main

import "fmt"

type SimpleRpcResponse struct {
	Code int `json:"code"`
}

func (this *SimpleRpcResponse) String() string {
	return fmt.Sprintf("<Code: %d>", this.Code)
}

type NullRpcRequest struct {
}
func main()  {
	s:=SimpleRpcResponse{1}
	fmt.Println(s.String())

	}