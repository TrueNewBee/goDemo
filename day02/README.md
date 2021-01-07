2021-1-6 20:14:08

2021-1-7 22:26:21 明天继续

# 02day笔记

## 今日内容  

### 流程控制 

  for 多层循环,  switch ,goto(不推荐使用 )

### 运算符

见代码

### 复杂数据类型

函数  

### 数组

![image-20210107160323637](D:\Go\src\chentianxiang.vip\studygo\day02\README.assets\image-20210107160323637.png)

### 切片(slice)

切片指向了一个底层的数组

切片的长度就是它元素的个数

切片的容量就是底层数组从切片的第一个元素到最后一个元素的数量

### 切片的本质

切片就是一个框,框住了一块连续的内存.

切片属于引用类型,真正的数据都是保存在底层的数组里的

`s1 = append(s1, ss...) //  ...表示拆开`

### 切片扩容原理

![image-20210107195859554](D:\Go\src\chentianxiang.vip\studygo\day02\README.assets\image-20210107195859554.png)



### 面试题

```go

//练习题  面试题
func main() {
	var a = make([]string, 5, 10) //创建切片,长度为5,容量为10
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)      // [     0 1 2 3 4 5 6 7 8 9]
	fmt.Println(cap(a)) // 20
}

```

### 指针

Go语言中不存在指针操作,只需要记住两个符号

1. `&`:取地址
2. `*`:根据地址取值

### make和new的区别 (面试题)

1. make和new都是用来申请内存的
2. new很少用,一般用来给基本数据类型申请内存,`sting` , `int`,返回的是对应类型的指针(*string, *int)
3. make是用来给`slice`,`map`,`chan`申请内存的,make函数返回的是对应的这三个类型本身

### map



匿名函数和封包



##  内容回顾

### GO安装

`$GOPATH`: 你写的Go代码的工作区,保存你的Go代码

`go env`

`GOPATH/bin` 添加环境变量  `go install`命名会把生成的二进制可执行文件拷贝到``

`GOPATH/bin`



## Go 命令

`go build`: 编译Go程序

`go build -o "xx.exe" `: 编译成xx.exe文件

`go  run main.go`: 像执行脚本一样执行main.go文件

`go install`: 先编译后靠背

## 快捷键配置

![image-20210107214539773](D:\Go\src\chentianxiang.vip\studygo\day02\README.assets\image-20210107214539773.png)

## Go语言文件基础语法

存放Go源代码的文件后缀名是`.go`

文件第一行: `package`关键字声明包名 

**如果要编译一个可执行文件,必须要有main包和main函数(入口函数)**

```go
//  单行注释
/*
多行注释
*/
```

Go语言函数外的语句必须以关键字开头

函数内部定义的变量必须使用



## 变量

3种声明方式:

1. `var name1 string`
2. `var name2 = 沙河娜扎`
3. 函数内部专属:  `name3 := 沙河王子` 

匿名变量(哑元变量)

当有些数据必须要用变量接收但是又不使用它时,就可以用_来接收



## 常量

`const PI = 3.1415926`

`const UserNotExit = 10000`



iota:实现枚举

两个要点:

1. `iota`在const关键字出现时将被重置为0
2. const中每新增一行常量声明,iota加1



## 流程控制

```go
	// age := 19
	// if age > 18 {
	// 	fmt.Println("澳门首家线上赌场上线了")
	// } else {
	// 	fmt.Println("改写暑假作业了")
	// }

	// 作用域
	// age变量此时只在if条件判断语句中生效, 节约内存
	if age := 19; age > 18 {
		fmt.Println("澳门首家线上赌场上线了")
	} else if age < 1 {
		fmt.Println("改写暑假作业了")
	} else {
		fmt.Println("玩吧")
	}
	// fmt.Println(age)
```

 

## for 循环

```go
	// 基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种1  省略初始语句 常用
	// var i = 5
	// for ; i < 19; i++ {
	// 	fmt.Println(i)
	// }

	// 变种2 少用
	var i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 无限循环  不要轻易用
	// for {
	// 	fmt.Println("123")
	// }

	// for range循环
	s := "hello 沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

```

## 99乘法表

```go
// 打印9*9乘法表 
for i := 1; 1<10 ; i++{
    for j := 1; j <= i; j++{
        fmt.Printf("%d*%d=%d\t", j, i, j*i )
    } 
    fmt.Println()
}
```



## 基本数据类型

### 整型

​	无符号整型: `unit8` `unit`16 `unit32`  `unit64`

​	带符号整型: `int8` `int16` `int32` `int64`

​	`uint`和`int`:具体是看系统32位还是64位

​	`uinptr`:表示指针

### 其他进制数

Go语言中没办法直接定义二进制数

```go
// 八进制数
var n1 = 0777
// 十六进制数
var n2 = 0xfff
fmt.Println(n1, n2)
fmt.Printf("%o\n", n1)
fmt.Printf("%x\n", n2)
```



### 浮点数

​	`float64`和`floate32`

​	Go语言中浮点数默认是`float64`

###  复数

`complex128`和`complex64`

### 布尔值

`true`和`flase`

不能和其他的类型做转换

### 字符串

常用方法

字符串不能修改

### byte和rune类型

都属于类型别名

````go
	// for range循环
	s := "hello 沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

````





### 字符串,字符,字节都是什么?

字符串: 双引号包裹的是字符串

字符: 单引号包裹的是字符,单个字母,单个符号,单个文字

字节: 1byte=8bit

go语言中字符串都是utf8编码中一个常用汉字一般占用3个字节











 





 