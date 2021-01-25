# day09 课上笔记

# 课程安排:

日志收集项目:9.7\9.21\9.28

gin框架和微服务: 4天

Docker和k8s: 4天

# 今日分享

注释\日志\单元测试\

要做就做专业的,从一开始就规范起来,

# 内容回顾

## 互斥锁

`sync.Mutex`

 是一个结构体,是值类型,给函数传参数的时候要传指针

### 两个方法

```go
var lock sync.Muxtex
lock.Lock()	// 加锁
lock.Unlock()// 解锁
```

### 为什么要用锁?

防止同一时刻多个goroutine操作同一个资源

(有修改公共资源时候,考虑加锁)

## 读写互斥锁

### 应用场景

适用于读多写少的场景下,才能提高程序的执行效率.

### 特点

1. 读的goroutine来了获取的是**读锁**,后续的goroutine能读不能写
2. 写的goroutine来了获取的是**写锁**,后续的goroutine不管是读还是写都要等待获取锁

### 使用

```go
var rwLock sync.RWMutex
rwLock.RLock()	// 获取读锁
rwLock.RUnlock()// 释放读写

rwLock.Lock()	// 获取写锁
rwLock.Unlock()	// 释放写锁
```





## 等待组

`sync.Waitgroup`

用来等goroutine执行完再继续

是结构体,是值类型,给函数传参数时候要传指针

### 使用

```go
var wg sync.WaitGroup

wg.Add(1)	//起几个goroutine就加几个计数
wg.Done() 	// 在goroutine对应的函数中,函数要结束的时候表示goroutine完成,计数器 -1
wg.Wait()	// 阻塞,等待所有的goroutine都结束
```



## Sync.once

### 使用场景

  某些函数只需要执行一次的时候,就可以使用`sync.Once`

比如 bolg加载图片的那个例子

``` go
var once sync.Once
once.Do()	// 接受一个没有参数也没有返回值的函数,如有需要可以使用闭包(匿名函数)
```

## sync.Map

### 使用场景

并发操作一个map的时候,内置的map不是并发安全的

### 使用

是一个开箱急用的并发安全的map (不需要make初始化)

```go
var syncMap sync.Map

// Map[key] = value // 原生map
syncMap.Store(key, value)
syncMap.Load(key)
syncMap.Delete()
syncMap.Range()
syncMap.LoadOrStore()
syncMap.Range()

```

### 原子操作

Go语言内置了一些针对内部的基本类型的一些并发安全的操作

```go
var i int64 =10
atomic.AddInt64(&i, 1)
```

## 网络编程

### 

![image-20210123114042831](D:\Go\src\chentianxiang.vip\studygo\day09\README.assets\image-20210123114042831.png)

# 今日内容

## 互联网的协议

 OSI七层模型

HTTP: 超文本传输协议

**裸体的人**

规定了:浏览器和网站服务器之间通信的规则

HTML: 超文本标记语言

学的就是标记的符号,标签

CSS:层叠样式表

**让人穿上衣服/化妆**

规定了HTML中标签的具体样式(颜色\背景\大小\浮动...)

javascrip

**让人动起来**

## 请求发送流程图



![image-20210123150857838](D:\Go\src\chentianxiang.vip\studygo\day09\README.assets\image-20210123150857838.png)

## 单元测试

开发自己给自己的代码写的测试

测试函数覆盖率:100%

代码覆盖率:60%

```go
以 xx_test.go 命名文件
测试函数命名:  func TestXxx(t *testing.T) {
    ret := xxx  	// 输入测试条件
    want :=  xxx	// 想得到结果
    if  ret != want {
        t.Errorf("want :%v but got : %v\n", want, ret)
    }
}
在终端中 go test  -v 即可测试
查看测试覆盖率 go test -cover
讲覆盖率相关记录输出到一个文件中, go test -cover -coverprofile=cover.out
然后执行go tool cover -html=cover.out,
```



## pprof调试工具

[pprof性能调优](https://www.liwenzhou.com/posts/Go/performance_optimisation/)

### context