package main

import "fmt"

// array数组练习题
// 1. 求数组[1, 3, 5, 7, 8]所有元素的和
// 2. 找出数组中和为指定值得两个元素的下标,比如上述数组找出和为8的两个元素的下标分别为

func main() {
	// // 1. 求数组[1, 3, 5, 7, 8]所有元素的和
	a1 := [...]int{1, 3, 5, 7, 8}
	// // 把数组中的每一个元素遍历出来,累加到一个sum变量中就可以了
	// sum := 0
	// for _, v := range a1 {
	// 	sum = sum + v
	// }
	// fmt.Println(sum)

	// 2.找出和为8的下标
	// 两层for循环
	// 内层遍历从哪找?
	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i+a1[j]] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
	// for a, b1 := range a1 {
	// 	for c, b2 := range a1 {
	// 		if b1+b2 == 8 {
	// 			fmt.Printf("(%d , %d)", a, c)
	// 		}
	// 	}
	// }
}
