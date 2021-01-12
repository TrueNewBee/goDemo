package main

import "fmt"

// 结构体遇到的问题

// 1. myInt(100) 是这个啥
// 这是一个强制类型转换
type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

type person struct {
	name string
	age  int
}

func main() {
	// 声明一个int32类型的变量x, 它的值是10
	// 方法1:
	// var x int32
	// x = 10
	// 方法2:
	// var x int32 = 10
	// 方法3:
	// var x = int32(10)
	// 方法4:
	// x := int32(10)
	// fmt.Println(x)

	// 声明一个myInt类型的变量m,它的值是100
	// 方法1:
	// var m int32
	// m = 10
	// 方法2:
	// var m int32 = 10
	// 方法3:
	// var m = int32(10) // 强制类型转换
	// 方法4:
	// m := int32(10)
	// fmt.Println(m)

	// 问题2 : 结构体初始化
	// 方法1:
	var p person // 声明一个person类型的变量p
	p.name = "元帅"
	p.age = 18
	fmt.Println(p)
	var p1 person
	p1.name = "周林"
	p1.age = 19
	fmt.Println(p1)
	// 方法2:
	// 键值对初始化, 声明的同时初始化
	var p2 = person{
		name: "冠华",
		age:  15,
	}
	fmt.Println(p2)
	// 值列表初始化
	var p3 = person{
		"理想",
		18,
	}
	fmt.Println(p3)

}

// 问题3: 为什么要有构造函数:别人调用我,我能给她一个特定类型的变量
func newPerson(name string, age int) person {
	// 别人调用我,我能给她一个person类型的变量
	return person{
		name: name,
		age:  age,
	}
}

// 返回一个指针类型
// func newPerson(name string, age int) *person {
// 	// 别人调用我,我能给她一个person类型的变量
// 	return &person{
// 		name: name,
// 		age:  age,
// 	}
// }

// 2021-01-11 23:07:12 键盘到了 用着很爽,茶轴!
// 明天开始步入正轨!!!!
