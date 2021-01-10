package main

import "fmt"

// 结构体

// 仿佛是个类
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明一个person类型的变量p
	var p person
	// 通过字段赋值
	p.name = "周林"
	p.age = 9000
	p.gender = "男"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)
	// 访问变量p的字段
	fmt.Println(p.name)
	// 仿佛用类 new个对象 在go中就是 类型而已
	var p2 person
	p2.name = "理想"
	p2.age = 18
	fmt.Printf("type:%T  value: %v\n", p2, p2)

	// 匿名结构体 (类似于java中的匿名内部类)
	// 多用于 临时的chengjing
	var s struct {
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)

}
