package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步```")
}
func (c cat) eat(food string) {
	fmt.Printf("猫吃%s```", food)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动```")
}
func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s```", food)
}

func main() {
	var a1 animal //	定义一个接口类型的变量
	fmt.Printf("%T\n", a1)
	blueCat := cat{ // 定义一个cat类型的变量 blueCat  其实就是创建个对象
		name: "淘气",
		feet: 4,
	}
	a1 = blueCat
	a1.eat("小黄鱼")
	fmt.Println(a1)

	kfc := chicken{
		feet: 2,
	}

	a1 = kfc
	a1.eat("鸡饲料")
	fmt.Printf("%T\n", a1)
}
