package d

import (
	_ "net/http/pprof"

	"time"
	"testing"
)
var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)
	time.Sleep(1*time.Second)

	return sData
}
const url = "https://github.com/EDDYCJY"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}

//func main() {
//	go func() {
//		for {
//			log.Println(Add("https://github.com/EDDYCJY"))
//		}
//	}()
//
//	http.ListenAndServe("0.0.0.0:6060", nil)
//}