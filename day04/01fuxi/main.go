package main

import "fmt"

// 函数作为参数
func lixiang(f func(string), name string) {
	f(name)
}

// 函数作为返回值
func zhoulin() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

// 闭包示例 把yuanshuai 传到 low 里面
func yuanshuai(name string) {
	fmt.Println("hello", name)
}

func low(f func()) {
	f()
}

// 闭包
func bi(f func(string), name string) func() {
	// 把当前函数修饰成闭包符合了目标函数要求,间接传给了目标函数
	return func() {
		f(name)
	}
}

func main() {
	lixiang(yuanshuai, "理想")
	// 闭包示例
	fc := bi(yuanshuai, "元帅")
	low(fc)
}
