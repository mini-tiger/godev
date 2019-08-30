// example.go
package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/chenjiandongx/go-echarts/charts"
)
// 使用go run saveHtml.go 运行
var nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
var seed = rand.NewSource(time.Now().UnixNano())

func randInt() []int {
	cnt := len(nameItems)
	r := make([]int, 0)
	for i := 0; i < cnt; i++ {
		r = append(r, int(seed.Int63()) % 50)
	}
	return r
}

func main() {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true},charts.InitOpts{PageTitle:"this is tj first go-charts"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(f)
}