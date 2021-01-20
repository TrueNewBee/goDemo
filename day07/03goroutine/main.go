package main

import (
	"fmt"
	"time"
)

// goroutine

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动只有会创建一个主goroutine去执行
func main() {
	for i := 0; i < 1000; i++ {
		go hello(i) // 开启一个单独的goroutine去执行hello函数(任务)
		// 匿名函数
		go func(i int) {
			fmt.Println(i) //	用的是参数的的那个i而不是外面的i
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
}
