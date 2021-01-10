package main

import "fmt"

// 上台阶的面试题

// 递归: 函数自己调用自己
// 递归适合处理那种问题相同\问题的规模越来越小的场景
// 递归一定要有一个明确的退出条件
// 永远不要高估自己!

// 上台阶的面试题
// n个台阶,一次可以走1步,也可以走2步,有多少种走法
func tiajie(n uint64) uint64 {
	if n == 1 {
		// 如果只有一个台阶就一种走法
		return 1
	}
	if n == 2 {
		return 2
	}
	// 剩下的走到最后一步,要么剩下1个台阶或者剩下2个台阶,,故把两种方案加一块就是总共路线!
	return tiajie(n-1) + tiajie(n-2)
}

// 3! =3*2*1	= 3*2!
// 4! = 4*3*2*1	= 4*3!

// 计算n的阶乘

func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}
func main() {
	ret := f(7)
	fmt.Println(ret)
}
