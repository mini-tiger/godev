package main

import (
	"encoding/json"
	"fmt"
)

type SuccessResponse struct {
	Code int         `json:"code,omitempty"` // xxx omitemty 忽略空值
	Msg  string      `json:"_"`              // xxx  _ 不解析此字段
	Data interface{} `json:"data"`
}

func main() {
	s := &SuccessResponse{1, "2", 3}
	js, _ := json.Marshal(s)
	fmt.Println(string(js))

	s = &SuccessResponse{}
	js, _ = json.Marshal(s)
	fmt.Println(string(js)) // xxx 不传值 默认为类型的空值

}
