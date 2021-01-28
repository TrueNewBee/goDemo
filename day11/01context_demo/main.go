package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要

var wg sync.WaitGroup
var notify bool

func f() {
	defer wg.Done()
	for {
		fmt.Println("周林")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	// 通知子goroutine退出
	notify = true
	wg.Wait()
}
