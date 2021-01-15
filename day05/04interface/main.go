package main

import "fmt"

// 接口实例2
// 不管什么牌子的车,都能跑

// 定义一个car接口类型
// 不管是什么结构体 只要有run方法都能是car类型
type car interface {
	run()
}
type baoshijie struct {
	brand string
}

func (b baoshijie) run() {
	fmt.Printf("%s速度700迈", b.brand)
}

type falali struct {
	brand string
}

func (f falali) run() {
	fmt.Printf("%s速度700迈", f.brand)
}

func drive(c car) {
	c.run()
}

// drive函数接收一个car类型的变量
func main() {
	var f1 = falali{
		brand: "法拉利",
	}

	var b1 = baoshijie{
		brand: "保时捷",
	}
	drive(f1)
	drive(b1)

}
