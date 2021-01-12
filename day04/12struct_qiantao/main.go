package main

import "fmt"

// 嵌套 把变量重复部分抽出
// 这tm就是 json !

//
type address struct {
	province string
	city     string
}
type workPlace struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	// addr address
	address //匿名嵌套结构体
	workPlace
}

type compay struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "周林",
		age:  9000,
		address: address{
			province: "山东",
			city:     "威海",
		},
		workPlace: workPlace{
			province: "杭州",
			city:     "威海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)
	// fmt.Println(p1.city) // 现在自己结构体找这个字段,找不到就去嵌套的结构体中查找该字段
	// 当有多个结构体有匿名字段是,就老老实实用这种调用↓↓
	fmt.Println(p1.name, p1.workPlace.city)
}
