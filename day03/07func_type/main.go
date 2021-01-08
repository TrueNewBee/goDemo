package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("hello 沙河")
}

func f2() int {
	return 10
}

// 函数也可以作为参数的类型 func()是无参的
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func ff(a, b int) int {
	return a + b
}

// 函数还可以作为返回值
func f5(x func() int) func(int, int) int {
	return ff
}

func main() {
	// func() 为函数类型
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
}
