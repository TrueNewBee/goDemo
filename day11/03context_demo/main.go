package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么需要context  超时解决问题
// 多个goroutine可以同时推退出

var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("周林")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func f(ctx context.Context) {
	go f2(ctx)
	defer wg.Done()
LOOP:
	for {
		fmt.Println("周林")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	// 通知子goroutine退出
	cancel()
	wg.Wait()
}
