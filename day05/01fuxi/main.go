package main

import (
	"encoding/json"
	"fmt"
)

// 结构体

type temp struct {
	x int
	y int
}

type person struct {
	name string
	age  int
}

// 普通有返回函数
// 返回值是对应的int类型
func sum(x int, y int) int {
	return x + y
}

// 构造(结构体变量的)函数: 由于每次都实例化结构体麻烦,故封装到函数里面,叫构造函数
// 返回值是对应的结构体类型
func newPerson(n string, i int) person {
	return person{
		name: n,
		age:  i,
	}
}

// 方法
// 接受者使用对应类型的首字母小写
// 指定了接受者之后,只有接受者这个类型的变量才能调用这个方法
func (p person) dream(str string) {
	fmt.Printf("%s的梦想是学好go语言,然后%s", p.name, str)
}

// func (p person) guonian() {
// 	p.age++ // 此处是p1的副本,改的是副本
// }

// 指针接受者
// 1. 需要修改结构体变量的值时需要用指针接受者
// 2. 结构体本身比较大,拷贝的内存开销比较大时也要使用指针接受者
// 3. 保持一致性:如果一个方法使用了指针接受者,其他的方法为了统一也要使用指针接受者
func (p *person) guonian() {
	p.age++
}

func main() {
	//普通类型
	var a = temp{
		1,
		2,
	}
	// 匿名结构体
	var b = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(a)
	fmt.Println(b)

	// 对比记忆 基本数据类型和自定义数据类型
	// 基本数据类型使用
	var x int
	y := int8(18)
	fmt.Println(x, y)

	// 自定义数据类型使用
	var p1 person // 结构体实例化1 :一个一个调用传值
	p1.name = "周林"
	p1.age = 90000
	p2 := person{"保德路", 18}   // 结构体实例化2 : 用{} 一次写完(少可以用)
	p3 := newPerson("周林", 13) // 结构体实例化3 : 用构造函数直接用完! (多次使用推荐)
	fmt.Println(p1, p2, p3)
	p1.dream("好好学习")
	p2.dream("天天向上")

	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)

	// 结构体嵌套
	type addr struct {
		province string
		city     string
	}
	type student struct {
		name string
		addr // 匿名嵌套别的结构体,就使用类型名做名称
	}

	type point struct {
		X int `json:"zhoulin"` // 给json使用的名字
		Y int `json:"baodelu"`
	}

	// 序列化
	q1 := point{100, 12}
	b1, err := json.Marshal(q1)
	// 如果出错了
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Println(string(b1))

	// 反序列化: 有字符串---->Go中的结构体变量
	str1 := `{"zhoulin":10010,"baodelu":10086}`
	var po2 point                            // 造一个结构体变量,准备接受反序列化的值
	err = json.Unmarshal([]byte(str1), &po2) // 需要传入个指针,修改值
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
	}
	fmt.Println(po2)
}
