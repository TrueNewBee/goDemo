package split_string

import (
	"strings"
)

// Split切割字符串
// example
// abc, b => [a c]

// Split  实现按字符切片功能
func Split(str string, sep string) []string {
	// str: "babcbef"  spe = "b"  [a cdef]
	// 在函数外初始化,仅申请一次内存
	var ret = make([]string, 0, strings.Count(str, sep)+1)
	// 找到目标字符的位置
	index := strings.Index(str, sep)
	for index >= 0 {
		// 把切片追加进去
		ret = append(ret, str[:index]) // [ ) 不包含b
		// [)从b后面一个字符,到后面所有 , 跳过了b
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}

//Fib 斐波那契函数
// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
