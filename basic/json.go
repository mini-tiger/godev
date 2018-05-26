package main

import (
	"encoding/json"
	"fmt"
)

// Product 商品信息
type Product struct {
	Name      string `json:"name,omitempty"`       //tag
	ProductID int64  `json:"product_id,omitempty"` //omitempy，可以在序列化的时候忽略0值或者空值
	Number    int    `json:"-"`                    //tag  标签，json不序列化 此项
	Price     float64
	IsOnSale  bool `json:"is_on_sale,string"` // 序列化转换 字符串
	Bt        bool //首字母必须大写
}

func tjson() {

}

func main() {

	p := &Product{}
	(*p).Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 0
	p.Bt = false

	data, err := json.Marshal(p)
	fmt.Println(string(data), err) //"Price":2499,"is_on_sale":"true","Bt":false} <nil>
	fmt.Printf("%T\n", data)

	data1, _ := json.MarshalIndent(p, "[+]", "   ") //可读性, [+] 前缀  “   ”格式
	fmt.Printf("%s ,%[1]T \n", data1)
	/*
		{
		[+]   "Price": 2499,
		[+]   "is_on_sale": "true",
		[+]   "Bt": false
		[+]},[]uint8
	*/

	fmt.Printf("%s \n", "__________________________")
	shuzu() //
	fmt.Printf("%s \n", "__________________________")
	jiexi(&data)
}
func shuzu() {

	pp := []Product{} //初始化数组
	p := &(Product{}) // 初始化 结构体

	for i := 0; i < 5; i++ {
		p = &Product{} //重置结构体内存地址

		p.Name = "Xiao mi 6"
		p.IsOnSale = true
		p.Number = 10000
		p.Price = 2499.00 + float64(i)
		p.ProductID = int64(i)
		p.Bt = false
		pp = append(pp, *p) //数组每个元素 是结构体
	}

	fmt.Printf("%T \n %[1]v\n", pp)
	/*
	   []main.Product
	    [{Xiao mi 6 0 10000 2499 true false} {Xiao mi 6 1 10000 2499 true false} {Xiao mi 6 2 10000 2499 true false} {Xiao mi 6 3 10000 2499 true false} {Xiao mi 6 4 10000 2499 true false}]
	*/

	data, err := json.Marshal(pp)
	fmt.Printf("%s %[1]T\n", data, err)

	jiexi_shuzu(&data) //任意json解析
}

func jiexi_shuzu(j *[]uint8) {
	fmt.Printf("%50s \n", "__________________________任意json解析")

	var dat []map[string]interface{}

	if err := json.Unmarshal([]byte(*j), &dat); err == nil {
		fmt.Println(dat) //slice 每项是map
		fmt.Println(dat[0]["name"])
	} else {
		fmt.Println(err)
	}
}

func jiexi(j *[]uint8) {
	p := &Product{} // 关联 struct

	e := json.Unmarshal([]byte(*j), p) //解析

	fmt.Println(e)  // nil
	fmt.Println(*p) //{Xiao mi 6 0 10000 2499 true false}

	//重新定义结构体 提取指定项
	type Product_min struct {
		Name string `json:"name,omitempty"`
		// ProductID int64  `json:"product_id,omitempty"` //omitempy，可以在序列化的时候忽略0值或者空值
		// Number    int    `json:"-"`                    //tag  标签，json不序列化 此项
		// Price     float64
		// IsOnSale  bool `json:"is_on_sale,string"` // 序列化转换 字符串
		// Bt        bool //首字母必须大写
	}
	pp := &Product_min{}                 // 关联 struct
	ee := json.Unmarshal([]byte(*j), pp) //解析

	fmt.Println(ee)  // nil
	fmt.Println(*pp) //{Xiao mi 6 0 10000 2499 true false}
}
