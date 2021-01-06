package main

import "fmt"

// 常量

const pi = 3.1415926

// 批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// 批量声明变量时,如果某一行后没有赋值,默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota  面试注意: iota是go语言的常量计数器,只能在常量表达式中使用
// 每新增一行常量, iota +1
const (
	a1 = iota // 0
	a2 = iota // 1
	a3 = iota //2
)

// 跳过
const (
	b1 = iota // 0
	b2 = iota // 1
	_  = iota // 2
	b3 = iota // 3
)

// 插队
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4        // 3
)

// 多个变量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // 	d1 : 1  d2 : 2
	d3, d4 = iota + 1, iota + 2 //	d3 : 2 	d4 : 3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n1:", n2)
	fmt.Println("n1:", n3)
	fmt.Println("n1:", a1)
	fmt.Println("n1:", a2)
	fmt.Println("n1:", a3)
}
