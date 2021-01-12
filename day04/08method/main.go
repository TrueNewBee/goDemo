package main

import "fmt"

// 方法
// 标识符: 变量名 函数名 类型名 方法名
// Go语言中如果标识符首字母是大写的,就表示对外部课件(暴露的,共有的)
// 小写类似于 private 如果要是大写,则写上注释

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量,多用类型名
// 类似于java中的 类方法

type dog struct {
	name string
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

type person struct {
	age int
}

// 构造函数
func newPerson(age int) person {
	return person{
		age: age,
	}
}

// 方法
func (d dog) wang() { // 用类型首字母小写, 类似于Python中的 self
	fmt.Printf("%s:汪汪汪", d.name)
}

// 使用值接收者: 传拷贝进去
func (p person) guoninan() {
	p.age++
}

// 指针接收者,传内存地址进去
func (p *person) zhenguoninan() {
	p.age++
}

func (p *person) dream() {
	fmt.Println("不上班也能挣钱")
}

func main() {
	d1 := newDog("zhoulin")
	d1.wang()
	p1 := newPerson(18)
	fmt.Println(p1.age) // 18
	p1.guoninan()
	fmt.Println(p1.age) // 18   通过值接收,传拷贝 ,仅修改副本,不修改原始值
	p1.zhenguoninan()
	fmt.Println(p1.age) // 19	通过指针接收,直接用地址,修改接受者中的值
}
