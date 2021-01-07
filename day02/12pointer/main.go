package main

import "fmt"

// pointer

func main() {
	// 1. &: 取地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p) // *int: int类型的指针
	// 2.* :根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m) // int

	// var a *int // nil pointer
	var a = new(int) // new函数申请了一个内存地址
	*a = 100
	fmt.Println(*a)

}
