# day10课上笔记

# 本周分享

两个面试题:

**leetcode刷题**,**每天一道题强壮程序员**

**数据结构和算法很重要,找机会抓紧补上**  **数据机构好像补过了23333333!**

**面试题**:

**如何判断链表有闭环**

![image-20210125145618834](D:\Go\src\chentianxiang.vip\studygo\day10\README.assets\image-20210125145618834.png)

**走楼梯, 一次一步和一次两步,一共有多少种方案**

最后要么剩下一个台阶,要么剩下两个台阶,取二者之和

```go
func(n int)int {
    if n==1 {
        return 1
    }
    if n=== 2 {
        return 2
    }
    return f(n-1)+f(n+2)
}
```

// 由于递归效率太低,下面是非递归思路

```go
讨论情况,x 存第一次 n-2个台阶的路数, y存第二次n-1个台阶的路数,第三次n个台阶的路数就为x+y
因为算f(n-1)的时候,已经包含了f(n-2)的了,故每次算好路数存到对应变量中,最后return x+y 即可
```



![image-20210125150347302](D:\Go\src\chentianxiang.vip\studygo\day10\README.assets\image-20210125150347302.png)

# 内容回顾

## net/http包的用法

如何发请求

当需要频繁发送请求的时候(每5秒从阿里云同步接口数据):定义一个全局的var clinet (单例),后续发送请求的操作就是这个全局的client

## 单元测试

### 单元测试

xxx/cc.go

单元测试文件名必须是`xxx/cc_test.go`

go内置的测试工具

```go
go test -v
```

单元测试函数的格式

```go
import "testing"
// Test开头后函数名字
func TestSplit (t *testing.T){
    t.Fatal()
}
```

### 性能测试/基准测试

函数格式:

```go
func BenchmarkSplit(b *test.B){
    // b.N: 被测试函数执行的次数
}
```

执行命令

`go test -bench -v`

并行测试

### setup和teardown

## pprof

记录cpu快照信息

记录内存的快照信息

**参考昨天代码或者博客** [pprof](https://www.liwenzhou.com/posts/Go/performance_optimisation/)

## flag

### os.Args

```
./xxx   a b c
```

osArgs:["./xxx" "a" " b" "c"]

### flag标准库

#### 声明变量

```bash
./xxx -name="周林"  -age=8999 a b c
```

两种方法

```go
nameVal1 := flag.String("name", "卢明辉","请输入dsb的名字")// 返回的是指针类型

var nameVal2 string
flag.StringVar(&nameVal2, "name", "卢明辉", "请输入dsb的名字")// 把一个已经存在的变量绑定到命令行flag参数
```

必须要调用:

```go
flag.Parse() // 解析命令传入的参数,赋值给对应的变量
```

其他方法:

```go
flag.Args() // 返回命令行参数后的其他参数,以[]string类型
flag.NArg()	// 返回命令行参数后的其他参数个数
flag.NFlag() // 返回使用的命令行参数个数
```

# 今日内容

## MySQL

### 数据库

常见的数据库SQLite,MySQL,SQLServer,postgreSQL,Oracle

MySQl主流的关系型数据库,类似的还有`postgreSQL`

**关系型数据库**:

用表来存一类数据

表结构设计的三大范式: (漫画数据库)



### MSQL知识点

### SQL语句

DDL: 操作数据库的

DML: 表的增删改查

DCL: 用户及权限

### 存储引擎

MySQL支持插件式的存储引擎

常见的存储引擎: MyISAM和InnoDB.

MyISAM:

1. 查询速度快
2. 只支持表锁
3. 不支持事务

InnoDB:

1. 整体速度快
2. 支持表锁和行锁
3. 支持事务

**事务**:

把多个SQL操作当成一个整体.

**事务特点**:

ACID:

1. 原子性: 事务要么成功要么失败,没有中间状态
2. 一致性:..
3. 隔离性:事务之间是相互隔离的
   1. 隔离的四个级别 如图(面试会问)
4. 持久性: 事务操作的结果是不会丢失的

![image-20210125161415726](D:\Go\src\chentianxiang.vip\studygo\day10\README.assets\image-20210125161415726.png)

**索引:**

索引的原理是啥: B树和B+树(**还好老子学了!)**

索引的类型

索引的命中

分库分表

SQL注入

SQL慢查询优化



**MySQL主从**:

​	binlog

MySQL读写分离



### Go操作MySQL

####  database/sql

原生支持连接池,是并发安全的

这个库标准没有具体的实现,只是列出了一些需要第三方库实现的具体内容

#### 下载驱动

```bash
go get -u github.com/go-sql-driver/mysql
```

`go get 包路径` 就是下载第三方依赖

将第三方的依赖默认保存在$GOPATH/src/

`ps,我反正没用命令行下载成功`

![image-20210125165116376](D:\Go\src\chentianxiang.vip\studygo\day10\README.assets\image-20210125165116376.png)

**于是乎我先把该包从github上下载到了src下,然后执行**

`go install-u github.com/go-sql-driver/mysql` 这样!

**错误1**

```go
GoLang 解决VsCode中提示错误 go: cannot find main module, but found .git/config in D:\XXX\src\XXX to create a module there, run: cd .. && go mod init
```

<img src="D:\Go\src\chentianxiang.vip\studygo\day10\README.assets\image-20210125172201907.png" alt="image-20210125172201907" style="zoom: 80%;" />

### 使用



![image-20210125173906990](D:\Go\src\chentianxiang.vip\studygo\day10\README.assets\image-20210125173906990.png)

在`github.com/go-sql-driver/mysql/driver.go ` 里面执行了init(),把sql注册了进去,所以我在Open()的时候填上mysql就可以加载该驱动了

![image-20210125174000588](D:\Go\src\chentianxiang.vip\studygo\day10\README.assets\image-20210125174000588.png)

#### 注意

`若是加上 :=  会报空指针错误`

![image-20210125175323066](D:\Go\src\chentianxiang.vip\studygo\day10\README.assets\image-20210125175323066.png)

## Redis

KV数据库 

```bash
go get -u github.com/go-redis/redis
```

Redis的用处:

1. cache缓存
2. 简单示例
3.  排行榜

书:<Redis实战>

**由于网络等原因,故下载东西有些出路,看代码即可**

**等用到后,再做详细探讨!**

[Redis的使用详情](https://www.liwenzhou.com/posts/Go/go_redis/)

## NSQ

[NSQ详细](https://www.liwenzhou.com/posts/Go/go_nsq/)



Go语言开发的轻量级消息队列

![image-20210127140429242](D:\Go\src\chentianxiang.vip\studygo\day10\README.assets\image-20210127140429242.png)

### 组件

```bash
nsqdlookeupd.exe

```

默认在本机的127.0.0.1:4160启动

```bash
nsqd.exe -broadcast -address=127.0.0.1 -lookupd -tcp -address=127.0.0.1:4160
```





## 包的依赖管理 go module

Go1.11之后官方出的依赖管理工具







