package utils

import (
	"log"
	"math/rand"
	"net/http"
	"path"

	"github.com/chenjiandongx/go-echarts/charts"
	"strings"
)

var IpAddr = "http://127.0.0.1"
var Port = "8081"
var Host = strings.Join([]string{IpAddr, Port}, ":")

const (
	maxNum = 50
)

type router struct {
	name string
	charts.RouterOpts
}

var (
	nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
	foodItems = []string{"面包", "牛奶", "奶茶", "棒棒糖", "加多宝", "可口可乐"}

	rangeColor = []string{
		"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
		"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
	}

	Hours = [...]string{
		"12a", "1a", "2a", "3a", "4a", "5a", "6a", "7a", "8a", "9a", "10a", "11a",
		"12p", "1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p", "9p", "10p", "11p",
	}

	Days = [...]string{"Saturday", "Friday", "Thursday", "Wednesday", "Tuesday", "Monday", "Sunday"}

	mapData = map[string]float32{
		"北京":   float32(rand.Intn(150)),
		"上海":   float32(rand.Intn(150)),
		"深圳":   float32(rand.Intn(150)),
		"辽宁":   float32(rand.Intn(150)),
		"青岛":   float32(rand.Intn(150)),
		"山西":   float32(rand.Intn(150)),
		"陕西":   float32(rand.Intn(150)),
		"乌鲁木齐": float32(rand.Intn(150)),
		"齐齐哈尔": float32(rand.Intn(150)),
	}

	guangdongData = map[string]float32{
		"深圳市": float32(rand.Intn(150)),
		"广州市": float32(rand.Intn(150)),
		"湛江市": float32(rand.Intn(150)),
		"汕头市": float32(rand.Intn(150)),
		"东莞市": float32(rand.Intn(150)),
		"佛山市": float32(rand.Intn(150)),
		"云浮市": float32(rand.Intn(150)),
		"肇庆市": float32(rand.Intn(150)),
		"梅州市": float32(rand.Intn(150)),
	}

	shantouData = map[string]float32{
		"澄海区": float32(rand.Intn(150)),
		"潮阳区": float32(rand.Intn(150)),
		"潮南区": float32(rand.Intn(150)),
		"南澳县": float32(rand.Intn(150)),
	}

	routers = []router{
		{"bar3D", charts.RouterOpts{URL: Host + "/bar3D", Text: "Bar3D-(3D 柱状图)"}},
		{"bar", charts.RouterOpts{URL: Host + "/bar", Text: "Bar-(柱状图)"}},
		{"boxPlot", charts.RouterOpts{URL: Host + "/boxPlot", Text: "BoxPlot-(箱线图)"}},
		{"effectScatter", charts.RouterOpts{URL: Host + "/effectScatter", Text: "EffectScatter-(动态散点图)"}},
		{"funnel", charts.RouterOpts{URL: Host + "/funnel", Text: "Funnel-(漏斗图)"}},
		{"gauge", charts.RouterOpts{URL: Host + "/gauge", Text: "Gauge-仪表盘"}},
		{"geo", charts.RouterOpts{URL: Host + "/geo", Text: "Geo-地理坐标系"}},
		{"graph", charts.RouterOpts{URL: Host + "/graph", Text: "Graph-关系图"}},
		{"heatMap", charts.RouterOpts{URL: Host + "/heatMap", Text: "HeatMap-热力图"}},
		{"kline", charts.RouterOpts{URL: Host + "/kline", Text: "Kline-K 线图"}},
		{"line", charts.RouterOpts{URL: Host + "/line", Text: "Line-(折线图)"}},
		{"line3D", charts.RouterOpts{URL: Host + "/line3D", Text: "Line3D-(3D 折线图)"}},
		{"liquid", charts.RouterOpts{URL: Host + "/liquid", Text: "Liquid-(水球图)"}},
		{"map", charts.RouterOpts{URL: Host + "/map", Text: "Map-(地图)"}},
		{"overlap", charts.RouterOpts{URL: Host + "/overlap", Text: "Overlap-(重叠图)"}},
		{"parallel", charts.RouterOpts{URL: Host + "/parallel", Text: "Parallel-(平行坐标系)"}},
		{"pie", charts.RouterOpts{URL: Host + "/pie", Text: "Pie-(饼图)"}},
		{"radar", charts.RouterOpts{URL: Host + "/radar", Text: "Radar-(雷达图)"}},
		{"sankey", charts.RouterOpts{URL: Host + "/sankey", Text: "Sankey-(桑基图)"}},
		{"scatter", charts.RouterOpts{URL: Host + "/scatter", Text: "Scatter-(散点图)"}},
		{"scatter3D", charts.RouterOpts{URL: Host + "/scatter3D", Text: "Scatter-(3D 散点图)"}},
		{"surface3D", charts.RouterOpts{URL: Host + "/surface3D", Text: "Surface3D-(3D 曲面图)"}},
		{"themeRiver", charts.RouterOpts{URL: Host + "/themeRiver", Text: "ThemeRiver-(主题河流图)"}},
		{"wordCloud", charts.RouterOpts{URL: Host + "/wordCloud", Text: "WordCloud-(词云图)"}},
		{"page", charts.RouterOpts{URL: Host + "/page", Text: "Page-(顺序多图)"}},
	}
)

func OrderRouters(chartType string) []charts.RouterOpts {
	for i := 0; i < len(routers); i++ {
		if routers[i].name == chartType {
			routers[i], routers[0] = routers[0], routers[i]
			break
		}
	}

	rs := make([]charts.RouterOpts, 0)
	for i := 0; i < len(routers); i++ {
		rs = append(rs, routers[i].RouterOpts)
	}
	return rs
}

func getRenderPath(f string) string {
	return path.Join("html", f)
}

func logTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func genKvData() map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(nameItems); i++ {
		m[nameItems[i]] = rand.Intn(maxNum)
	}
	return m
}
