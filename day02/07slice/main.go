package main

import "fmt"

// 切片 slice

func main() {
	// 切片的定义  就是没有长度的数组
	var s1 []int    // 定义一个存放int类型的切片
	var s2 []string // 定义一个存放string类型的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) // true
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"Java", "Python", "go"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil) // false
	// 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 2. 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] //	基于一个数组切割,左包含又不包含
	s4 := a1[1:6] //	基于一个数组切割,左包含又不包含
	s5 := a1[:5]  // 从第一个 0 开始切
	s6 := a1[:]   // 从第一个切到最后一个
	fmt.Println(s3, s4, s5, s6)
	// 切片的容量是指底层数组的容量
	fmt.Printf("len(s3):%d cap(s3):%d\n", len(s3), cap(s3))
	// 切片的容量就是底层数组从切片的第一个元素到最后一个元素的数量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	// 切片再切片
	s8 := s6[3:]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))
	// 切片是一个引用类型,指向了底层的一个数组
	fmt.Println("s6:", s6)
	a1[6] = 1300 // 修改了底层数组的值
	fmt.Println("s6:", s6)
	fmt.Println("s8:", s8)

}
