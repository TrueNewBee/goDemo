package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChanner() {
	fmt.Println(b)     //nil
	b = make(chan int) // 不带缓冲区通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台gorountine从管道b中取到了", x)
	}()

	b <- 10 // hang住了
	fmt.Println("10发送到通道b中了....")
	b = make(chan int, 16) // 带缓冲区的通道的初始化
	fmt.Println(b)
	wg.Wait()
}
func bufChannel() {
	fmt.Println(b)         //nil
	b = make(chan int, 10) // 带缓冲区通道的初始化
	b <- 10                // hang住了
	fmt.Println("10发送到通道b中了....")
	x := <-b
	fmt.Println("从管道b中取到了", x)
	close(b) // 关闭通道  其实不关闭还可以,垃圾自动回收了
}
func main() {
	bufChannel()
}
