# day03课上笔记

如果什么都是很容易的,那还叫什么选择.

走上坡路肯定是会累的

永远不要高估自己



// 明天继续 2021-01-08 22:56:42

// 2021-01-09 23:01:53  明天继续 ,书已经到齐, 键盘明天到,开始正式规划3周冲刺计划!

// 2021-01-10 10:52:10 键盘明天到

# 今日内容

## 函数

### 函数的定义

### 基本格式

参数的格式

有参数的函数

参数类型简写

可变参数

### 返回值的格式

有返回值

多返回值

命名返回值

### 变量作用域

1.全局作用域

2.函数作用域

	1. 先在函数内部找变量 ,找不到往 外部找

3.代码块作用域

### 高阶函数

函数也是一种类型,他可以作为参数,也可以作为返回值



### 匿名函数

没有名字的函数

### 闭包

 **就是把当前函数包装成目标函数想要的那种**

```go
// 什么是闭包
// 闭包是一个函数,这个函数包含了他外部作用域的一个变量

// 底层原理:
// 1.函数可以作为返回值
// 2.函数内部查找变量的顺序,现在自己内部找,找不到往外层找
```

 

 ![image-20210109110726126](D:\Go\src\chentianxiang.vip\studygo\day03\README.assets\image-20210109110726126.png)



### fmt 库  

[参考文档 API]( https://studygolang.com/pkgdoc)



### 今日难点:

 1. 函数的定义

 2. 高阶函数

 3. 函数类型

 4. 闭包

 5. defer

 6. panic/ recover

    



# 内容回顾

## 运算符

### 算数运算符

+-*/

### 逻辑运算符

&&  || !  .....

### 位运算符

`>>` `<=` ......

### 赋值运算符

=  +=  ......

`++` `--`  是独立的语句,不属于赋值运算符

### 比较运算符

< <=  !=  .....

## 数组(Array)

`var ages [30]int `

`var ages [30]string `

数组包含元素的类型和元素的个数. 元素的个数(数组的长度)属于数组类型的一部分

数组是值类型

## 切片

切片的定义

切片不存值,他就像一个框,去底层数组框值

切片: 指针,长度,容量

切片的扩容策略

1. 如果小于1024,那么就直接两倍
2. 如果大于1024,就按照1.25倍去扩容
3. 如果申请的的容量大于原来的2倍,直接扩容至新申请的容量
4. 具体储存的值类型不同,扩容策略也有一定的不同

切片的删除:  通过copy向前挪移,仿佛数组删除那样

## 指针

只需要记住两个符号: `&` 和`*`

## map



