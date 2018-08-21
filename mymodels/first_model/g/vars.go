package g

import "fmt"

type Food interface {
	Zhengqian() float64
}

type Cook_food struct {
	Food map[string]Food

}

type Cook_food1 struct {
	Food []*Fish

}

type Milk struct {
	RMB float64
	Chengben float64
}

type Fish struct {
	Name string
	RMB float64
	Chengben float64
}

func (self Fish)Zhengqian() (float64) {
	return self.RMB - self.Chengben
}
func (self Milk)Zhengqian() (float64) {
	return self.RMB - self.Chengben
}

func (self *Cook_food)Cooks() (total float64) {
	//var total float64
	for k,v := range self.Food {
		fmt.Printf("food:%s,zhangqian:%0.2f\n",k,v.Zhengqian())
		total= v.Zhengqian()+total
	}
	return
}