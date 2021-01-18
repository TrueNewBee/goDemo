package main

import "fmt"

// 类型断言

func main() {
	var a interface{} // 定义一个空接口变量a

	a = 100
	// 如何判断a保存的值的具体类型是什么?
	// 类型断言
	// 1. x.(T)

	v1, ok := a.(int8)
	if ok {
		fmt.Println("猜对了,a是int8", v1)
	} else {
		fmt.Println("猜错了,不是int8")
	}

	// 2. switch
	switch v2 := a.(type) {
	case int8:
		fmt.Println(v2)
	case int16:
		fmt.Println(v2)
	case int32:
		fmt.Println(v2)
	case string:
		fmt.Println(v2)
	default:
		fmt.Println("滚")
	}

}
