package main

import "fmt"

// 结构体是值类型

// 仿佛是个类
type person struct {
	name   string
	age    int
	gender string
	// hobby  []string
}

// go语言中函数传参永远是拷贝
func f(x person) {
	x.gender = "女" // 修改的是副本的gender
}

func f2(x *person) {
	// (*x).gender = "女" // 根据内存地址找到哪个原变量,修改的就是原来的变量
	x.gender = "女" // 语法塘, 自动根据指针找对应的变量
}

func main() {
	var p person
	f2(&p) // 传入个指针

	// 结构体指针1
	var p2 = new(person) // 声明一个变量,然后初始化
	p2.name = "理想"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)  // p2保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2) // 求b2的内存地址

	// 2.结构体指针2
	// 2.1 key - value初始化
	var p3 = person{
		name:   "元帅",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)

	// 2.2 使用值列表的形式初始化,值的顺序要和结构体定义时字段的顺序一致
	p4 := person{
		"小王子",
		16,
		"22",
	}
	fmt.Printf("%#v\n", p4)

	// // 声明一个person类型的变量p
	// var p person
	// // 通过字段赋值
	// p.name = "周林"
	// p.age = 9000
	// p.gender = "男"
	// p.hobby = []string{"篮球", "足球", "双色球"}
	// fmt.Println(p)
	// // 访问变量p的字段
	// fmt.Println(p.name)
	// // 仿佛用类 new个对象 在go中就是 类型而已
	// var p2 person
	// p2.name = "理想"
	// p2.age = 18
	// fmt.Printf("type:%T  value: %v\n", p2, p2)

	// // 匿名结构体 (类似于java中的匿名内部类)
	// // 多用于 临时的chengjing
	// var s struct {
	// 	x string
	// 	y int
	// }
	// s.x = "嘿嘿嘿"
	// s.y = 100
	// fmt.Printf("type:%T value:%v\n", s, s)

}
