package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var bodyMap map[string]interface{} = make(map[string]interface{}, 3)
	for i := 0; i < 10; i++ {

		bodyMap[strconv.Itoa(i)] = i
		fmt.Println(unsafe.Sizeof(bodyMap), len(bodyMap))
	}

}
