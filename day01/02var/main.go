package main

import "fmt"

// Go语言中,推荐使用驼峰式命名  首字母小写
var studentName string

var student_name string
var StudentName string

// 声明变量
// var name  string
// var age int
// var isOk bool

// 批量声明
var (
	name string // ""
	age  int    // 0
	isOk bool   // flase
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	// Go语言中非全局变量声明必须使用.不适用编译不过去

	fmt.Print(isOk) // 在终端中输出打印的内容
	fmt.Println()
	fmt.Printf("name: %s", name) //  &s:站位符 使用name这个变量的值去替换占位符
	fmt.Println(age)             // 打印完指定的内容之后会在后面加一个换行符

	// 声明变量同时赋值(不推荐)
	var s1 string = "who"
	fmt.Println(s1)

	// 类推导(根据值判断该变量是什么类型)
	var s2 = "20"
	fmt.Println(s2)

	// 简短变量声明,只能在函数中使用(推荐)
	s3 := "哈哈哈哈"
	fmt.Println(s3)

	// _ 为匿名变量,像一个树洞,表示忽略值
	// 函数外每个语句都必须以关键字开始
	// 同一个作用域中不能重复声明同名的变量
}
