package main

import "fmt"

// fmt库用法,可以参考博客或者官方文档
// https://studygolang.com/pkgdoc
func main() {
	// %T : 查看类型
	// %d : 十进制数
	// %b : 二进制数
	// %o : 八进制数
	// %x : 十六进制数
	// %c : 字符
	// %s : 字符串
	// %p : 指针
	// %v : 值
	// %f : 浮点数
	// %t : 布尔值

	// 获取用户输入
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("用户输入的内容是:", s)
	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println(name, age, class)
}
