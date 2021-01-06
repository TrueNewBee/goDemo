package main

import "fmt"

// for 循环

func main() {
	// 基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种1  省略初始语句 常用
	// var i = 5
	// for ; i < 19; i++ {
	// 	fmt.Println(i)
	// }

	// 变种2 少用
	var i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 无限循环  不要轻易用
	// for {
	// 	fmt.Println("123")
	// }

	// for range循环
	s := "hello 沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

}

// 99 乘法表  for 循环
// 越努力,越幸运!
