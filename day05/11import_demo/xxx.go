package main

import (
	// 包名有数字 就起了个别名
	"fmt"

	calc "chentianxiang.vip/studygo/day05/10calc"
)

func main() {
	ret := calc.Add(1, 2)
	fmt.Println(ret)
}
