package main

import "fmt"

// switch
// 简化大量判断(一个变量和具体的值作比较)

func main() {
	var n = 5
	if n == 1 {
		fmt.Println("1")
	} else if n == 2 {
		fmt.Println("2")
	} else if n == 3 {
		fmt.Println("3")
	} else if n == 4 {
		fmt.Println("4")
	} else if n == 5 {
		fmt.Println("5")
	} else {
		fmt.Println("无效数字")
	}
}
