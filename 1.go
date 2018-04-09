package main

import (
	"fmt"
)

//飞的接口
type IFly interface {
	Fly()
}

//吃的接口
type IEat interface {
	Eat()
}

//跑的接口
type IRun interface {
	Run()
}

//狗的实现类
type Dog struct {
	name string
}

func (dog Dog) Eat() {
	fmt.Println(dog, " eat")
}

func (dog *Dog) Run() {
	fmt.Println(dog.name + " run")
}

//鸟的实现类
type Bird struct {
	name string
}

func (bird *Bird) Fly() {
	fmt.Println(bird.name + " fly")
}

func (bird *Bird) Eat() {
	fmt.Println(bird.name + " eat")
}

func main() {
	var d = Dog{"Dog"}
	fmt.Println(d.name)
	d.Eat()

}
