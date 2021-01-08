package main

import "fmt"

// defer

// defer 多用于函数结束之前释放资源(文件句柄,数据库连接,socker链接)
func deferdemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿")  // defer把它后面的语句延迟到函数返回的时候再执行
	defer fmt.Println("end") // 一个函数中可以有多个defer语句
	defer fmt.Println("end") // 多个defer语言按照先进后出(后劲先出)的顺序延迟执行
	fmt.Println("111")
}

func main() {
	deferdemo()
}
