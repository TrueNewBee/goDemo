package main

import "fmt"

// 面试题

// Go语言中的函数的return不是原子操作,在底层分为两步来执行的
// 第一步: 返回值赋值
// defer
// 第二步: 真正的RET返回
// 函数中如果存在defer, 那么defer执行的时机实在第一步和第二步之间

// 只有传参是修改拷贝副本,不传参和传入指针都是直接修改返回值

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x不是返回值
	}()
	return x
}

// 没有传参,故返回值并没有拷贝,直接进行了修改
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值=x
}

// 传参,故返回值拷贝,修改了副本而不是返回值
func f3() (y int) {
	x := 5
	defer func() {
		x++ // 修改的是x 而不是返回值
	}()
	return x // 返回值 = y = x =5 2.defer修改的是x   3.真正的返回
}

// 传参,故返回值拷贝,修改了副本而不是返回值
func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中的副本
	}(x)
	return 5 // 返回值 = x = 5
}

// 匿名函数
func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5 // 返回值 = x = 5
}

// 传一个x的指针的匿名函数, 直接修改的是指针地址
func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 // 返回值 = x = 5  2.defer x=6   3.RET返回
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f5()) // 5
}
