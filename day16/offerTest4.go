package main

import (
	"math/rand"
	"sync"
)

//const n = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	// 加上锁,解决并发
	flag := sync.Mutex{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			m[rand.Int()] = rand.Int()
			flag.Unlock()
		}()
	}
	wg.Wait()
	println(len(m))

}
