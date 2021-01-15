package main

import "fmt"

// 引出接口的实例

// 定义一个能叫的类型
type speaker interface {
	// 只要实现了speak方法的变量都是spealer类型
	// 方法可以有多个
	speak()
}

type cat struct{}

type person struct{}

type dog struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~~")
}

func (p person) speak() {
	fmt.Println("啊啊啊~~")
}
func da(x speaker) {
	// 接受一个参数,传进什么,我就打什么
	x.speak() //	挨打了就要叫
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)
}
