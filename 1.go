package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Data struct {
	X string      `json:"x"`
	Y interface{} `json:"y"`
	S string      `json:"s"`
}

func main() {

	GenTime()
	//GenUser()
	//Both()
}

func GenTime() {
	startTime := time.Now()
	max := 7
	fmt.Println("[")
	for i := 0; i < max; i++ {
		//m:=make(map[string]interface{},0)
		var d Data
		d.X = startTime.Add(time.Duration(i) * 24 * time.Hour).Format("2006-01-02")
		d.Y = rand.Intn(10)

		resp, _ := json.MarshalIndent(d, "", "    ")

		if i == max-1 {
			fmt.Printf("%+v\n", string(resp))
		} else {
			fmt.Printf("%+v,\n", string(resp))
		}
	}
	fmt.Println("]")
}

func GenUser() {
	start := "人员"
	max := 7

	for i := 1; i < max; i++ {
		//m:=make(map[string]interface{},0)
		var d Data
		d.X = fmt.Sprintf("%s%d", start, i)
		d.Y = rand.Intn(10)

		resp, _ := json.MarshalIndent(d, "", "    ")

		switch true {
		case i == max-1:
			fmt.Printf("%+v]\n", string(resp))
			break
		case i == 1:
			fmt.Printf("[\n%+v,\n", string(resp))
		default:
			fmt.Printf("%+v,\n", string(resp))
		}

	}
}

func Both() {
	start := "人员"
	startTime := time.Now()
	max := 7
	fmt.Println("[")
	for i := 0; i <= max; i++ {
		//m:=make(map[string]interface{},0)
		var d Data
		d.X = startTime.Add(time.Duration(i) * 24 * time.Hour).Format("2006-01-02")
		var maxii = rand.Intn(max)
		for ii := 1; ii <= maxii; ii++ {
			d.Y = rand.Intn(10)
			d.S = fmt.Sprintf("%s%d", start, ii)
			resp, _ := json.MarshalIndent(d, "", "    ")

			if maxii == ii && i == max {
				fmt.Printf("%+v\n", string(resp))
			} else {
				fmt.Printf("%+v,\n", string(resp))
			}

			//if i==max-1{
			//	fmt.Printf("%+v\n",string(resp))
			//}
		}

	}
	fmt.Println("]")
}
