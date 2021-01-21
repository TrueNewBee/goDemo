# day08 课上笔记

# 内容回顾

## 并发之goroutine

  ### 并发与并行的区别

### goroutine的启动

将要并发执行的任务包装秤一个函数,调用函数的时候前面加上`go`关键字, 就能够开启一个gorouine去执行该函数的任务

goroutine对应的函数执行完,该goroutine就结束了

程序启动的时候会自动的创建一个goroutine去执行main函数

main函数结束了,那么程序就结束了,由该程序启动的所有其他goroutine也都结束了

### goroutine的本质

goroutine的调度模型: `GMP`

`M:N`:把m个goroutine分配给n个操作系统线程

### goroutine与操作系统线程(os线程)的区别

goroutine是用户态的线程, 比内核态的线程更轻量级一点 ,初始时只占用2k的栈空间

可以轻轻松松开启数十万的goroutine也不会蹦内存

### runtime.GOMAXPROCS

Go1.5之后默认就是操作系统的逻辑核心数,默认跑满CPU

`runtime.GOMAXPROCS(1)`: 只占用一个核.

### work pool 模式

开启一定数据的goroutine去干活



### sync.WaitGroup

`var wg sync.WaitGroup`

- wg.Add(1):计数器+1
- wg.Done(): 计数器-1
- wg.Wait(): 等待

## channel

### 为什么需要channel?

通过channel实现多个goroutine之间得通信

'GSP': 通过通信来共享内存

channel是一种类型,一种引用类型.make函数初始化后才可以使用(slice, map.channel)

**channel的声明**: `var ch chan 元素类型`

**channel的初始化**: `cha = make(chan 元素类型, [可自选 缓冲区大小])`

### channel的操作:

- 发送: `ch<- 100`
- 接受:`x := <- ch`
- 关闭:`close(ch)`   //可以不关闭,因为不用可以被垃圾回收掉

### 带缓冲区的通道和无缓冲区的通道

快递员送快递的实例,有缓冲区就是有快递柜



从通道中取值:

```go
	for {
		x, ok := <-results
		if !ok { // 什么时候ok=false? results被关闭的时候
			break
		}
		fmt.Println(x)
	}
	for x := range results {
		fmt.Println(x)
	}
```

### 单项通道

通常是用做函数的参数,只读通道`<-chan int`和只写通道`chan<-int`

### 常见的异常

![image-20210121145607437](D:\Go\src\chentianxiang.vip\studygo\day08\RAEDME.assets\image-20210121145607437.png)

![image-20210121150439164](D:\Go\src\chentianxiang.vip\studygo\day08\RAEDME.assets\image-20210121150439164.png)

## select

同一时刻有多个通道要操作的场景下,使用select

![image-20210121150634237](D:\Go\src\chentianxiang.vip\studygo\day08\RAEDME.assets\image-20210121150634237.png)

## 关于函数传参

函数穿进去的参数,是一个slice的切片, 直接传a  则是把整体传入, a... 是把正体里面的元素分开了

![image-20210121154943073](D:\Go\src\chentianxiang.vip\studygo\day08\RAEDME.assets\image-20210121154943073.png)

![image-20210121154953214](D:\Go\src\chentianxiang.vip\studygo\day08\RAEDME.assets\image-20210121154953214.png)

# 今日内容

sync包

context

网络编程

