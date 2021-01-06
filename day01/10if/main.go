package main

import "fmt"

// if条件判断

func main() {

	// age := 19
	// if age > 18 {
	// 	fmt.Println("澳门首家线上赌场上线了")
	// } else {
	// 	fmt.Println("改写暑假作业了")
	// }

	// 作用域
	// age变量此时只在if条件判断语句中生效, 节约内存
	if age := 19; age > 18 {
		fmt.Println("澳门首家线上赌场上线了")
	} else if age < 1 {
		fmt.Println("改写暑假作业了")
	} else {
		fmt.Println("玩吧")
	}
	// fmt.Println(age)

}
