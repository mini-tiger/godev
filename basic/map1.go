package main

import "fmt"

func main()  {

	var map1 map[string]string

	fmt.Println(map1)
	var map2 map[string]string = make(map[string]string)
	fmt.Println(map2)
	//增加
	map2[string("a")]= "a"
	fmt.Println(map2)
	//更改
	map2["a"]="b"
	fmt.Println(map2)
	//删除
	delete(map2,"a")
	fmt.Println(map2)
	//遍历
	map3 := map[string]string{
		"a":"aa",
		"b":"bb",
		"c":"cc",
	}
	for k,v:=range map3{
		fmt.Println(k,v)
	}
	//判断
	if k,ok:=map3["a"];ok{
		fmt.Println(map3[k])
	}
	if k,ok:=map3["aaaaa"];ok{
		fmt.Println(k)
	}
}
