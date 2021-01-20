package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

// rand随机数
func f() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		r1 := rand.Int()    // int64
		r2 := rand.Intn(10) // 0<= x <10
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wg.Done() // 完成一次goroutine就减1
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

// goroutine 计数器 多个goroutine的同步
var wg sync.WaitGroup

func main() {
	// f()
	// wg.Add(10)
	for i := 0; i < 10; i++ {
		wg.Add(1) // 开始一次goroutine就加1
		go f1(i)
	}
	wg.Wait() // 等待wg的计数器减为0
}
