1. 下go  配置环境,  系统和当前环境 改成自己工作目录,, D:\go
2. 不过1.14后不要求,可以参考这篇博客
3. https://www.liwenzhou.com/posts/Go/install_go_dev/

 go evn 





go build (推荐使用)

1.在项目目录下 go build 

2.  编译到其他目录,,,  go  build    项目路径(从GOPATH的src后面开始)

go  run   执行脚本文件

go install    分为两步:

	1. 先编译得到一个可执行文件
	2. 讲可执行文件拷贝到 "GOPTH  D:\Go\bin"



### Go语言文件的基本结构

****



```go
package main

// 导入包
import "fmt"

// 函数外只能放置标识符(变量/常量/函数/类型)的声明

// 程序的入口函数
func main() {
	fmt.Println("hello world")
}

```



# Go语言基础之变量和常量

[慢慢接触就认识了](https://www.liwenzhou.com/posts/Go/01_var_and_const/)

 

### 变量声明



变量声明后才可以使用

``var  s1 string;    `` 声明一个s1字符串变量

面试注意:

`iota`是go语言的常量计数器，只能在常量的表达式中使用。

`iota`在const关键字出现时将被重置为0。const中每新增一行常量声明将使`iota`计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。



### 数据类型



[数据类型](https://www.liwenzhou.com/posts/Go/02_datatype/)



进制转换



### 字符串

Go语言中字符串是用双引号包裹的!

Go语言中的单引号包裹的是字符!

```go
// 字符串
s := "hello  沙河"
// 单独的字母,汉字,符号表示一个字符
c1 := 'h'
c2 := '1'

// 字节: 1字节=8bit (8个进制位)
// 1个字符'A'=1个字节
// 1个utf8编码的汉子 '沙' = 一般3个字节

```

// 2021-1-4 21:40:45 明天早起继续 明天做笔记



**go真tm简洁,很舒服,比python爽 ,,以后早上go,下午java!!!!!**

坚持!!





for  if  rune 看代码 明天早上做笔记

// 99 乘法表



**晚上回来补笔记!**  `2021-1-6 10:52:01`

