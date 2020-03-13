package main

import (
	"fmt"
	"github.com/coocood/freecache"
	"strconv"
)

//doc : https://pkg.go.dev/github.com/coocood/freecache?tab=doc
//download : https://github.com/coocood/freecache

func main() {
	cacheSize := 50 * 1024 * 1024 // 最小512K   512 * 1024
	cache := freecache.NewCache(cacheSize)

	for i := 0; i < 65535*10; i++ {
		key := []byte("abc" + strconv.Itoa(i))
		val := []byte("def")
		expire := 60 // expire in 60 seconds
		e := cache.Set(key, val, expire)
		if e != nil {
			fmt.Println("set " + e.Error())
		}

		//affected := cache.Del(key)
		//fmt.Println("deleted key ", affected)

	}

	fmt.Println("entry count ", cache.EntryCount())
	fmt.Println("EvacuateCount count ", cache.EvacuateCount())
	//got, err := cache.Get(key)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(string(got))
	//}

}
