package g

import "fmt"

func (self Fish)zhengqian() (float64) {
	return self.RMB - self.Chengben
}
func (self Milk)zhengqian() (float64) {
	return self.RMB - self.Chengben
}

func (self *Cook_food)Cooks()  {
	var total float64
	for k,v := range self.Food {
		fmt.Printf("food:%s,zhangqian:%f",k,v.zhengqian())
		total= v.zhengqian()+total
	}
}
