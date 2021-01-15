# day05 课上笔记

# 今日内容

## 接口(interface)

接口是一种类型,是一种特殊的类型,它规定了变量有哪些方法

**在编程中会遇到一下场景**:

我不关心一个变量是什么类型,我只关心能调用它的什么方法

## 包(package)

## 文件操作



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