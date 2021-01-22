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

## sync包

某个操作想执行一次

`var once sync.Once`

```go
once.Do(func() {
       ...
})
```

![image-20210122150711981](D:\Go\src\chentianxiang.vip\studygo\day08\RAEDME.assets\image-20210122150711981.png)

### sync.map

像这种场景下就需要为map加锁来保证并发的安全性了，Go语言的`sync`包中提供了一个开箱即用的并发安全版map–`sync.Map`。开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。同时`sync.Map`内置了诸如`Store`、`Load`、`LoadOrStore`、`Delete`、`Range`等操作方法。

```go
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```



## atomic

原子操作  底层还是锁,比直接用互斥锁效率高

## 网络编程

![image-20210122161614532](D:\Go\src\chentianxiang.vip\studygo\day08\RAEDME.assets\image-20210122161614532.png)

## TCP客户端和服务端

**server**

```go
package main

import (
	"fmt"
	"net"
)

// tcp server 端
func processConn(conn net.Conn) {
    defer conn.Close()
	// 3.与客户端通信
	var temp [128]byte
	n, err := conn.Read(temp[:])
	if err != nil {
		fmt.Printf("read from conn failed,err:%v\n", err)
		return
	}
	fmt.Println(string(temp[:n]))
}

func main() {
	// 1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("start tcp server on  127.0.0.1:20000,err:%v\n", err)
		return
	}
	for {
		// 2. 等待别人来跟我建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed,err:%v\n", err)
			return
		}
		go processConn(conn)
	}

}

```



**client**

```go
package main

import (
	"fmt"
	"net"
	"os"
)

// tcp client

func main() {
	//1. 与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
    defer conn.Close()
	if err != nil {
		fmt.Printf("dial 127.0.0.1:20000 failed,err:%v\n", err)
		return
	}
	// 2. 发送数据
	var msg string
	if len(os.Args) < 2 {
		msg = "hello wang ye "
	} else {
		msg = os.Args[1]
	}
	conn.Write([]byte(msg))
	conn.Close()
}

```

## TCP黏包

### 大端和小端

![image-20210122185617268](D:\Go\src\chentianxiang.vip\studygo\day08\RAEDME.assets\image-20210122185617268.png)