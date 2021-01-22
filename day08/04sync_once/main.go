package main

import (
	"sync"
)

// 借助sync.Once实现的并发安全的单例模式：

type singleton struct{}

var instance *singleton
var once sync.Once

// GetInstance ...
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
