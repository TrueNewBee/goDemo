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

```go
// 标识符: 变量名 函数名 类型名 方法名
// Go语言中如果标识符首字母是大写的,就表示对外部课件(暴露的,共有的)
// 小写类似于 private 如果要是大写,则写上注释
```



## 方法 

```go
// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量,多用类型名
// 类似于java中的 类方法
```



 func (接受者变量x   接受类型xxx)  方法名 (参数) 返回类型{

} 

### 什么时候应该实用指针类型接收者

1. 需要修改接收者中的值
2. 接受者是拷贝代价比较大的大对象
3. 保证一致性,如果有某个方法使用了指针接收者,那么其他的方法也应该使用指针接收者

[在Go中恰到好处的内存对齐](https://segmentfault.com/a/1190000017527311?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com)

### 接受者 

接受者通常使用类型首字母的小写,不建议使用诸如`this`和`self`这样的

### 标识符

// 标识符: 变量名 函数名 类型名 方法名

// go语言如果标识符首字母大写的,就表示对外部包可见(暴露的,共有的)

### 匿名字段

没有名字的

### 嵌套结构体

### 匿名嵌套结构体

### 匿名其那套结构体的字段冲突

老老实实标准调用 直接写全

```go
fmt.Println(p1.name, p1.workPlace.city)
fmt.Println(p1.name, p1.address.city)
```

## 结构体与JSON



## 作业

1. 把上课写的代码自己敲一遍
2. 把上课的函数版学生管理系统自己写一遍
3. 把函数版学员信息管理系统改写成方法版本(提示: 谁的方法) `可参照哪个08method那个  `