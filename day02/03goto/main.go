package main

import "fmt"

// goto

func main() {
	// 推荐使用
	// 如何跳出多层循环
	var flag =false
	for i := 0, i <10; i++{
		for j :='A'; j <'Z'; j++{
			if j == 'C'{
				flag = true
				break	//跳出内循环
			}
			fmt.Println("%v-%c\n", i, j)
		}
		if flag{
			break 	// 跳出for循环(外层的for循环)
		}
	}

	// 不推荐使用
	// goto+label实现跳过多层for循环
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto xx // 跳到我指定的那个标签
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
xx: // 标签
	fmt.Println("结束for循环")
}
