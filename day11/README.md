# day11课上笔记

# 今日内容

## 依赖管理go module

1.15默认开启GO111MODULE的

直接无脑安装,不需要配置工作区的环境变量

可以手动开启 `set GO111MODULE=on ` 

可以设置环境变量 `go env -w GO111MODULE=on`

更改 go get 的源 `go env -w GOPROXY=https://goproxy.cn,direct`

初始化init 需要 在项目下  `go mod init 该包的名字`  如`go mod init xxx`  这样go build后的 exe文件就是该名字
ps: **go buidl 之前,必须有mod文件, 先 执行上面步骤**

![image-20210127183720501](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210127183720501.png)

![image-20210127183713603](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210127183713603.png)

然后 `go get ` 下载文件就好了

![image-20210128221342596](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210128221342596.png)

### go.mod

包都放在了GOPATH下的pkg

### go.sum 文件

详细的包名和版本信息

### 常用命令

```bash
go mod init [包名] // 初始化项目
go mod tidy // 检查代码里的依赖去更新go.mod文件的依赖
go get
go mod download

```



## Context

[Context 的代码](https://www.liwenzhou.com/posts/Go/go_context/)

非常重要!

### 两个默认值

```go
context.Background()
context.TODO()
```

### 四个方法

```go
context.WithCancel(context.Background())
context.WithDeadline(context.Background(), time.Time)
context.WithTimeout(context.Background(), time.Duration)
context.WithValue(context.BackGround(), key, value)
```



## 服务端Agent开发(日志收集项目)

zoolleeper , kafka部署文档 : https://docs.qq.com/doc/DTmdldEJJVGtTRkFi

kafka的使用:

1. ![image-20210129210553200](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129210553200.png)

2.  当前文件夹下cmd   

   1. ```bash
      bin\windows\kafka-server-start.bat   config\server.properties
      ```

   2. ```bash
      bin\windows\zookeeper-server-start.bat config\zookeeper.properties
      ```

​      

​	kafka终端读取数据:

```bash
bin\windows\kafka-console-consumer.bat   --bootstrap-server=127.0.0.1:9092 --topic=web_log --from-beginning
```

### 配置加载ini     gopkg.in/ini.v1 

# 截图和列出来的问题都是重点



### kafka

![image-20210129202241168](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129202241168.png)

1. Kafka集群的架构

   1. broker
   2. topic
   3. partition: 分区, 把用一个topic分成不同的分区,提高负载
      1. leader: 分区的主节点(老大)
      2. follower: 分区的从节点(小弟)
   4.  COnsumer Group

2. 生产者往Kafka发送数据的流程(6步)

   ![image-20210129202604992](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129202604992.png)

3. Kafka选择分区的模式(3种)

   1. 指定往哪个分区写
   2. 指定key,kafka根据key做hash然后决定写哪个分区
   3. 轮询方式

4. 生产者发往kafka发送数据的模式(3种)

   1. `0`: 把数据发送给leader就成功,效率高,安全性最低
   2. `1`: 把数据发送给leader, 等待leader回ACK
   3. `all`: 把数据发送给leader,follower从leader拉取数据恢复ack给leader,leader在回复ACK

5. 分区存储文件的原理

6. 为什么kafka快?  随机读变成了顺序读

7. 消费者消费数据的原理

8. ![](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129212029558.png)

   1. **LogAgent的工作流程**

      1.  读体质-- `tailf`第三方库
      2. 往kafka写日志-- `sarama`第三方库

      ![image-20210130105637665](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210130105637665.png)

   ​	









![image-20210129195726592](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129195726592.png)

![image-20210129195827494](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129195827494.png)



![image-20210129195941227](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129195941227.png)

### 项目架构设计

![image-20210129200012755](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129200012755.png)

![image-20210129200257240](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129200257240.png)

![image-20210129200346320](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129200346320.png)

![image-20210129200529619](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129200529619.png)

### kafka和zookeeper

![image-20210129200728372](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129200728372.png)

![image-20210129201037210](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129201037210.png)

![image-20210129200928566](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129200928566.png) ![image-20210129201227361](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129201227361.png)

![image-20210129201243290](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129201243290.png)

![image-20210129201413978](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129201413978.png)

![image-20210129201835816](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129201835816.png)

![image-20210129202031235](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129202031235.png)

![image-20210129203258634](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129203258634.png)

![image-20210129203310444](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129203310444.png)
![image-20210129203501934](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129203501934.png)

![image-20210129203549532](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129203549532.png)

![image-20210129203751222](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129203751222.png)

![](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129204040264.png)

![image-20210129204104216](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129204104216.png)

![image-20210129204136009](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129204136009.png)

![image-20210129211858090](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129211858090.png)

![image-20210129211911196](D:\Go\src\chentianxiang.vip\studygo\day11\README.assets\image-20210129211911196.png)



 

