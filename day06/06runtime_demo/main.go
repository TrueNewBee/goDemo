package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller()

func f1() {
	// 传入0 当前函数名
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)            // 文件名字
	fmt.Println(path.Base(file)) // main.go
	fmt.Println(line)
}
func main() {
	f1()
}
