# GoS5 day04

# 内容回顾

**函数**

函数定义

```go
func  函数名(参数1 ,参数2...)返回值{
	函数体
}
```



**函数进阶**

​	内置函数 : `panic`和 `recover`

​	defer: 延迟调用,多用于处理资源释放

​	高阶函数 : 函数可以作为参数也可以作为返回值

​	闭包:  函数和其外部变量的引用  // 闭包是一个函数,这个函数包含了他外部作用域的一个变量

```go
// 闭包示例 把yuanshuai 传到 low 里面
func yuanshuai(name string) {
	fmt.Println("hello", name)
}

func low(f func()) {
	f()
}

// 闭包
func bi(f func(string), name string) func() {
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

```









# 今日内容

结构体(struct)   (类似于 类,在go中是类型,而在java中是class类  )

结构体都是值类型

构造函数: 返回一个结构体变量的函数





方法

[在Go中恰到好处的内存对齐](https://segmentfault.com/a/1190000017527311?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com)