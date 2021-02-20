package main

import "sync"

const n = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(n)
	// 之前是for的问题,突然进去了,电脑好了就结束了
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
}
