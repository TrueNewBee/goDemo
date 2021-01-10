package main

import "fmt"

// 自定义类型和类型别名

// type 后面跟的是类型

type myInt int     // 自定义类型
type yourInt = int // 类别名

func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m yourInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune // rune是 int32的别名
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
