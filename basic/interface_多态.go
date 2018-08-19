package main

import (
	"fmt"
	"os"
)

type people_money interface {
	money() float64
	cook_name() string
}

type fish struct {
	name         string
	dazhe        float64
	jiage        float64
	maichu_money float64
}

type milk struct {
	name         string
	dazhe        float64
	jiage        float64
	maichu_money float64
}

type water struct {
	name         string
	maichu_money float64
}

func ABc() {
	os.Open()
}

//fish 绑定的方法
func (self *fish) money() float64 {
	return self.maichu_money - self.jiage*self.dazhe
}
func (self *fish) cook_name() string {
	return self.name
}

//milk 绑定的方法
func (self *milk) money() float64 {
	return self.maichu_money - self.jiage*self.dazhe
}
func (self *milk) cook_name() string {
	return self.name
}

//water 绑定的方法
func (self *water) money() float64 {
	return self.maichu_money
}
func (self *water) cook_name() string {
	return self.name
}

func main() {
	//无论 任何结构体不是 其它类型，只有满足接口的方法，就能统一处理计算
	var f1 *fish = &fish{"鲤鱼", 0.8, 15, 70}
	var f2 *fish = &fish{"鲨鱼", 0.95, 20, 200}
	var m1 *milk = &milk{"牛奶", 0.8, 12, 16}
	var m2 *milk = &milk{"酸奶", 0.65, 10, 13}
	var w1 *water = &water{"自来水", 2}
	var w2 *water = &water{"矿泉水", 6}

	p := []people_money{f1, f2, m1, m2, w1, w2}
	fmt.Println(return_totle_money(&p))
}

func return_totle_money(p *[]people_money) (total_money float64) {

	for _, value := range *p {
		fmt.Printf("cookname:%s, money:%0.2f\n", value.cook_name(), value.money())
		total_money += value.money()
	}
	return
}
