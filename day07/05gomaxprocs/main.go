package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS
var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1) // 默认CPU的逻辑核心数, 默认跑满CPU
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	// time.Sleep(time.Second)
	wg.Wait()
}
