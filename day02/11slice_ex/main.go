package main

import "fmt"

//练习题  面试题
func main() {
	var a = make([]string, 5, 10) //创建切片,长度为5,容量为10
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)      // [     0 1 2 3 4 5 6 7 8 9]
	fmt.Println(cap(a)) // 20
}
