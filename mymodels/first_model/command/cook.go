package command

import "godev/mymodels/first_model/g"




func Cook_being()  {
	cook_food_all :=make(map[string]g.Food)
	var cook_food g.Food
	cook_food=&g.Fish{"鲤鱼",100,10}
	cook_food_all.Food["鲤鱼"]=cook_food
	var cook_food1 g.Food
	cook_food1=&g.Fish{"鲟鱼",200,40}
	cook_food_all.Food["鲟鱼"]=cook_food1
}
