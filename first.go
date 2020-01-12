package main

import (
	"fmt"
	"reflect"
)


type A struct {
	D map[int]int
}

func (a *A)haha()  {
	fmt.Println(a.D)
}


type II interface {
	haha()
}

func tttt(i II)  {
	i.haha()
}
func main()  {
	var a []int = make([]int,0,1<<13)
	fmt.Printf("%p,len:%d,cap:%d\n",a,len(a),cap(a))

	for i:=0;i<1<<13;i++{
		a=append(a,i)
		//a[i]=ir
	}

	fmt.Printf("%p,len:%d,cap:%d\n",a,len(a),cap(a))
	var b []int = make([]int,len(a),len(a))
	copy(b,a)
	fmt.Printf("b:%p,len:%d,cap:%d\n",b,len(b),cap(b))


	a=nil
	fmt.Printf("%p,len:%d,cap:%d\n",a,len(a),cap(a))
	a=append(a,1)
	a[0]=1
	fmt.Printf("%p,len:%d,cap:%d\n",a,len(a),cap(a))


	aa:=A{}
	aa.D=make(map[int]int,0)
	aa.D[1]=1
	aa.haha()
	//var aai II
	//aai := &aa
	tttt(&aa)

	var ddda *string
	ddda="!"
	//var DetailFieldsMap map[int]string = map[int]string{0: "DATACLIENT", 1: "AgentInstance", 2: "BackupSetSubclient", 3: "Job ID (CommCell)(Status)",
	//	4: "Type", 5: "Scan Type", 6: "Start Time(Write Start Time)", 7: "End Time or Current Phase", 8: "Size of Application", 9: "Data Transferred",
	//	10: "Data Written", 11: "Data Size Change", 12: "Transfer Time", 13: "Throughput (GB/Hour)", 14: "Protected Objects", 15: "Failed Objects",
	//	16: "Failed Folders"}
	//fmt.Println(len(DetailFieldsMap))
	//fmt.Printf("%p\n",DetailFieldsMap)
	//var dda map[int]string
	//
	//dda=DetailFieldsMap
	//fmt.Printf("%p\n",dda)


}
func Clear(v interface{}) { // 必须传入指针
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}