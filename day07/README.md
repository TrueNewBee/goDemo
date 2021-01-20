# day07课上笔记

# 内容回顾

## time

`2006-01-02 15:04:05:000`

### 时间类型

- `time.Time`:`time.now()`
- 时间戳:
  - `time.Now().Unix()`: 1970.1.1到现在的秒数
  - `time.Now().UnixNano()`:1970.1.1 到现在的纳米秒

### 时间间隔类型

- `time.Duration`: 时间间隔类型
  - `time.Second`

```go
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
```

## 时间操作

时间对象+/- 一个时间间隔对象

```go
// 时间间隔
fmt.Println(time.Second)
// now + 24小时
fmt.Println(now.Add(24 * time.Hour))	
// Sub 两个时间相减
nextYear, err := time.Parse("2006-01-02", "2021-08-04")
if err != nil {
    fmt.Printf("parse time failed, err:%v\n", err)
    return
}
d := nextYear.Sub(now)
```

after/befor

### 时间格式化

2006-01-02  15:04:05.000

### 定时器

```go
// 定时器	
timer := time.Tick(time.Second)
for t := range timer {
    fmt.Println(t) // 1秒钟执行一次
}
```

### 解析字符串格式的时间(时区)

```go
func f2() {
	now := time.Now() // 本地时间
	fmt.Println(now)
	// 明天这个时间
	// 按照指定格式取解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2021-01-17 16:31:29")
	// 按照东八区的时区和格式取解析一个字符串格式的时间
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed ,err: %v\n", err)
		return
	}
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-01-17 16:31:29", loc)
	if err != nil {
		fmt.Printf("load loc failed ,err: %v\n", err)
		return
	}
	fmt.Println(timeObj)
	// 时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}
```

## 日志库

time

文件操作

`runttime.Caller()`

之前学的内容整合

## 反射

 接口类型的变量底层是分为两个部分:动态类型和动态值

反射的应用:`json`等数据解析\ORM等工具

### 反射的两个方法

- `reflect.TypeOf()`
- `reflect.ValueOf()`



## ini解析作业

自己可以封装成 go ini   发到GitHub 开源项目 

### 面向GitHub 编程  233333

![image-20210119204310451](D:\Go\src\chentianxiang.vip\studygo\day07\README.assets\image-20210119204310451.png)



# 今日内容

## 并发

并发：同一时间段内执行多个任务（你在用微信和两个女朋友聊天）。

并行：同一时刻执行多个任务（你和你朋友都在用微信和女朋友聊天）。

Go语言的并发通过`goroutine`实现。`goroutine`类似于线程，属于用户态的线程，我们可以根据需要创建成千上万个`goroutine`并发工作。`goroutine`是由Go语言的运行时（runtime）调度完成，而线程是由操作系统调度完成。

Go语言还提供`channel`在多个`goroutine`间进行通信。`goroutine`和`channel`是 Go 语言秉承的 CSP（Communicating Sequential Process）并发模式的重要实现基础。

## goroutine

Go语言中的`goroutine`就是这样一种机制，`goroutine`的概念类似于线程，但 `goroutine`是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。

### 启动goroutine

```go
// goroutine

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动只有会创建一个主goroutine去执行
func main() {
	for i := 0; i < 1000; i++ {
		go hello(i) // 开启一个单独的goroutine去执行hello函数(任务)
		// 匿名函数
		go func(i int) {
			fmt.Println(i) //	用的是参数的的那个i而不是外面的i
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
}

```

### `goroutine`什么时候结束?

goroutine 对应的函数结束了,goroutine结束了

`main`函数执行完了,有`main`函数创建的那些`goroutine`都结束了

### math/rand

```go
// rand随机数
func f() {
	rand.Seed(time.Now().UnixNano())	// 保证每次执行的时候有点不一样

	for i := 0; i < 5; i++ {
		r1 := rand.Int()    // int64
		r2 := rand.Intn(10) // 0<= x <10
		fmt.Println(r1, r2)
	}
}
```

[并发](https://www.liwenzhou.com/posts/Go/14_concurrence/)

### `goroutine`调度模型

`GMP`   面试会问

`M:N`	把m个goroutine 分配给n个操作系统线程

goroutine初始栈的大小是2k



## channel

```go
var b chan int // 需要指定通道中元素的类型
```

通道必须使用make函数初始化才能使用

```go
b = make(chan int)     // 不带缓冲区通道的初始化
b = make(chan int, 16) // 带缓冲区通道的初始化
```

### 通道的操作

`<-`

1. 发送: `cha1 <-1`	  发送到ch1
2. 接收:`x := <- ch1` x接受ch1的数据
3. 关闭: `close()`
4. 单向通道  

```go
// ch1 仅能取值, ch2 仅能存值
func f2(ch1 <-chan, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() { close(ch2) }) // 确保某个操作只执行一次
}
```

![image-20210120162719704](D:\Go\src\chentianxiang.vip\studygo\day07\README.assets\image-20210120162719704.png)



## 作业

1. 为了保证业务代码的执行性能将之前写的日志库改写为异步记录日志方式。

   业务代码记日记先存放到通道中,然后起一个后台goroutine专门从通道中取日志往文件里写

   1. 日志库中channel怎么用?
   2. 什么时候起后台的goroutine去写日志到文件中

