package main

import (
	"fmt"
	"strings"
)

const l =4

func main()  {
	var arr [l]string = [l]string{"abc","bcd"}
	fmt.Printf("arr 的值: %v ,地址： %p \n",arr,&arr)
	upper11(&arr)//指向数组的指针
	fmt.Printf("经过upper11 method 后，arr 的值: %v ,地址： %p \n",arr,&arr)

	fmt.Println("______________________________________")
	fmt.Printf("arr_p 的值: \n")
	var arr_p [l]*string
	for i,_:=range arr { //生成app_p
		arr_p[i] = &(arr[i])
		fmt.Println(*(arr_p[i]))
	}

	lower(arr_p) //指针数组
	fmt.Printf("经过lower method后,arr_p 的值: %v ,地址： %x \n",arr_p,arr_p)
	for _,v:=range arr_p{
		fmt.Println(*v)
	}
	}

//指向数组的指针
func upper11(arr *[l]string)  (*[l]string){
		for i:=0;i<len(*arr);i++{
			(*arr)[i]=strings.ToUpper((*arr)[i])
			//fmt.Printf("type: %T, 数值： %v\n",(*arr)[i],(*arr)[i])
	}
	return arr
}

//数组指针
func lower(arr [l]*string)  ([l]*string){
	for i:=0;i<len(arr);i++{
		*(arr[i])= strings.ToLower(*(arr[i]))
		//fmt.Printf("type: %T, 数值： %v\n",arr[i],arr[i])
	}
	return arr
}