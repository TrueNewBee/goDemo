package main

import (
	"fmt"
	"sync"
	"time"
)

// rwlock

var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex   // 互斥锁
	rwLock sync.RWMutex // 读写锁
)

// 读操作
func read() {
	defer wg.Done()
	// lock.Lock()
	rwLock.RLock() // 读写锁
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwLock.RUnlock()
}

// 写操作
func write() {
	defer wg.Done()
	// lock.Lock()
	rwLock.RLock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	// lock.Unlock()
	rwLock.RUnlock()
}
func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	fmt.Println(time.Now().Sub(start))
}
