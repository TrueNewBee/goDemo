# day05 课上笔记

# 今日内容

## 插入一个刘谦小魔术
(x *9 (个位+十位  肯定是9 自己记住) x 9 -2 =79  可以写成47 (七九 倒过来) )
## 接口(interface)

接口是一种类型,是一种特殊的类型,它规定了变量有哪些方法

**在编程中会遇到一下场景**:

我不关心一个变量是什么类型,我只关心能调用它的什么方法

### 接口的定义

```go
type 接口名 interface {
    方法名1(参数1, 参数2 ...)(返回值1, 返回值2,)
    方法名2(参数1, 参数2 ...)(返回值1, 返回值2,)
}
```

用来给变量\参数\返回值等设置类型

### 接口的实现

一个变量如果实现了接口中规定的所有的方法,那么这个变量就实现了这个接口,可以称为这个接口类型的变量

![image-20210116111541137](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210116111541137.png)

![image-20210116113139750](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210116113139750.png)

### 使用值接受者实现接口与使用指针接受者实现接口的区别

**使用值接受者实现接口,结构体类型和结构体指针类型的变量都能存**

**指针接受者实现接口只能存结构体指针类型的变量**

### 接口和类型的关系

多个类型可以实现同一个接口

一个类型可以实现多个接口

### 空接口

没有必要起名字,通常定义成下面的格式

```go
interface{}	// 空接口
```

所有的类型都实现了空接口,也就是任意类型的变量都能保存到空接口中



#### 注意

```go
关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。
```



**类型断言**

![image-20210116132558249](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210116132558249.png)

## 包(package)

[包](https://www.liwenzhou.com/posts/Go/11_package/)

注意事项：

- import导入语句通常放在文件开头包声明语句的下面。
- 导入的包名需要使用双引号包裹起来。
- 包名是从`$GOPATH/src/`后开始计算的，使用`/`进行路径分隔。
- Go语言中禁止循环导入包。
- 向被别的包调用的标识符都要首字母大写!
- 单行和多行导入
- 导入包的时候可以指定别名
- 导入包不想使用包内的标识符,需要使用匿名导入 _
- 每个包导入的时候回自动执行一个名为`init()`函数,他没有参数也没有返回值也不能手动调用
- 多个包都定义了`init()`函数,它们执行顺序如下图



**main包中的操作**

![image-20210116180203225](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210116180203225.png)

### 导入包的顺序

![image-20210116180310812](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210116180310812.png)

## 文件操作

[文件操作](https://www.liwenzhou.com/posts/Go/go_file/)

```go
package main

import (
	"fmt"
	"io"
	"os"
)

// 打开文件

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 读文件
	// var temp = make([]byte, 128)	// 指定读的长度
	var temp [128]byte
	for {
		n, err := fileObj.Read(temp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Printf(string(temp[:n]))
		if n < 128 {
			return
		}
	}

}

```

**bufio.NewReader() 是一个接口类型,只要出来的文件有东西就行**

![image-20210116225822382](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210116225822382.png)

## 作业

1. 把课上读文件的3种方式写一下
2. 把课上写文件的3中方式和copyFile哪个函数写一下  博客那的

自己写一个日志库

接口: 用处 日志可以输出到终端,也可以输出到kafka

需求:

	1. 可以往不同的位置记录日志
 	2.  日志分为五种级别    就是那个 11111  吃 睡 打

  ![image-20210116231209833](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210116231209833.png)

# 内容回顾

## 自定义类型和类型别名

```go
type myInt int  // 自定义类型
type newInt = int	// 类型别名
```

**类型别名只在代码编写过程中有效,编译完之后就不存在了,内置的*byte和rune都属于类型别名**

## 结构体

基本的数据类型: int float等表示现实中的物件有局限性

编程是代码解决现实生活中的问题:

```go
var name = "保德路"
```

结构体是一种数据类型,一种我们自己造的可以保存多个维度数据类型

```go
// 有名字的结构体
type person struct{
    name string
    age int
    addr string
}
var a = person{
    "元帅",
    12,
    "沙河",
}
```



## 匿名结构体

多用于临时场景

```go
// 普通 有名字结构体
var temp = struct {
		x int
		y int
	}
var a = temp{
    1,
    2,
}
// 匿名结构体
var b = struct {
		x int
		y int
	}{10, 20}
```



## 构造函数

```go
// 结构体类型为person
type person struct {
	name string
	age  int
}
// 构造(结构体变量的)函数: 由于每次都实例化结构体麻烦,故封装到函数里面,叫构造函数
// 返回值是对应的结构体类型
func newPerson(n string, i int) person {
	return person{
		name: n,
		age:  i,
	}
}
func main (){
    //用构造函数直接用完! (多次使用,结构体字段多,推荐)
    p3 := newPerson("周林", 13) 
	fmt.Println(p3)
}
```

![image-20210114113433181](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210114113433181.png)

## 方法和接受者

方法是有接受者的函数,接受者指的是哪个类型的变量可以调用这个函数

![image-20210114154419827](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210114154419827.png)

结构体是值类型

![image-20210114162925396](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210114162925396.png)

## 结构体的嵌套

![image-20210114163350505](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210114163350505.png)

## 结构体的匿名字段

![image-20210114163303210](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210114163303210.png)

## JSON序列化与反序列化

 经常出现问题:

​	1.结构体内部的字段要大写!!!!不大写别人是访问不到的

​	2. 反序列化时要传递指针!

![image-20210114163707782](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210114163707782.png)

![image-20210114164815021](D:\Go\src\chentianxiang.vip\studygo\day05\README.assets\image-20210114164815021.png)