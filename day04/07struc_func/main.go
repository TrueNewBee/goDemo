package main

import "fmt"

// 构造函数 : 返回一个结构体变量的函数

type person struct {
	name string
	age  int
}

// 构造函数: 约定成俗用 new 开头
// 返回的是结构体还是结构体指针
//当结构体比较大的时候尽量使用结构体指针,减少程序的开销

// 返回的是结构体
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

// 返回的是结构体指针
func newPersonPoint(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("元帅", 18)
	p2 := newPerson("理想", 1800)
	fmt.Println(p1, p2)
}
