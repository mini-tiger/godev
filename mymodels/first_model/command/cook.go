package command

import (
	"godev/mymodels/first_model/g"
	"fmt"
)



func Cook_being()  {
	cook_food_all := map[string]g.Food{
		"鲤鱼":g.Fish{"鲤鱼",100,10},
		"鲟鱼":g.Fish{"鲟鱼",200,40},
	}
	fmt.Println(cook_food_all)
	for k,v :=range cook_food_all{
		fmt.Println(k,v.Zhengqian())
	}
	s:=g.Cook_food{cook_food_all}
	fmt.Printf("%0.2f\n",s.Cooks())
}

func create_flsh1() (*g.Fish){
	return &g.Fish{"鲤鱼",100,10}
}

func create_flsh2() (*g.Fish){
	return &g.Fish{"鲫鱼",200,20}
}
func Cook_being1()  {
	cook_food_all :=[]*g.Fish{
		create_flsh1(),
		create_flsh2(),
	}
	fmt.Println(cook_food_all)
	for k,v :=range cook_food_all{
		fmt.Println(k,v.Zhengqian())
	}

}
