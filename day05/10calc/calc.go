package calc

import "fmt"

// 包中的标识符(变量名\函数体\结构体\接口等)如果有首字母小写的,表示私有(只能在当前包中使用)
// 首字母大写的标识符可以被外部调用
func init() {
	fmt.Println("improt包时候自动执行")
}

func Add(x, y int) int {
	return x + y
}
