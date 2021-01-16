package main

import "fmt"

// 使用值接受者和指针接受者的区别
type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 使用了值接受者实现了接口的所有方法
// func (c cat) move() {
// 	fmt.Println("走猫步```")
// }
// func (c cat) eat(food string) {
// 	fmt.Printf("猫吃%s```", food)
// }

// 使用指针接受者实现了接口的所有方法
func (c *cat) move() {
	fmt.Println("走猫步```")
}
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s```", food)
}

func main() {
	var a1 animal

	// c1 := cat{"tom", 4}  //cat 值类型
	c2 := &cat{"假老练", 4} // *cat	指针类型

	// a1 = c1 //	实现animal这个接口的是cat的指针类型 报错
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
